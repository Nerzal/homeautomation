package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var leds [60]color.RGBA

func main() {
	machine.D2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledStrip := ws2812.New(machine.D2)

	for {
		for i := 0; i < len(leds); i++ {
			leds[i] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		}

		ledStrip.WriteColors(leds[:])
		time.Sleep(500 * time.Millisecond)

		for i := 0; i < len(leds); i++ {
			leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
		}

		ledStrip.WriteColors(leds[:])
		time.Sleep(500 * time.Millisecond)
	}
}
