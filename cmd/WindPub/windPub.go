package main

import (
	"fmt"
	"time"
	"tools"
	"util"
)

func main() {

	client := tools.Connect(util.HOST, util.CLIENT_WIND_PUB)

	for {
		client.Publish(util.TOPIC_WIND, 2, false, fmt.Sprintf("%d | %s | %s | %d | %s",
			util.ID_CAPTOR_WIND, util.IATA_CODE, "Wind speed", tools.WindNumberGenerator(), time.Now()))
		time.Sleep(time.Second * 10)
	}

}
