package main

import (
	"context"
	"fmt"
	"time"
	"tools"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	initDb()

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

func initDb() {

}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Airport:Airport@cluster0.0c6je.mongodb.net/AirportDataBase?retryWrites=true&w=majority"))
	fmt.Sprintf("%d", err)
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Printf("Ajout en BD de  - %s\n", msg.Payload())
	collection := dbClient.Database("AirportDataBase").Collection("Pressure")
	res, err := collection.InsertOne(ctx, bson.D{{"idCaptor", 1}, {"iATA", "TLS"}, {"value", 40}, {"pickingDate", time.Now()}})
	fmt.Printf("res  - %s\n", res)
	fmt.Printf("err  - %s\n", err)
	fmt.Sprintf("%d", res)
	fmt.Sprintf("%d", err)
}
