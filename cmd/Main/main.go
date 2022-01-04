package main

import (
	"context"
	"net/http"

	"github.com/AntonMestre/AirportProject/cmd/Main/api"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func main() {
	//Connection to MongoDB
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error
	dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Airport:Airport@cluster0.0c6je.mongodb.net/AirportDataBase?retryWrites=true&w=majority"))
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	api.InitApiDatabaseVar(dbClient)

	//Route
	router := mux.NewRouter()
	router.HandleFunc("/data", api.GetData).Methods("GET")
	router.HandleFunc("/mean", api.GetMean).Methods("GET")

	http.ListenAndServe(":3000", router)
}
