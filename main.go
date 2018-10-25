package main

import (
	"time"
)

var now = func() time.Time {
	return time.Now()
}

func GetNow() time.Time {
	return now()
}
