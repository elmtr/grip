package grip

import "github.com/deta/deta-go/service/base"

type AverageMark struct {
	Key        string `json:"key,omitempty"`
	Value     int    `json:"value,omitempty"`
	SubjectKey string `json:"subjectKey,omitempty"`
	StudentKey string `json:"studentKey,omitempty"`
}

func (averageMark *AverageMark) Put() (error) {
	averageMark.Key = GenKey()
	_, err := AverageMarks.Put(averageMark)
	
	return err
}

func GetAverageMarks(query base.Query) ([]AverageMark, error) {
	var averageMarks []AverageMark

	_, err := AverageMarks.Fetch(&base.FetchInput{
		Q: query,
		Dest: &averageMarks,
	})

	return averageMarks, err
}