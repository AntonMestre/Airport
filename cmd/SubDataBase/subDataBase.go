package main

import (
	"context"
	"fmt"
	"strings"
	"time"
	"tools"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Connection to database (MongoDb Cloud)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(util.DB_URI))
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

		value := msg.Payload()

		data := strings.Split(string(value), "|")

		// loc, err := time.LoadLocation("Europe/Paris")
		// if err != nil {
		// 	panic(err)
		// }

		res, err := collection.InsertOne(ctx, bson.D{primitive.E{Key: "idCaptor", Value: data[0]}, primitive.E{Key: "iATA", Value: data[1]}, primitive.E{Key: "value", Value: data[3]}, primitive.E{Key: "pickingDate", Value: time.Now()}})
		if err != nil {
			fmt.Printf("Une erreur est survenue à l'enregistrement de la donnée\n")
			fmt.Printf("Plus d'infos : %s\n", err.Error())
		}
		fmt.Printf("ID : %s\n", res.InsertedID)
		fmt.Printf("Enregistrement réussie\n")
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
