package system

import (
	"time"
	"url_shortener/internal/domain"
)

type clock struct{}

func NewClock() domain.Clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now()
}
