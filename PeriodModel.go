package grip

import (
	"github.com/deta/deta-go/service/base"
)

type Period struct {
  Key string `json:"key"`

  Day       int    `json:"day"`
  Interval  int    `json:"interval"`

  // modifiable
  Room      string `json:"room"`
  
  // indexable
  Subject Subject   `json:"subject"`
}

func (period *Period) Put() (error) {
	if (period.Key == "") {
		period.Key = GenKey()
	}
	
	_, err := Periods.Put(period)

	return err
}

func GetPeriod(query base.Query) (Period, error) {
  var periods []Period

  _, err := Periods.Fetch(&base.FetchInput {
    Q: query,
    Dest: &periods,
    Limit: 1,
  })

  if len(periods) > 0 {
    return periods[0], err 
  } else {
    return Period {}, err
  }
}

func GetPeriods(query base.Query) ([]Period, error) {
  var periods []Period

  _, err := Periods.Fetch(&base.FetchInput {
    Q: query,
    Dest: &periods,
  })

  return periods, err
}

func UpdatePeriod(key string, subject Subject, room string) (error) {
  var period Period

  err := Periods.Update(key, base.Updates {
    "room": room,
    "subject": subject,
  },)

  period.Room = room
  period.Subject = subject

  return err
}

func DeletePeriod(key string) (error) {
  err := Periods.Delete(key)

  return err
}