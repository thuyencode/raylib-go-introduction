package timer

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	Duration     float64
	StartTime    float64
	IsActive     bool
	ShouldRepeat bool
	callback     func()
}

func NewTimer(duration float64, repeat bool, callback func()) *Timer {
	return &Timer{
		Duration:     duration,
		StartTime:    0,
		IsActive:     false,
		ShouldRepeat: repeat,
		callback:     callback,
	}
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
		if timer.callback != nil && timer.StartTime > 0 {
			timer.callback()
		}

		timer.Deactivate()
	}
}
