package grip

import "github.com/deta/deta-go/service/base"

type Mark struct {
	Key        string `json:"key"`
	Value      int    `json:"value"`
	DateDay    int 		`json:"dateDay"`
	DateMonth  int 		`json:"dateMonth"`
	SubjectKey string `json:"subjectKey"`
	StudentKey string `json:"studentKey"`
}

func GetMarks(query base.Query) ([]Mark, error) {
	var marks []Mark

	_, err := Marks.Fetch(&base.FetchInput{
		Q: query,
		Dest: &marks,
	})

	return marks, err
}

func (mark *Mark) Put() (error) {
	mark.Key = GenKey()
	_, err := Marks.Put(mark)
	
	return err
}