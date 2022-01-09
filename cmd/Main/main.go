package main

import (
	"context"
	"net/http"
	"util"

	"github.com/AntonMestre/AirportProject/cmd/Main/api"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var dbClient *mongo.Client

func main() {
	//Connection to MongoDB
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error
	dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(util.DATABASE_CLOUD_URI))
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	api.InitApiDatabaseClient(dbClient)

	//Route
	router := mux.NewRouter()

	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/data", api.GetData).Methods("GET", "OPTIONS")
	router.HandleFunc("/mean", api.GetMean).Methods("GET", "OPTIONS")
  router.HandleFunc("/airport", api.GetDataFromAirport).Methods("GET", "OPTIONS")

	http.ListenAndServe(util.API_URI, handlers.CORS(origins)(router))
}
