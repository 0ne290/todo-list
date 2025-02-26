package internal

import "time"

type ITimeProvider interface {
	Now() time.Time
}

type TimeProvider struct { }

func NewTimeProvider() *TimeProvider {
	return &TimeProvider{}
}

func (timeProvider *TimeProvider) Now() time.Time {
	return time.Now()
}