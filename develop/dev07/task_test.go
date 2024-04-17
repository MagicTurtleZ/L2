package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Run("Single channel", func(t *testing.T) {
		done := make(chan interface{})
		go func() {
			time.Sleep(100 * time.Millisecond)
			close(done)
		}()
		<-or(done)
	})

	t.Run("Multiple channels", func(t *testing.T) {
		done1 := make(chan interface{})
		done2 := make(chan interface{})
		done3 := make(chan interface{})

		go func() {
			time.Sleep(200 * time.Millisecond)
			close(done1)
		}()

		go func() {
			time.Sleep(100 * time.Millisecond)
			close(done2)
		}()

		go func() {
			time.Sleep(300 * time.Millisecond)
			close(done3)
		}()

		<-or(done1, done2, done3)
	})
}

func TestOr_EmptyChannels(t *testing.T) {
	done := or()
	if done != nil {
		t.Error("Expected nil channel, got non-nil channel")
	}
}