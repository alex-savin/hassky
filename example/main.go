package main

import (
	"fmt"

	"github.com/alex-savin/hassky"
)

func main() {
	h, _ := hassky.New("[YOUR-IP-ADDREDD]:[PORT]", "[YOUR-HASS-API-TOKEN]", false, "[debug]")

	fmt.Println(h.R().GetConfig().Exec())

	h.R().Ping().Exec()

	h.R().GetState("light.kitchen_led_01").Exec()

	h.R().CallServiceWS("light.toggle", "light.kitchen_01").
		SetParameters(map[string]interface{}{"transition": 15, "brightness_pct": 100}).
		Exec()

	h.R().SetState("sensor.api_request", "on").
		SetAttributes(map[string]interface{}{"device_class": "golang"}).
		Exec()

	h.R().CallService("light.toggle", "light.porch_01").
		SetParameters(map[string]interface{}{"transition": 15, "brightness_pct": 100}).
		Exec()

	h.R().Ping().Exec()

	h.R().CallService("fan.toggle", "fan.living_room_bathroom_fan").Exec()

	h.R().GetEventListeners().Exec()

	h.R().Ping().Exec()

	h.R().GetConfigWS().Exec()
	h.R().DiscoveryInfo().Exec()
}
