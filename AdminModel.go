package grip

import (
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
  Key string `json:"key"`
  Email string `json:"email"`
  Password string `json:"password"`
}

type AdminClaims struct {
  Key string `json:"key"`
  Email string `json:"email"`
  Password string `json:"password"`

  jwt.StandardClaims
}

func GenAdminToken(admin Admin) (string, error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &AdminClaims {
    Key: admin.Key,
    Email: admin.Email,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
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