# Airport Project

GO project that retrieve and return weather data from airports sensors (temperature, wind speed and atmospheric pressure).

![mongo](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)
![maintened](https://img.shields.io/badge/Maintained%3F-yes-green.svg)
![askeme](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)
![go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)

## Built with
- [Golang](https://go.dev/)
- [Mosquitto](https://mosquitto.org/)
- [MongoDB](https://www.mongodb.com/)
- [VueJS](https://vuejs.org/)


## Getting started

### Prerequisites

You need to install the following in order to run this project :
- [Go](https://go.dev/dl/)
- [Mosquitto](https://mosquitto.org/download/)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/AntonMestre/Airport.git
   ```

## Usage

Because this project requires Mosquitto, you need to run it first with ``mosquitto -v``

To collect data from all the sensors, you will need to run each of them (PressurePub, TempPub and WindPub) with 
- ``go run sensor_name.go>``

Then, run the subscriber ``go run subDatabase.go``

Your subscriber shall now retrieve data from the sensors.

**Run API with :**
- `go run .\cmd\Main\main.go`

**To run client :**
- `cd .\client\`
- If it is the first time you lauch the client : `npm install`
- `npm run serve`

## Contributors

IMT Atlantique - FIL 2024

- Antonin Maystre
- Aurélie Lemoine
- Gabin Raapoto
- Luka Signe--Morice
- Fabien Hannon
