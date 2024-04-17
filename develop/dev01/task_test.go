package main

import (
	"testing"
	"github.com/beevik/ntp"
)

func TestTime(t *testing.T) {
	ntpAddress := "0.beevik-ntp.pool.ntp.org"

	_, err := ntp.Time(ntpAddress)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
