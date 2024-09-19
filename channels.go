package mygolib

import (
  "errors"
)

type StopCloseChan chan struct{}

var ErrStopped = errors.New("StopCloseChan stopped")

func IsStopped(ch StopCloseChan) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
