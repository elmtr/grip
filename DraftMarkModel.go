package grip

import (
	"github.com/deta/deta-go/service/base"
)

type DraftMark struct {
  Key        string `json:"key"`
	Value      int    `json:"value"`
	DateDay    int 		`json:"dateDay"`
	DateMonth  int 		`json:"dateMonth"`
	SubjectKey string `json:"subjectKey"`
	StudentKey string `json:"studentKey"`
}

func (draftMark *DraftMark) Put() (error) {
	draftMark.Key = GenKey()
	_, err := DraftMarks.Put(draftMark)
	
	return err
}

func GetDraftMark(key string) (DraftMark, error) {
	var draftMark DraftMark

	err := DraftMarks.Get(key, &draftMark)

	return draftMark, err
}

func GetDraftMarks(query base.Query) ([]DraftMark, error) {
	var draftMarks []DraftMark

	_, err := DraftMarks.Fetch(&base.FetchInput{
		Q: query,
		Dest: &draftMarks,
	})

	return draftMarks, err
}

func (draftMark *DraftMark) Update() (error) {
  _, err := DraftMarks.Put(draftMark)
	
	return err
}

func DefinitivateDraftMark(key string) (Mark, error) {
	var draftMark DraftMark

	err := DraftMarks.Get(key, &draftMark)
	if err != nil {
		return Mark {}, err
	}
	err = DraftMarks.Delete(key)
	if err != nil {
		return Mark {}, err
	}

  mark := Mark(draftMark)
	err = mark.Put()

  return mark, err
}
