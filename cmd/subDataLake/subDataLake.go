package main

import (
	"fmt"
	"os"
	"strings"
	"tools"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	// Callback function called when a message is recieved from the broker
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Reception d'un message  - %s\n", msg.Payload())

		//fmt.Println(msg.Payload())
		value := msg.Payload()
		data := strings.Split(string(value), "|")

		var sensor string

		if msg.Topic() == util.TOPIC_WIND {
			sensor = "Wind"
		} else if msg.Topic() == util.TOPIC_TEMP {
			sensor = "Temperature"
		} else {
			sensor = "Pressure"
		}

		dateSplit := strings.Split(data[4], "-")
		date := []string{dateSplit[0], dateSplit[1], dateSplit[2]}
		csvFileName := data[1] + "-" + strings.Join(date, "-") + "-" + sensor + ".csv"
		filePath := "../../datalake/" + csvFileName

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		content := strings.Replace(string(value), "|", ";", -1) + "\n"

		if _, err = file.WriteString(content); err != nil {
			panic(err)
		}
	}

	// Connecting to the broker in subscriber mode
	client := tools.Connect(util.HOST, util.CLIENT_DATABASE_SUB)
	tokenWind := client.Subscribe(util.TOPIC_WIND, 2, messagePubHandler)
	tokenTemp := client.Subscribe(util.TOPIC_TEMP, 2, messagePubHandler)
	tokenPressure := client.Subscribe(util.TOPIC_PRESSURE, 2, messagePubHandler)

	// Waiting for message from broker
	for {
		tokenWind.Wait()
		tokenTemp.Wait()
		tokenPressure.Wait()
	}

}
