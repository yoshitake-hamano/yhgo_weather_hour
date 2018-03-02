package main

import (
	"fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"
	"time"
)

func main() {
	// START CONST OMIT
	const (
		GPIO_PIN           = 18
		LED_COUNT          = 3
		DEFAULT_BRIGHTNESS = 255
		COLOR_RED          = 0x00FF00
		COLOR_BLACK        = 0x000000
	)
	// END CONST OMIT
	// START EXECUTE OMIT
	if err := ws2811.Init(GPIO_PIN, LED_COUNT, DEFAULT_BRIGHTNESS); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {
		ws2811.SetLed(0, COLOR_RED)
		if err := ws2811.Render(); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)

		ws2811.SetLed(0, COLOR_BLACK)
		if err := ws2811.Render(); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
	// END EXECUTE OMIT
}
