package util

const MINIMUM_VALUE_WIND = 0
const MAXIMUM_VALUE_WIND = 50
const MINIMUM_VALUE_TEMP = 0
const MAXIMUM_VALUE_TEMP = 10
const MINIMUM_VALUE_PRESSURE = 1016
const MAXIMUM_VALUE_PRESSURE = 1030

const TOPIC_WIND = "topic/wind"
const TOPIC_TEMP = "topic/temp"
const TOPIC_PRESSURE = "topic/pressure"

const HOST = "tcp://localhost:1883"
const CLIENT_DATABASE_SUB = "CLIENT_DATABASE_SUB"

//API
const API_URI = ":3000"
const DATABASE_CLOUD_URI = "mongodb+srv://Airport:Airport@cluster0.0c6je.mongodb.net/AirportDataBase?retryWrites=true&w=majority"
const DATABASE_NAME = "AirportDataBase"

const DATE_LAYOUT = "2006-01-02T15:04:05.000+00:00" //golang time layout in mongodb format
const MIN_DATE = "0000-01-01T00:00:00.000+00:00"
const MAX_DATE = "9999-12-31T23:59:59.999+00:00"

var SENSORS_NAMES = []string{"Pressure", "Temp", "Wind"}

type Config struct {
	Nature   string `json:"nature"`
	IataCode string `json:"iatacode"`
	IdSensor byte   `json:"idsensor"`
	Broker   string `json:"broker"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Qoslevel byte   `json:"qoslevel"`
	ClientId string `json:"clientId"`
}
