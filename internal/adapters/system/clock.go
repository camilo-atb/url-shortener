package system

import "time"

type clock struct{}

func NewClock() *clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now()
}
