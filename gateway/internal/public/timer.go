package public

import (
	"time"
)

var startTime time.Time

func TimerStart() {
	startTime = time.Now()
}

func TimerEnd() string {
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	return elapsed.String()
}
