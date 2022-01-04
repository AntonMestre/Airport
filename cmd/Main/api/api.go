package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"tools"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataFormat struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdCaptor    int                `json:"idCaptor,omitempty" bson:"idCaptor,omitempty"`
	IATA        string             `json:"iATA,omitempty" bson:"iATA,omitempty"`
	Value       int                `json:"value,omitempty" bson:"value,omitempty"`
	PickingDate time.Time          `json:"pickingDate,omitempty" bson:"pickingDate,omitempty"`
}

type AverageCaptor struct {
	Name    string
	Average int
}

var captorsNamesInDb = []string{"Pressure", "Temp", "Wind"}

var dbClient *mongo.Client

func InitApiDatabaseVar(myDbClient *mongo.Client) {
	dbClient = myDbClient
}

func GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

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

func GetMean(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	//Retrieving query parameters
	queryValues := request.URL.Query()
	dateQuery := queryValues.Get("date")

	//Converting dates to match mongodb date format
	dateLayout := "2006-01-02T15:04:05.000+00:00" //golang time layout in mongodb format
	minDateQuery := dateQuery + "T00:00:00.000+00:00"
	maxDateQuery := dateQuery + "T23:59:59.999+00:00"

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

	dataSet := make(map[string]AverageCaptor)
	//Retrieving collection from mongodb
	for i := 0; i < len(captorsNamesInDb); i++ {
		cursor, ctx := getCollection(captorsNamesInDb[i], minDate, maxDate)
		totalValue := 0
		nbRow := 0

		//Parsing dataset
		for cursor.Next(ctx) {
			var data DataFormat
			cursor.Decode(&data)
			totalValue += data.Value
			nbRow++
		}

		defer cursor.Close(ctx)

		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}

		data := AverageCaptor{captorsNamesInDb[i], totalValue / nbRow}

		dataSet[captorsNamesInDb[i]] = data
	}

	json.NewEncoder(response).Encode(dataSet)
}

func getCollection(name string, minDate time.Time, maxDate time.Time) (*mongo.Cursor, context.Context) {
	collection := dbClient.Database("AirportDataBase").Collection(name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"pickingDate": bson.M{"$gt": minDate, "$lt": maxDate}})

	if err != nil {
		return nil, nil
	}

	return cursor, ctx
}
