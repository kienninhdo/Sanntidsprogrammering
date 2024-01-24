package main

import (
	//"fmt"
	"time"
)

var timerEndTime float64
var timerActive int 

type Timeval struct{
	Sec int64
	Usec int64
}

func get_wall_time() float64 {
	var wall_time float64

	currentTime := Timeval{Sec: time.Now().Unix(), Usec: int64(time.Now().UnixNano() / 1000)}
	wall_time = float64(currentTime.Sec) + float64(currentTime.Usec)*0.000001
	return wall_time
}

func timer_start(duration float64) {
	
	timerEndTime := get_wall_time() + duration
	timerActive := 1
}

func timer_stop() {
	timerActive := 0
}

func timer_timedOut() bool {
	
	return float64(timerActive) !=0 && get_wall_time() > timerEndTime
}