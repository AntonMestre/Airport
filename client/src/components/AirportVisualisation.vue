<template>
  <div v-if="isLoaded == true">
    <div id="airport-header-container">
        <airport-header id="airport-header" v-bind:airport-name="airport.name" v-bind:last-update="lastUpdate"></airport-header>
    </div>
    <div id="value-displayer-container">
      <h2>
        Currently
      </h2>
      <div id="value-components-container">
        <value-displayer class="value-displayer" v-bind:value-object="temperatureValue"></value-displayer>
        <value-displayer class="value-displayer" v-bind:value-object="windValue"></value-displayer>
        <value-displayer class="value-displayer" v-bind:value-object="pressureValue"></value-displayer>
      </div>
    </div>
    <div id="graph-container">
      <h2>
        Graphical view
      </h2>
      <graph-container :iataCode="airport.iata"></graph-container>
    </div>
    <div id="values-list-container">
      <h2>
        Last readings
      </h2>
      <values-list v-bind:values-list="sortValuesList(valuesList)"></values-list>
    </div>
  </div>
</template>

<script>
import ValueDisplayer from "./ValueDisplayer";
import AirportHeader from "./AirportHeader";
import GraphContainer from "./GraphContainer";
import ValuesList from "./ValuesList";
import axios from 'axios';

let date = new Date();

export default {
  name: "AirportVisualisation",
  components: {ValuesList, GraphContainer, AirportHeader, ValueDisplayer,},
  props: {
    airport: Object
  },

  mounted() {
    this.fetchDataFromSensors();
  },

  data: function () {
    return {
      isLoaded: false,
      temperatureValue: null,
      windValue: null,
      pressureValue: null,
      airportName: null,
      lastUpdate: null,
      valuesList: {},
    }
  },

  methods: {

    fetchDataFromSensors(){
      this.fetchData("Temp");
      this.fetchData("Wind");
      this.fetchData("Pressure");
    },
    // Temperature data --------------
    mapTempData(data){
      let lastData = data[data.length-1];
      this.temperatureValue = {
        name: "Temperature",
        currentValue: lastData.value,
        avgValue: 0,
        unit: "°C"
      };
      this.setLastUpdate(lastData.pickingDate);
      data.forEach(d => {
        let date = d.pickingDate.slice(0,19);
        if (date in this.valuesList) {
          this.valuesList[date].temp = d.value; 
        }
        else{
          this.valuesList[date] = {
            pressure: null,
            wind: null,
            temp: d.value,
          }
        }
      })
    },
  

    // Wind data --------------
    mapWindData(data){
      let lastData = data[data.length-1];
      this.windValue = {
        name: "Wind",
        currentValue: lastData.value,
        avgValue: 0,
        unit: "km/h"
      };
      this.setLastUpdate(lastData.pickingDate);
      data.forEach(d => {
        let date = d.pickingDate.slice(0,19);
        if (date in this.valuesList) {
          this.valuesList[date].wind = d.value; 
        }
        else{
          this.valuesList[date] = {
            pressure: null,
            wind: d.value,
            temp: null,
          }
        }
      })
    },

    // Pressure data --------------
    mapPressureData(data){
      let lastData = data[data.length-1];
      this.pressureValue = {
        name: "Pressure",
        currentValue: lastData.value,
        avgValue: 0,
        unit: "hPa"
      };
      this.setLastUpdate(lastData.pickingDate);
      data.forEach(d => {
        let date = d.pickingDate.slice(0,19);
        if (date in this.valuesList) {
          this.valuesList[date].pressure = d.value; 
        }
        else{
          this.valuesList[date] = {
            pressure: d.value,
            wind: null,
            temp: null,
          }
        }
      })
    },

    fetchData(sensor){
      this.fetchDataFromSensor(sensor);
    },
  
    fetchDataFromSensor(sensor){
      let data = null;
      axios
          .get('http://localhost:3000/data',{
            params: {
              sensor: sensor,
              // minDate: date.toISOString().slice(0, 10) + "T00:00:00.000",
              // maxDate: date.toISOString().slice(0, 10) + "T23:59:59.999",
              minDate: "2021-12-23T14:57:49.076",
              maxDate: "2021-12-29T15:16:29.801",
              iATA: this.airport.iata,
            }
          })
          .then(response =>{
            data = response.data
            if(sensor == "Temp")
              this.mapTempData(data);
            else if (sensor == "Wind")
              this.mapWindData(data);
            else if (sensor == "Pressure")
              this.mapPressureData(data);
            this.fetchAverageValues();
            this.isLoaded = true;
          })
          .catch(e => console.log(e))
    },

    fetchAverageValues(){
      let data = null;
      axios
          .get('http://localhost:3000/mean',{
            params: {
              // minDate: date.toISOString().slice(0, 10),
              // maxDate: date.toISOString().slice(0, 10),
              date: "2021-12-23",
              iATA: this.airport.iata,
            }
          })
          .then(response => {
            data = response.data;
            this.temperatureValue.avgValue = data.Temp.Average;
            this.pressureValue.avgValue = data.Pressure.Average;
            this.windValue.avgValue = data.Wind.Average;
          })
          .catch(e => console.log(e))
    },

    sortValuesList(valuesList){
      let valuesListArray = [];
      for (const [key, value] of Object.entries(valuesList)){
        let currDate = new Date(key.slice(0,19));
        valuesListArray.push([currDate, value.pressure, value.wind, value.temp]);
      }
      return valuesListArray.sort().reverse();
    },

    setLastUpdate(dateString){
      let date = new Date(dateString.slice(0,19));
      if(this.lastUpdate == null){
        this.lastUpdate = [dateString.slice(0,10), dateString.slice(11,19)];
      }
      else{
        if(this.lastUpdate < date){
          this.lastUpdate = date;
        }
      }
    }
  },

  watch: {
    // airport(){
    //   console.log("bipbap");
    //   this.fetchDataFromSensors();
    // }
  }
}

</script>

<style scoped>
  h2{
    margin: 0;
    padding-bottom: 15px;
    font-weight: normal;
    color: #4C4C51;
  }

  #airportVisualisation{
    display: grid;
    grid-template-rows: repeat(10,1fr);
    grid-template-columns: repeat(3,1fr);

    row-gap: 40px;
    column-gap: 50px;

    margin: 0 60px 10px 60px;
    height: 100vh;
  }

  #airport-header-container{
    grid-row: 1/3;
    grid-column: 1/4;
  }

  #value-displayer-container{
    grid-row: 4;
    grid-column: 1/4;

  }

  #value-components-container{
    display: grid;
    grid-template-columns: repeat(3,1fr);
    column-gap: 50px;
  }

  #graph-container{
    grid-row: 5/10;
    grid-column: 1/3;
  }

  #values-list-container{
    grid-row: 5/10;
    grid-column: 3;
  }

</style>