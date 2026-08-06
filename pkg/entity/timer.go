package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	Duration     float64
	StartTime    float64
	ShouldRepeat bool
	IsActive     bool
	Callback     func()
}

func NewTimer(duration float64, repeat bool, autostart bool, callback func()) *Timer {
	timer := Timer{
		Duration:     duration,
		StartTime:    0,
		ShouldRepeat: repeat,
		IsActive:     false,
		Callback:     callback,
	}

	if autostart {
		timer.Activate()
	}

	return &timer
}

func (timer *Timer) Activate() {
	timer.IsActive = true
	timer.StartTime = rl.GetTime()
}

func (timer *Timer) Deactivate() {
	timer.IsActive = false
	timer.StartTime = 0

	if timer.ShouldRepeat {
		timer.Activate()
	}
}

func (timer *Timer) Update() {
	if !timer.IsActive {
		return
	}

	if rl.GetTime()-timer.StartTime >= timer.Duration {
		if timer.Callback != nil && timer.StartTime > 0 {
			timer.Callback()
		}

		timer.Deactivate()
	}
}
