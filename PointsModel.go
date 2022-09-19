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
  var points []Points

  _, err := PointsBase.Fetch(&base.FetchInput {
    Q: query,
    Dest: &points,
    Limit: 1,
  })

  return points[0], err
}

func ModifyPoints(key string, amount int) (error) {

  err := PointsBase.Update(key, base.Updates {
    "value": PointsBase.Util.Increment(amount),
  })

  return err
}

func IncreasePoints(key string) (error) {
  return ModifyPoints(key, 1)
}

func DecreasePoints(key string) (error) {
  return ModifyPoints(key, -1)
}