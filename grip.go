package grip

import (
	"context"
	"fmt"

	// deta
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"

	// redis
	"github.com/go-redis/redis/v8"
)

// administrative bases
var Schools *base.Base
var Grades *base.Base
var Subjects *base.Base
var Periods *base.Base

// main bases
var Marks *base.Base
var Truancies *base.Base
var DraftMarks *base.Base
var PointsBase *base.Base
var AverageMarks *base.Base

// accounts bases
var Teachers *base.Base
var Students *base.Base
var Parents *base.Base
var Admins *base.Base

var ctx = context.Background()
var RDB *redis.Client

// initializing database
func InitDB(DetaKey string) {
    // initializing deta
  d, err := deta.New(deta.WithProjectKey(DetaKey))
  if err != nil {
    fmt.Println("failed to init new Deta instance:" , err)
  }

  // adminstrative bases
  Schools, _ = base.New(d, "schools")
  Grades, _ = base.New(d, "grades")
  Subjects, _ = base.New(d, "subjects")
  Periods, _ = base.New(d, "periods")

  // main bases
  Marks, _ = base.New(d, "marks")
  Truancies, _ = base.New(d, "truancies")
  DraftMarks, _ = base.New(d, "draftMarks")
  PointsBase, _ = base.New(d, "points")
  AverageMarks, _ = base.New(d, "averageMarks")

  // accounts bases
  Teachers, _ = base.New(d, "teachers")
  Students, _ = base.New(d, "students")
  Parents, _ = base.New(d, "parents")
  Admins, _ = base.New(d, "admins")

  fmt.Println("connected to deta")
}

func InitCache(RedisOptions *redis.Options) {
  RDB = redis.NewClient(RedisOptions)
  
  pong, _ := RDB.Ping(context.Background()).Result()
  if pong == "PONG" {
    fmt.Println("connected to redis")
  } else {
    fmt.Println("not connected to redis")
  }
}

func Set(key string, value string) error {
  err := RDB.Set(ctx, key, value, 0).Err()

  return err
}

func Get(key string) (string, error) {
  val, err := RDB.Get(ctx, key).Result()

  return val, err
}

func Del(key string) error {
  _, err := RDB.Del(ctx, key).Result()

  return err
}
