package grip

import (
	"github.com/deta/deta-go/service/base"
)

type Truancy struct {
	Key        string `json:"key"`
	Motivated  bool   `json:"motivated"`
	DateDay    int `json:"dateDay"`
	DateMonth  int `json:"dateMonth"`
	SubjectKey string `json:"subjectKey"`
	StudentKey string `json:"studentKey"`
}

func (truancy *Truancy) Put() (error) {
	truancy.Key = GenKey()
	_, err := Truancies.Put(truancy)
	
	return err
}

func GetTruancies(query base.Query) ([]Truancy, error) {
	var truancies []Truancy

	_, err := Truancies.Fetch(&base.FetchInput{
		Q: query,
		Dest: &truancies,
	})

	return truancies, err
}

func MotivateTruancy(key string) (error) {
	err := Truancies.Update(key, 
		base.Updates {
			"motivated": true,
		},
	)
	return err
}