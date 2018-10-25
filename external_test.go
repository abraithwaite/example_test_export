package main

import "time"

func SetNow(f func() time.Time) {
	now = f
}
