package alarmer

import (
	"fmt"
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
		alarms: alarms,
		done: done,
	}
}

func (a * alarmer) Alarm() <-chan struct{} {
	return a.alarms
}

func (a * alarmer) Init() {
	go func () {
		timer := time.NewTicker(a.timeout)
		defer timer.Stop()
		for {
			select {
				case <- timer.C:
					fmt.Println("Tick")
					a.alarms <- struct{}{}
				case <-a.done:
					close(a.alarms)
					close(a.done)
					fmt.Println("close alarmer")
					return

			}
		}
	} ()
}

func (a * alarmer) Close() {
	fmt.Println("CLOSE!!!")
	a.done <- struct{}{}
}