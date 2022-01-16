module github.com/AntonMestre/AirportProject

go 1.17

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	tools v1.0.0
	util v1.0.0
)

require (
	github.com/go-errors/errors v1.4.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/viant/toolbox v0.34.5
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.8.1 // indirect
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace util v1.0.0 => ./internal/util

replace tools v1.0.0 => ./internal/tools
