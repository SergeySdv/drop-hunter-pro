package lib

import (
	"math/rand"
	"time"
)

var s = rand.NewSource(time.Now().UnixNano())

func Contains(s []string, t string) bool {
	for _, ss := range s {
		if ss == t {
			return true
		}
	}
	return false
}
