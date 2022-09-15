package grip

import (
	"github.com/deta/deta-go/service/base"
)

type Truancy struct {
	Key        string `json:"key"`
	Motivated bool   `json:"motivated"`
	DateDay   string `json:"dateDay"`
	DateMonth string `json:"dateMonth"`
	SubjectKey string `json:"subjectKey"`
	StudentKey string `json:"studentKey"`
}

func (truancy *Truancy) Put() (error) {
	truancy.Key = GenKey()
	_, err := Truancies.Put(truancy)
	
	return err
}

func GetTruancies(query base.Query) ([]Mark, error) {
	var marks []Mark

	_, err := Marks.Fetch(&base.FetchInput{
		Q: query,
		Dest: &marks,
	})

	return marks, err
}

func MotivateTruancy(key string) (error) {
	err := Truancies.Update(key, 
		base.Updates {
			"motivated": true,
		},
	)
	return err
}