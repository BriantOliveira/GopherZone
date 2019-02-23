package clock

import "fmt"

type Clocks int

func (clock Clocks) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}

func (clock Clocks) Add(minutes int) Clocks {
	time := (clock + Clocks(minutes)) % 1440
	if time < 0 {
		time += 1440
	}
	return Clocks(time)
}

func (clock Clocks) Subtract(minutes int) Clocks {
	return clock.Add(-minutes)
}

func New(hour int, minute int) Clocks {
	time := (hour*60 + minute) % 1440
	if time < 0 {
		time += 1440
	}
	return Clocks(time)
}
