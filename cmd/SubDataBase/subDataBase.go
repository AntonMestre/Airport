package main

import (
	"fmt"
	"tools"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	client := tools.Connect(util.HOST, util.CLIENT_DATABASE_SUB)
	tokenWind := client.Subscribe(util.TOPIC_WIND, 2, messagePubHandler)
	tokenTemp := client.Subscribe(util.TOPIC_TEMP, 2, messagePubHandler)
	tokenPressure := client.Subscribe(util.TOPIC_PRESSURE, 2, messagePubHandler)

	for {
		tokenWind.Wait()
		tokenTemp.Wait()
		tokenPressure.Wait()
	}

}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("%s\n", msg.Payload())
}
