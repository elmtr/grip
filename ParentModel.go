package grip

import (
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/dgrijalva/jwt-go"
)

type Parent struct {
  Key       string          `json:"key"`
  LastName  string          `json:"lastName"`
  FirstName string          `json:"firstName"`
  Phone     string          `json:"phone"`
  Students  []ParentStudent `json:"students"`
  Password  string          `json:"password"`
  Passcode  string          `json:"passcode"`
}

type ParentStudent struct {
  ID        string `json:"id"`
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
}

type ParentClaims struct {
  Key       string          `json:"key"`
  LastName  string          `json:"lastName"`
  FirstName string          `json:"firstName"`
  Phone     string          `json:"phone"`
  Students  []ParentStudent `json:"students"`

  jwt.StandardClaims
}

func GenParentToken(parent Parent) (string, error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &ParentClaims {
    Key: parent.Key,
    FirstName: parent.FirstName,
    LastName: parent.LastName,
    Phone: parent.Phone,
    Students: parent.Students,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
}

func (parent *Parent) Put() (error) {
  parent.Key = GenKey()
  parent.Students = []ParentStudent {}
  
  _, err := Parents.Put(parent)
  return err
}

func GetParent(query base.Query) (Parent, error) {
  var parents []Parent

  _, err := Parents.Fetch(&base.FetchInput {
    Q: query,
    Dest: &parents,
    Limit: 1,
  })

  return parents[0], err
}

func AddParentStudent(key string, students []ParentStudent, student ParentStudent) ([]ParentStudent, error) {
  students = append(students, student)

  err := Teachers.Update(key, base.Updates {
    "students": students,
  })

  return students, err
}

func UpdateParent(key string, updates base.Updates) (error) {
  err := Teachers.Update(key, updates)

  return err
}
