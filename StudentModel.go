package grip

import (
	"time"

	sj "github.com/brianvoe/sjwt"
	"github.com/deta/deta-go/service/base"
)

type Student struct {
  Key       string         `json:"key"`
  LastName  string         `json:"lastName"`
  FirstName string         `json:"firstName"`
  Phone     string         `json:"phone"`
  Grade  Grade             `json:"grade"`
  Subjects  []ShortSubject `json:"subjects"`
  Password  string         `json:"password"`
  Passcode  string         `json:"passcode"`
}

type ShortSubject struct {
  Key  string `json:"key"`
	Name string `json:"name"`
}

func GenStudentToken(student Student) (string)  {
  info := &Student {
    Key: student.Key,
    FirstName: student.FirstName,
    LastName: student.LastName,
    Phone: student.Phone,
    Grade: student.Grade,
    Subjects: student.Subjects,
  }
  claims, _ := sj.ToClaims(info)
  claims.SetExpiresAt(time.Now().Add(8760 * time.Hour))

  token := claims.Generate(JWTKey)
  return token
}

func ParseStudentToken(token string) (Student, error) {
  hasVerified := sj.Verify(token, JWTKey)

  if !hasVerified {
    return Student {}, nil
  }

  claims, _ := sj.Parse(token)
  err := claims.Validate()
  student := Student {}
  claims.ToStruct(&student)

  return student, err
}

func (student *Student) Put() (error) {
  student.Key = GenKey()
  student.Grade = Grade {}
  student.Subjects = []ShortSubject {}
  
  _, err := Students.Put(student)
  return err
}

func GetStudent(query base.Query) (Student, error) {
  var students []Student

  _, err := Students.Fetch(&base.FetchInput {
    Q: query,
    Dest: &students,
    Limit: 1,
  })

  return students[0], err
}

func GetStudents(query base.Query) ([]Student, error) {
  var students []Student

  _, err := Students.Fetch(&base.FetchInput {
    Q: query,
    Dest: &students,
    Limit: 1,
  })

  return students, err
}

func AddStudentSubject(key string, subjects []ShortSubject, subject ShortSubject) ([]ShortSubject, error) {
  subjects = append(subjects, subject)

  err := Teachers.Update(key, base.Updates {
    "subjects": subjects,
  })

  return subjects, err
}

func RemoveStudentSubject(key string, subjects []ShortSubject, oldSubject ShortSubject) ([]ShortSubject, error) {
  var newSubjects []ShortSubject 
  for _, subject := range subjects {
    if (subject.Key != oldSubject.Key) {
      newSubjects = append(newSubjects, subject)
    }
  }

  err := Students.Update(key, base.Updates {
    "subjects": newSubjects,
  })

  return newSubjects, err
}

func UpdateStudentGrade(key string, grade Grade) (Student, error) {
  var student Student 

  err := Students.Update(key, base.Updates {
    "grade": grade,
  })

  student.Grade = grade

  return student, err
}

func StudentSetup(key string, grade Grade) (Student, error) {
  // getting the grade
  grade, err := GetGrade(
    base.Query {
      {"gradeNumber": grade.GradeNumber},
      {"gradeLetter": grade.GradeLetter},
    },
  )
  if err != nil {
    return Student {}, err
  }
  
  // getting the subjects
  subjects, err := GetSubjects(
    base.Query {
      {"grade.gradeLetter": grade.GradeLetter},
      {"grade.gradeNumber": grade.GradeNumber},
    },
  )
  if err != nil {
    return Student {}, err
  }

  // transforming subjects to short subjects
  var shortSubjects []ShortSubject
  for _, subject := range subjects {
    shortSubjects = append(
      shortSubjects,
      ShortSubject {
        Key: subject.Key,
        Name: subject.Name,
      },
    )
  }

  var student Student

  // updating the student
  err = Students.Update(key, base.Updates {
    "grade": grade,
    "subjects": shortSubjects,
  },)
  if err != nil {
    return Student {}, err
  }

  // getting the student
  err = Students.Get(key, &student)

  return student, err  
}


func UpdateStudent(key string, updates base.Updates) (error) {
  err := Students.Update(key, updates)

  return err
}
