package grip

import (
	"time"

	sj "github.com/brianvoe/sjwt"
	"github.com/deta/deta-go/service/base"
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

func GenParentToken(parent Parent) (string) {
  info := &Parent {
    Key: parent.Key,
    FirstName: parent.FirstName,
    LastName: parent.LastName,
    Phone: parent.Phone,
    Students: parent.Students,
  }
  claims, _ := sj.ToClaims(info)
  claims.SetExpiresAt(time.Now().Add(8760 * time.Hour))

  token := claims.Generate(JWTKey)
  return token
}

func ParseParentToken(token string) (Parent, error) {
  hasVerified := sj.Verify(token, JWTKey)

  if !hasVerified {
    return Parent {}, nil
  }

  claims, _ := sj.Parse(token)
  err := claims.Validate()
  parent := Parent {}
  claims.ToStruct(&parent)

  return parent, err
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
