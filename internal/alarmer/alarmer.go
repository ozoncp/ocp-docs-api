package alarmer

import (
	"time"
)

type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

type alarmer struct {
	timeout time.Duration
	alarms  chan struct{}
	done    chan struct{}
}

func New(timeout time.Duration) Alarmer {
	alarms := make(chan struct{})
	done := make(chan struct{})
	return &alarmer{
		timeout: timeout,
		alarms:  alarms,
		done:    done,
	}
}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarms
}

func (a *alarmer) Init() {
	go func() {
		timer := time.NewTicker(a.timeout)
		defer func() {
			timer.Stop()
			close(a.alarms)
			close(a.done)
		}()

		for {
			select {
			case <-timer.C:
				a.alarms <- struct{}{}
			case <-a.done:
				return
			}
		}
	}()
}

func (a *alarmer) Close() {
	a.done <- struct{}{}
}
