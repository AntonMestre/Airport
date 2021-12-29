package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"tools"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataFormat struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdCaptor    int                `json:"idCaptor,omitempty" bson:"idCaptor,omitempty"`
	IATA        string             `json:"iATA,omitempty" bson:"iATA,omitempty"`
	Value       int                `json:"value,omitempty" bson:"value,omitempty"`
	PickingDate time.Time          `json:"pickingDate,omitempty" bson:"pickingDate,omitempty"`
}

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

	//Route
	router := mux.NewRouter()
	router.HandleFunc("/data", GetData).Methods("GET")

	http.ListenAndServe(":1883", router)
}

func GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var captorsNamesInDb = []string{"Pressure", "Temp", "Wind"}

	//Retrieving query parameters
	queryValues := request.URL.Query()
	captor := queryValues.Get("captor")
	minDateQuery := queryValues.Get("minDate") + "+00:00"
	maxDateQuery := queryValues.Get("maxDate") + "+00:00"

	if !tools.StringInSlice(captor, captorsNamesInDb) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + "Wrong captor name" + `"}`))
		return
	}

	//Converting dates to match mongodb date format
	dateLayout := "2006-01-02T15:04:05.000+00:00" //golang time layout in mongodb format

	minDate, err := time.Parse(dateLayout, minDateQuery)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	maxDate, err := time.Parse(dateLayout, maxDateQuery)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	//Retrieving collection from mongodb
	var dataSet []DataFormat
	collection := dbClient.Database("AirportDataBase").Collection(captor)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"pickingDate": bson.M{"$gt": minDate, "$lt": maxDate}})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	//Parsing dataset
	for cursor.Next(ctx) {
		var data DataFormat
		cursor.Decode(&data)
		dataSet = append(dataSet, data)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(dataSet)
}
