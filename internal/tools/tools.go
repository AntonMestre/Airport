package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func WindNumberGenerator() int {
	return rand.Intn(util.MAXIMUM_VALUE_WIND-util.MINIMUM_VALUE_WIND) + util.MINIMUM_VALUE_WIND
}
func TempNumberGenerator() int {
	return rand.Intn(util.MAXIMUM_VALUE_TEMP-util.MINIMUM_VALUE_TEMP) + util.MINIMUM_VALUE_TEMP
}
func PressureNumberGenerator() int {
	return rand.Intn(util.MAXIMUM_VALUE_PRESSURE-util.MINIMUM_VALUE_PRESSURE) + util.MINIMUM_VALUE_PRESSURE
}

func CreateClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func Connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := CreateClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func ReadFile(filename string) []byte {
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

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
