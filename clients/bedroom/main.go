package main

import (
	"image/color"
	"machine"
	"math/rand"
	"runtime/interrupt"
	"time"

	"tinygo.org/x/drivers/net/mqtt"
	"tinygo.org/x/drivers/wifinina"
	"tinygo.org/x/drivers/ws2812"
)

var status = false
var leds [60]color.RGBA
var ledStrip ws2812.Device

func main() {
	machine.D2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledStrip = ws2812.New(machine.D2)

	connectWifi()
	mqttClient := connectMQTT()

	mqttClient.Subscribe("/noobygames/homeautomation/home/bedroom/light/on", 0, turnLightOn)
	mqttClient.Subscribe("/noobygames/homeautomation/home/bedroom/light/stop", 0, stop)
	mqttClient.Subscribe("/noobygames/homeautomation/home/bedroom/light/off", 0, turnLightOff)
	mqttClient.Subscribe("/noobygames/homeautomation/home/bedroom/light/status/request", 0, sendStatus)

	select {}
}

func sendStatus(client mqtt.Client, message mqtt.Message) {
	println("handling turn light on message")
	message.Ack()

	statusString := "off"
	if status {
		statusString = "on"
	}

	go func() {
		client.Publish("/noobygames/homeautomation/home/bedroom/light/status", 1, false, statusString)
	}()
}

func turnLightOn(client mqtt.Client, message mqtt.Message) {
	println("handling turn light on message")

	machine.D4.High()
	status = true
	go animationOne()
	message.Ack()
}

func stop(client mqtt.Client, message mqtt.Message) {
	println("handling turn light off message")

	machine.D4.Low()
	status = false
	message.Ack()
}

func turnLightOff(client mqtt.Client, message mqtt.Message) {
	for i := range leds {
		leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	mask := interrupt.Disable()
	ledStrip.WriteColors(leds[:])
	interrupt.Restore(mask)

	status = false
	message.Ack()
}

func animationOne() {
	br := false

	for {
		if !status {
			return
		}

		br = !br
		for i := range leds {
			if br {
				// Alpha channel is not supported by WS2812 so we leave it out
				leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0xff}
			} else {
				leds[i] = color.RGBA{R: 0xff, G: 0xff, B: 0x00}
			}
		}

		mask := interrupt.Disable()
		ledStrip.WriteColors(leds[:])
		interrupt.Restore(mask)

		time.Sleep(100 * time.Millisecond)
	}
}

func connectMQTT() mqtt.Client {
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://192.168.2.102:1883").
		SetClientID("bedroom" + randomString(5))
	client := mqtt.NewClient(opts)

	println("trying to connect to mqtt broker")

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		printError("could not establish mqtt connection:", token.Error())
	}

	println("successfully connected to mqtt broker")

	return client
}

func connectWifi() {
	err := machine.NINA_SPI.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})
	if err != nil {
		println("could not configure nina spi", err.Error())
	}

	wifi := wifinina.New(machine.NINA_SPI, machine.NINA_CS, machine.NINA_ACK, machine.NINA_GPIO0, machine.NINA_RESETN)

	wifi.Configure()

	time.Sleep(5 * time.Second)

	err = wifi.SetPassphrase("NoobyGames", "IchHasseLangeWlanZugangsDaten1312!")
	if err != nil {
		println("could not wifi credentials", err.Error())
	}

	for {
		status, err := wifi.GetConnectionStatus()
		if err != nil {
			println("could not get connection status", err.Error())
		}

		if status == wifinina.StatusConnected {
			println("successfully connected to wifi")
			break
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func printError(message string, err error) {
	for {
		println(message, err.Error())
		time.Sleep(time.Second)
	}
}
