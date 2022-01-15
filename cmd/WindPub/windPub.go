package main

import (
	"encoding/json"
	"fmt"
	"time"
	"tools"
	"util"
)

func main() {

	values := tools.ReadFile("config-wind")
	var config util.Config
	json.Unmarshal(values, &config)

	client := tools.Connect(config.Host, config.ClientId)

	for {
		client.Publish(config.Broker, config.Qoslevel, false, fmt.Sprintf("%d|%s|%s|%f|%s",
			config.IdSensor, config.IataCode, config.Nature, tools.FetchData(config.CityName).Wind_speed, time.Now()))
		time.Sleep(time.Second * 10)
	}

}
