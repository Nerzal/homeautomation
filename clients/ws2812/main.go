package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var leds [60]color.RGBA

func main() {
	machine.D2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledStrip := ws2812.New(machine.D2)

	rg := false

	for {
		rg = !rg
		for i := range leds {
			rg = !rg
			if rg {
				// Alpha channel is not supported by WS2812 so we leave it out
				leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			} else {
				leds[i] = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			}
		}

		mask := interrupt.Disable()
		ledStrip.WriteColors(leds[:])
		interrupt.Restore(mask)

		time.Sleep(100 * time.Millisecond)
	}
}
