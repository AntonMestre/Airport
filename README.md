# Airport Project

GO project that retrieve and return weather data from airports sensors (temperature, wind speed and atmospheric pressure).

## Built with
- [Golang](https://go.dev/)
- [Mosquitto](https://mosquitto.org/)
- [MongoDB](https://www.mongodb.com/)


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

## Contributors

IMT Atlantique - FIL 2024

- Antonin Maystre
- Aurélie Lemoine
- Gabin Raapoto
- Luka Signe--Morice
- Fabien Hannon