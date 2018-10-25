package main_test

import (
	"testing"
	"time"

	main "github.com/abraithwaite/blahblah"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {

	ts, _ := time.Parse(time.RFC3339, "2000-01-01T00:01Z")

	main.SetNow(func() time.Time {
		return ts
	})

	assert.Equal(t, ts, main.GetNow(), "should be equal")
}
