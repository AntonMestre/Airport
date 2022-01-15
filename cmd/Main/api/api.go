package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"tools"
	"util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Structures
type DataFormat struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdCaptor    int                `json:"idCaptor,omitempty" bson:"idCaptor,omitempty"`
	IATA        string             `json:"iATA,omitempty" bson:"iATA,omitempty"`
	Value       float64            `json:"value,omitempty" bson:"value,omitempty"`
	PickingDate time.Time          `json:"pickingDate,omitempty" bson:"pickingDate,omitempty"`
}

type AverageSensor struct {
	Name    string
	Average float64
}

//Global vars for the API
var sensorsNamesInDb = util.SENSORS_NAMES
var dbClient *mongo.Client

// Init dbClient
func InitApiDatabaseClient(myDbClient *mongo.Client) {
	dbClient = myDbClient
}

//returns the cursor of the collection whose name is passed as a parameter
func getCollectionCursor(collectionName string, minDate time.Time, maxDate time.Time, iata string) (*mongo.Cursor, context.Context) {
	collection := dbClient.Database(util.DATABASE_NAME).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cursor *mongo.Cursor
	var err error

	if iata == "" {
		cursor, err = collection.Find(ctx, bson.M{"pickingDate": bson.M{"$gt": minDate, "$lt": maxDate}})
	} else {
		cursor, err = collection.Find(ctx, bson.M{"pickingDate": bson.M{"$gt": minDate, "$lt": maxDate}, "iATA": bson.M{"$eq": iata}})
	}

	if err != nil {
		return nil, nil
	}

	return cursor, ctx
}

func getCollectionAirport(collectionName string, iata string) (*mongo.Cursor, context.Context) {
	collection := dbClient.Database(util.DATABASE_NAME).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"iATA": bson.M{"$eq": iata}})

	if err != nil {
		return nil, nil
	}

	return cursor, ctx
}

//Retrieve datas of specific sensor, between two dates
func GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	//Retrieving query parameters
	queryValues := request.URL.Query()
	sensor := queryValues.Get("sensor")
	iata := queryValues.Get("iATA")
	minDateQuery := queryValues.Get("minDate") + "+00:00"
	maxDateQuery := queryValues.Get("maxDate") + "+00:00"

	if !tools.StringInSlice(sensor, sensorsNamesInDb) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + "Wrong sensor name" + `"}`))
		return
	}

	//Converting dates to match mongodb date format

	minDate, err := time.Parse(util.DATE_LAYOUT, minDateQuery)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	maxDate, err := time.Parse(util.DATE_LAYOUT, maxDateQuery)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	//Retrieving collection from mongodb
	var dataSet []DataFormat

	cursor, ctx := getCollectionCursor(sensor, minDate, maxDate, iata)

	defer cursor.Close(ctx)

	//Parsing dataset
	for cursor.Next(ctx) {
		var data DataFormat
		cursor.Decode(&data)
		value, err := cursor.Current.Values()
		if err != nil {
			panic(err)
		}
		data.PickingDate = value[4].Time()
		data.IATA = value[2].StringValue()
		var sensorString string
		errtemp := value[3].Unmarshal(&sensorString)
		if errtemp != nil {
			panic(errtemp)
		}
		sensorValue, sensorValueError := strconv.ParseFloat(sensorString, 32)
		if sensorValueError != nil {
			panic(sensorValueError)
		}
		data.Value = sensorValue
		dataSet = append(dataSet, data)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(response).Encode(dataSet)
}

//Get the mean value of all sensors from a specific day
func GetMean(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	//Retrieving query parameters
	queryValues := request.URL.Query()
	dateQuery := queryValues.Get("date")
	iata := queryValues.Get("iATA")

	//Converting dates to match mongodb date format
	minDateQuery := dateQuery + "T00:00:00.000+00:00"
	maxDateQuery := dateQuery + "T23:59:59.999+00:00"

	minDate, err := time.Parse(util.DATE_LAYOUT, minDateQuery)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	maxDate, err := time.Parse(util.DATE_LAYOUT, maxDateQuery)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	dataSet := make(map[string]AverageSensor)
	//Retrieving collection from mongodb
	for i := 0; i < len(sensorsNamesInDb); i++ {
		cursor, ctx := getCollectionCursor(sensorsNamesInDb[i], minDate, maxDate, iata)
		totalValue := 0.0
		nbRow := 0

		//Parsing dataset
		for cursor.Next(ctx) {
			var data DataFormat
			cursor.Decode(&data)
			totalValue += float64(data.Value)
			nbRow++
		}

		defer cursor.Close(ctx)

		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}

		if nbRow == 0 {
			continue
		}

		data := AverageSensor{sensorsNamesInDb[i], totalValue / float64(nbRow)}
		dataSet[sensorsNamesInDb[i]] = data
	}

	json.NewEncoder(response).Encode(dataSet)
}

//Retrieve datas of specific airport
func GetDataFromAirport(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	//Retrieving query parameters
	queryValues := request.URL.Query()
	airportIata := queryValues.Get("iATA")

	//Retrieving collection from mongodb
	dataSet := make(map[string][]DataFormat)

	for i := 0; i < len(sensorsNamesInDb); i++ {
		cursor, ctx := getCollectionAirport(sensorsNamesInDb[i], airportIata)
		var tempDataSet []DataFormat

		//Parsing dataset
		for cursor.Next(ctx) {
			var data DataFormat
			cursor.Decode(&data)
			tempDataSet = append(tempDataSet, data)
		}

		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}

		dataSet[sensorsNamesInDb[i]] = tempDataSet
	}

	json.NewEncoder(response).Encode(dataSet)
}
