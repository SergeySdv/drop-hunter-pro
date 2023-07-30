package lib

import (
	"math/rand"
	"time"
)

func RandFloat(v float64, delta int) float64 {

	mind := (100.0 - float64(delta)) / 100.0

	min := v * mind
	max := v * ((100.0 + float64(delta)) / 100.0)

	randomFloatInRange := min + rand.New(s).Float64()*(max-min)
	return Round(randomFloatInRange)
}

func RandFloatRange(min, max float64) float64 {
	randomFloatInRange := min + rand.New(s).Float64()*(max-min)
	return Round(randomFloatInRange)
}

func RandDurationRange(min, max time.Duration) time.Duration {
	mini := int64(min)
	maxi := int64(max)
	if maxi-mini <= 0 {
		return time.Duration(mini)
	}

	d := mini + rand.New(s).Int63n(maxi-mini)
	delay := time.Duration(d)
	return delay
}
