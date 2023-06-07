package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/crgimenes/go-osc"
	"github.com/fatih/color"
)

func main() {
	mod := flag.Int("m", 4, "beats per bar")
	bpm := flag.Int("b", 110, "beats per minute")
	port := flag.Int("p", 10000, "port to send OSC messages")

	flag.Parse()

	start := time.Now()

	beatNo, totalNo := 0, 0

	client := osc.NewClient("127.0.0.1", *port)
	client2 := osc.NewClient("127.0.0.1", *port+1)

	dur := time.Duration(60000 / *bpm) * time.Millisecond
	var drift time.Duration

	for {

		t := time.Now()
		elapsed := t.Sub(start)

		// time.Sleep() is slightly drifting over time, correction needed here
		drift = time.Duration(elapsed.Milliseconds()%dur.Milliseconds()) * time.Millisecond

		if beatNo == 0 {
			color.Green("%v beat: %v drift: %v \n", beatNo, elapsed, drift)

		} else {
			fmt.Printf("%v beat: %v drift: %v \n", beatNo, elapsed, drift)
		}

		msg := osc.NewMessage("/osc/timer")
		msg.Append(int32(beatNo))
		client.Send(msg)
		client2.Send(msg)

		totalNo = totalNo + 1
		beatNo = beatNo + 1
		beatNo = beatNo % *mod

		// calculate drift correction
		ms := time.Duration(dur.Milliseconds()-drift.Milliseconds()) * time.Millisecond
		time.Sleep(ms)
	}

}
