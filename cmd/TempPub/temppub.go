package main

import (
	"encoding/json"
	"fmt"
	"time"
	"tools"
	"util"
)

func main() {

	values := tools.ReadFile("config-temp")
	var config util.Config
	json.Unmarshal(values, &config)

	client := tools.Connect(config.Host, config.ClientId)

	for {
		client.Publish(config.Broker, config.Qoslevel, false, fmt.Sprintf("%d|%s|%s|%f|%s",
			config.IdSensor, config.IataCode, config.Nature, tools.FetchData(config.CityName).Temp, time.Now().UTC().Format(time.RFC3339)))
		time.Sleep(time.Second * 10)
	}

}
