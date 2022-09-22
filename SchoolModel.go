package grip

import "github.com/deta/deta-go/service/base"

type School struct {
	Key       string   `json:"key"`
	Name      string   `json:"name"`
	Intervals []Interval `json:"intervals"`
}

type Interval struct {
	Number int		 `json:"number"`
	Start  float32 `json:"start"`
	End 	 float32 `json:"end"`
}

func (school *School) Put() (error) {
	school.Key = GenKey()
	_, err := Schools.Put(school)
	
	return err
}

func GetSchool(key string) (School, error) {
	var school School

	err := Schools.Get(key, &school)

	return school, err
}

func GetSchools(query base.Query) ([]School, error) {
	var schools []School

	_, err := Schools.Fetch(&base.FetchInput{
		Q: query,
		Dest: &schools,
	})

	return schools, err
}
