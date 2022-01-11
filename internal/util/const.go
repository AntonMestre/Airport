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
