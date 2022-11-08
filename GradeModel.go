package grip

import (
	"github.com/deta/deta-go/service/base"
)

type Grade struct {
	Key     		string `json:"key"`
	SchoolKey   string `json:"schoolKey"`
	YearFrom    int    `json:"yearFrom"`
	YearTo      int    `json:"yearTo"`
	GradeNumber int    `json:"gradeNumber"`
	GradeLetter string `json:"gradeLetter"`
	// Intervals   string `json:"intervals"`
}


func (grade *Grade) Put() (error) {
	grade.Key = GenKey()
	_, err := Grades.Put(grade)
	
	return err
}

func GetGrade(query base.Query) (Grade, error) {
	var grades []Grade

	_, err := Grades.Fetch(&base.FetchInput{
		Q: query,
		Dest: &grades,
		Limit: 1,
	})

	if len(grades) > 0 {
		return grades[0], err
	} else {
		return Grade {}, err
	}

}

func GetGrades(query base.Query) ([]Grade, error) {
	var grades []Grade

	_, err := Grades.Fetch(&base.FetchInput{
		Q: query,
		Dest: &grades,
	})

	return grades, err
}
