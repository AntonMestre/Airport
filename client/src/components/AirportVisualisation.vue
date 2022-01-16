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
        <value-displayer class="value-displayer" v-bind:value-object="temperatureValue" v-bind:sensor="'Temperature'"></value-displayer>
        <value-displayer class="value-displayer" v-bind:value-object="windValue" v-bind:sensor="'Wind'"></value-displayer>
        <value-displayer class="value-displayer" v-bind:value-object="pressureValue" v-bind:sensor="'Pressure'"></value-displayer>
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
    // Récupération des données nécessaires 
    this.fetchAll();
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
    fetchAll(){
      this.fetchDataFromSensor("Temp");
      this.fetchDataFromSensor("Wind");
      this.fetchDataFromSensor("Pressure");
      this.fetchAverageValues();
    },
    // Temperature data --------------
    mapTempData(data){
      if(data == null){
        this.temperatureValue = null;
        return;
      }
      let lastData = data[data.length-1];
      this.temperatureValue = {
        name: "Temperature",
        currentValue: lastData.value.toFixed(2),
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
      if(data == null){
        this.windValue = null;
        return;
      }
      let lastData = data[data.length-1];
      this.windValue = {
        name: "Wind",
        currentValue: lastData.value.toFixed(2),
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
      if(data == null){
        this.pressureValue = null;
        return;
      }
      let lastData = data[data.length-1];
      this.pressureValue = {
        name: "Pressure",
        currentValue: lastData.value.toFixed(2),
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
  
    fetchDataFromSensor(sensor){
      let data = null;
      axios
          .get('http://localhost:3000/data',{
            params: {
              sensor: sensor,
              minDate: date.toISOString().slice(0, 10) + "T00:00:00.000",
              maxDate: date.toISOString().slice(0, 10) + "T23:59:59.999",
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
            this.isLoaded = true;
          })
          .catch(e => console.log(e))
    },

    fetchAverageValues(){
      let data = null;
      axios
          .get('http://localhost:3000/mean',{
            params: {
              date: date.toISOString().slice(0, 10),
              iATA: this.airport.iata,
            }
          })
          .then(response => {
            data = response.data;
            if(data != null && Object.keys(data).length > 0){
              this.temperatureValue.avgValue = data.Temp.Average.toFixed(2);
              this.pressureValue.avgValue = data.Pressure.Average.toFixed(2);
              this.windValue.avgValue = data.Wind.Average.toFixed(2);
            }
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
      if(dateString == null){
        this.lastUpdate = null;
        return;
      }
      let date = new Date(dateString.slice(0,19));
      if(this.lastUpdate == null){
        this.lastUpdate = date;
      }
      else{
        if(this.lastUpdate < date){
          this.lastUpdate = date;
        }
      }
    }
  },

  watch: {
    airport(){
      this.fetchAll();
    }
  }
}

</script>

<style scoped>
  h2{
    margin: 0;
    padding-bottom: 1vw;
    font-weight: normal;
    color: #4C4C51;
    font-size: 1.55vw;
  }

  #airportVisualisation{
    display: grid;
    grid-template-rows: 15% 15% 43%;
    grid-template-columns: repeat(3,1fr);

    row-gap: 5vh;
    column-gap: 4vw;

    margin: 0 4vw 1vh 4vw;
    height: 100vh;
  }

  #airport-header-container{
    grid-row: 1;
    grid-column: 1/4;
  }

  #value-displayer-container{
    grid-row: 2;
    grid-column: 1/4;

  }

  #graph-container{
    grid-row: 3;
    grid-column: 1/3;
  }

  #value-components-container{
    display: grid;
    grid-template-columns: repeat(3,1fr);
    column-gap: 4vw;
  }

  #values-list-container{
    grid-row: 3;
    grid-column: 3;
  }

</style>