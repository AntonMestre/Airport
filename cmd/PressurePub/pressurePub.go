package main

import (
	"fmt"
	"time"
	"tools"
	"util"
)

func main() {

	client := tools.Connect(util.HOST, util.CLIENT_PRESSURE_PUB)

	for {
		client.Publish(util.TOPIC_PRESSURE, 2, false, fmt.Sprintf("%d | %s | %s |%d | %s",
			util.ID_CAPTOR_PRESSURE, util.IATA_CODE, "Atmospheric Pressure", tools.PressureNumberGenerator(), time.Now()))
		time.Sleep(time.Second * 10)
	}

}
