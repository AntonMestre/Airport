module github.com/AntonMestre/AirportProject

go 1.17

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	util v1.0.0
	tools v1.0.0
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
)

replace util v1.0.0 => ./internal/util
replace tools v1.0.0 => ./internal/tools