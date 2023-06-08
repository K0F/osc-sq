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
	bpm := flag.Float64("b", 110.0, "beats per minute")
	port := flag.Int("p", 10000, "port to send OSC messages")

	flag.Parse()

	start := time.Now()

	beatNo, barNo, totalNo := 0, 0, 0

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
			color.Green("%04d %04d %04d T %v\n", totalNo, barNo, beatNo, elapsed.Round(time.Duration(1*time.Millisecond)))

		} else {
			fmt.Printf("%04d %04d %04d T %v\n", totalNo, barNo, beatNo, elapsed.Round(time.Duration(1*time.Millisecond)))
		}

		msg := osc.NewMessage("/osc/timer")
		msg.Append(int32(beatNo))
		msg.Append(int32(totalNo))
		msg.Append(int32(*bpm))
		client.Send(msg)
		client2.Send(msg)

		totalNo = totalNo + 1
		beatNo = beatNo + 1

		if beatNo >= *mod {
			beatNo = 0
			barNo = barNo + 1
		}

		// calculate drift correction
		ms := time.Duration(dur.Milliseconds()-drift.Milliseconds()) * time.Millisecond
		time.Sleep(ms)
	}

}
