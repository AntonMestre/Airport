package main

import (
	"context"
	"fmt"
	"tools"
	"util"
	"strings"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Connection to database (MongoDb Cloud)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Airport:Airport@cluster0.0c6je.mongodb.net/AirportDataBase?retryWrites=true&w=majority"))
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Callback function called when a message is recieved from the broker
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Reception d'un message  - %s\n", msg.Payload())

		var collection *mongo.Collection

		if msg.Topic() == util.TOPIC_WIND {
			collection = dbClient.Database("AirportDataBase").Collection("Wind")
		} else if msg.Topic() == util.TOPIC_TEMP {
			collection = dbClient.Database("AirportDataBase").Collection("Temp")
		} else {
			collection = dbClient.Database("AirportDataBase").Collection("Pressure")
		}


		fmt.Println(msg.Payload())
		value := msg.Payload()
		data := strings.Split(string(value),"|")

		fmt.Println(data)
		strconv.Atoi("15256545")

		res, err := collection.InsertOne(ctx, bson.D{{"idCaptor", data[0]}, {"iATA", data[1]}, {"value", data[3]}, {"pickingDate", data[4]}})
		fmt.Printf("res  - %s\n", res)
		fmt.Printf("err  - %s\n", err)
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
