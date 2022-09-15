package grip

import (
	"github.com/deta/deta-go/service/base"
)

type Points struct {
  Key        string `json:"key"`
	Value     int    `json:"value"`
	SubjectKey string `json:"subjectKey"`
	StudentKey string `json:"studentKey"`
}

func (points *Points) Put() (error) {
  points.Key = GenKey()
	_, err := PointsBase.Put(points)
	
	return err
}

func GetPoints(query base.Query) (Points, error) {
  var points Points

  _, err := PointsBase.Fetch(&base.FetchInput {
    Q: query,
    Dest: &points,
    Limit: 1,
  })

  return points, err
}

func ModifyPoints(key string, amount int) (Points, error) {
  var points Points

  err := PointsBase.Update(key, base.Updates {
    "value": PointsBase.Util.Increment(amount),
  })

  points.Value += amount
  return points, err
}

func IncreasePoints(key string) (Points, error) {
  return ModifyPoints(key, 1)
}

func DecreasePoints(key string) (Points, error) {
  return ModifyPoints(key, -1)
}