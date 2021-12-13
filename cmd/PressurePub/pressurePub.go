package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"tools"
	"util"
)

type ConfigPressure struct {
	Nature   string `json:"nature"`
	IataCode string `json:"iatacode"`
	IdSensor byte   `json:"idsensor"`
	Broker   string `json:"broker"`
	Port     int    `json:"port"`
	Qoslevel byte   `json:"qoslevel"`
	ClientId string `json:"clientId"`
}

func main() {

	client := tools.Connect(util.HOST, util.CLIENT_PRESSURE_PUB)

	values := readfile("config-pressure")
	var config ConfigPressure
	json.Unmarshal(values, &config)

	for {
		client.Publish(config.Broker, config.Qoslevel, false, fmt.Sprintf("%d | %s | %s |%d | %s",
			config.IdSensor, config.IataCode, config.Nature, tools.PressureNumberGenerator(), time.Now()))
		time.Sleep(time.Second * 10)
	}

}

func readfile(filename string) []byte {
	jsonfile, err := os.Open("../../internal/util/" + filename + ".json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonfile.Close()
	bytevalue, err := ioutil.ReadFile(jsonfile.Name())
	if err != nil {
		fmt.Println(err)
	}
	return bytevalue
}
