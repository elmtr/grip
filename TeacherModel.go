package grip

import (
	"time"

	"github.com/deta/deta-go/service/base"
	"github.com/dgrijalva/jwt-go"
)

type Teacher struct {
  Key       string    `json:"key"`
  LastName  string    `json:"lastName"`
  FirstName string    `json:"firstName"`
  Phone     string    `json:"phone"`
  Homeroom  Grade     `json:"homeroom"`
  Subjects  []Subject `json:"subjects"`
  Password  string    `json:"password"`
  Passcode  string    `json:"passcode"`
}

type TeacherClaims struct {
  Key       string    `json:"key"`
  LastName  string    `json:"lastName"`
  FirstName string    `json:"firstName"`
  Phone     string    `json:"phone"`
  Homeroom  Grade     `json:"homeroom"`
  Subjects  []Subject `json:"subjects"`
  
  jwt.StandardClaims
}

func GenTeacherToken(teacher Teacher) (string, error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &TeacherClaims {
    Key: teacher.Key,
    FirstName: teacher.FirstName,
    LastName: teacher.LastName,
    Phone: teacher.Phone,
    Homeroom: teacher.Homeroom,
    Subjects: teacher.Subjects,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
}

func (teacher *Teacher) Put() (error) {
  teacher.Key = GenKey()
  teacher.Homeroom = Grade {}
  teacher.Subjects = []Subject {}
  
  _, err := Teachers.Put(teacher)
  return err
}

func GetTeacher(query base.Query) (Teacher, error) {
  var teachers []Teacher

  _, err := Teachers.Fetch(&base.FetchInput {
    Q: query,
    Dest: &teachers,
    Limit: 1,
  })

  return teachers[0], err
}

func AddTeacherSubject(key string, subjects []Subject, subject Subject) ([]Subject, error) {
  subjects = append(subjects, subject)

  err := Teachers.Update(key, base.Updates {
    "subjects": subjects,
  })

  return subjects, err
}

func RemoveTeacherSubject(key string, subjects []Subject, oldSubject Subject) ([]Subject, error) {
  var newSubjects []Subject 
  for _, subject := range subjects {
    if (subject.Key != oldSubject.Key) {
      newSubjects = append(newSubjects, subject)
    }
  }

  err := Teachers.Update(key, base.Updates {
    "subjects": newSubjects,
  })

  return newSubjects, err
}

func UpdateTeacherHomeroom(key string, homeroom Grade) (Teacher, error) {
  var teacher Teacher 

  err := Teachers.Update(key, base.Updates {
    "homeroom": homeroom,
  })

  teacher.Homeroom = homeroom

  return teacher, err
}

func UpdateTeacher(key string, updates base.Updates) (error) {
  err := Teachers.Update(key, updates)

  return err
}
