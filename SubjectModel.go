package grip

import (
	"github.com/deta/deta-go/service/base"
)

type Subject struct {
	Key 	string `json:"key"`
	Name  string `json:"name" `
	Grade Grade  `json:"grade"`
}

func (subject *Subject) Put() (error) {
	subject.Key = GenKey()
	_, err := Subjects.Put(subject)
	
	return err
}

func GetSubject(query base.Query) (Subject, error) {
	var subjects []Subject

	_, err := Subjects.Fetch(&base.FetchInput{
		Q: query,
		Dest: &subjects,
	})

	return subjects[0], err
}

func GetSubjects(query base.Query) ([]Subject, error) {
	var subjects []Subject

	_, err := Subjects.Fetch(&base.FetchInput{
		Q: query,
		Dest: &subjects,
	})

	return subjects, err
}
