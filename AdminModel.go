package grip

import (
	"time"

	sj "github.com/brianvoe/sjwt"
	"github.com/deta/deta-go/service/base"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
  Key      string `json:"key"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

func GenAdminToken(admin Admin) (string) {
  info := &Admin {
    Key: admin.Key,
    Email: admin.Email,
  }
  claims, _ := sj.ToClaims(info)
  claims.SetExpiresAt(time.Now().Add(8760 * time.Hour))

  token := claims.Generate(JWTKey)
  return token
} 

func ParseAdminToken(token string) (Admin, error) {
  hasVerified := sj.Verify(token, JWTKey)

  if !hasVerified {
    return Admin {}, nil
  }

  claims, _ := sj.Parse(token)
  err := claims.Validate()
  admin := Admin {}
  claims.ToStruct(&admin)

  return admin, err
}

func GetAdmin(email string) (Admin, error) {
  var admins []Admin

  _, err := Admins.Fetch(&base.FetchInput {
    Q: base.Query {
      {"email": email},
    },
    Dest: &admins,
    Limit: 1,
  })

  return admins[0], err
}

func (admin *Admin) Put() (error) {
  if admin.Key == "" {
    admin.Key = GenKey()
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
  if err != nil {
    return err
  }
  admin.Password = string(hashedPassword)

  _, err = Admins.Put(admin)

  return err
}