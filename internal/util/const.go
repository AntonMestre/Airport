package util

import "time"

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

const DB_URI = "mongodb://localhost:27017"

const NB_SEC = 10

//API
const API_URI = ":3000"
const DATABASE_CLOUD_URI = "mongodb://localhost:27017"
const DATABASE_NAME = "AirportDataBase"

const DATE_LAYOUT = "2006-01-02T15:04:05.000+00:00" //golang time layout in mongodb format
const MIN_DATE = "0000-01-01T00:00:00.000+00:00"
const MAX_DATE = "9999-12-31T23:59:59.999+00:00"
const DATE_FORMAT = "yyyy-MM-dd-HH-mm-ss"

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
	CityName string `json:"cityName"`
}

type Weather struct {
	Day   time.Time `json:"day"`
	Hours []Hours   `json:"hours"`
}

type Hours struct {
	Hour time.Time    `json:"hour"`
	Data Informations `json:"data"`
}

type Informations struct {
	Temp       float32 `json:"temp"`
	Pressure   float32 `json:"pressure"`
	Wind_speed float32 `json:"wind_speed"`
}
