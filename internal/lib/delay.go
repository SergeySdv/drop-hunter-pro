package lib

import (
	"time"

	"go.uber.org/zap"
)

func DelayBetween(min, max time.Duration, l *zap.SugaredLogger, op string) {
	delay := RandDurationRange(min, max)
	l.Info("zzz-zz..zz...zz... [" + op + "]: " + delay.String())
	time.Sleep(delay)
}
