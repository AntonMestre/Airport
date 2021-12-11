package main

import (
	"fmt"
	"time"
	"tools"
	"util"
)

func main() {

	client := tools.Connect(util.HOST, util.CLIENT_TEMP_PUB)

	for {
		client.Publish(util.TOPIC_TEMP, 2, false, fmt.Sprintf("%d | %s | %s | %d | %s",
			util.ID_CAPTOR_TEMP, util.IATA_CODE, "Temperature", tools.TempNumberGenerator(), time.Now()))
		time.Sleep(time.Second * 10)
	}

}
