package main

import (
	"encoding/json"
	"fmt"
	"time"
	"tools"
	"util"
)

func main() {

	values := tools.ReadFile("config-pressure")
	var config util.Config
	json.Unmarshal(values, &config)

	client := tools.Connect(config.Host, config.ClientId)

	for {
		client.Publish(config.Broker, config.Qoslevel, false, fmt.Sprintf("%d | %s | %s |%d | %s",
			config.IdSensor, config.IataCode, config.Nature, tools.PressureNumberGenerator(), time.Now()))
		time.Sleep(time.Second * util.NB_SEC)
	}
}
