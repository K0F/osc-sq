package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
        "github.com/crgimenes/go-osc"
	"time"
)

func main() {
        mod := flag.Int("m", 4, "beats per bar")
	bpm := flag.Int("b", 110, "beats per minute")
	port := flag.Int("p", 10000, "port to send OSC messages")
	
        flag.Parse()

	start := time.Now()

	beatNo := 0

        client := osc.NewClient("127.0.0.1", *port)
        client2 := osc.NewClient("127.0.0.1", *port+1)

	for {
		t := time.Now()
		elapsed := t.Sub(start)

		if beatNo == 0 {
			color.Green("%v beat: %s \n", beatNo, elapsed)

		} else {
			fmt.Printf("%v beat: %s \n", beatNo, elapsed)
		}


    go func(value int){
      msg := osc.NewMessage("/osc/timer")
      msg.Append(int32(value))
      client.Send(msg)
      client2.Send(msg)
    }(beatNo)
		beatNo = beatNo + 1
		beatNo = beatNo % *mod
		time.Sleep(time.Duration(60000 / *bpm) * time.Millisecond)
	}

}

