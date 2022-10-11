package grip

import (
	"math/rand"
	"time"
)

var Encoding string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

func GenKey() string {
  rand.Seed(time.Now().UnixNano())
  var ID string
  for i := 0; i < 6; i++ {
    ID += string(Encoding[rand.Intn(64)])
  }
  return ID
}