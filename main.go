package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
        "github.com/crgimenes/go-osc"
	"time"
)

func main() {

	delay := flag.Int("b", 110, "beats per minute")
	flag.Parse()

	start := time.Now()

	beatNo := 0

        client := osc.NewClient("localhost", 10000)

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
    }(beatNo)
		beatNo = beatNo + 1
		beatNo = beatNo % 4
		time.Sleep(time.Duration(60000 / *delay) * time.Millisecond)
	}

}

