<template>
  <div>
    <div id="airport-header-container">
        <airport-header id="airport-header" v-bind:airport-name="airportName" v-bind:last-update="lastUpdate"></airport-header>
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
      <graph-container></graph-container>
    </div>
    <div id="values-list-container">
      <h2>
        Last readings
      </h2>
      <values-list v-bind:values-list="valuesList"></values-list>
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
    iataCode: String
  },

  mounted() {
    this.fetchData("Temp");
    this.fetchData("Wind");
    this.fetchData("Pressure");
  },

  data: function () {
    return {
      temperatureValue: null,
      windValue: null,
      pressureValue: null,
      airportName: null,
      lastUpdate: null,
      valuesList: {},
    }
  },

  methods: {

    // Temperature data --------------
    mapTempData(data){
      let lastData = data[data.length-1];
      this.temperatureValue = {
        name: "Temperature",
        currentValue: lastData.value,
        avgValue: 12, //Faire la requete API
        unit: "°C"
      };

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
        avgValue: 12, //Faire la requete API
        unit: "km/h"
      };

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
        avgValue: 12, //Faire la requete API
        unit: "hPa"
      };

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

    fetchData(type){
      let data = null;
      axios
          .get('http://localhost:3000/data',{
            params: {
              sensor: type,
              // minDate: date.toISOString().slice(0, 10) + "T00:00:00.000",
              // maxDate: date.toISOString().slice(0, 10) + "T23:59:59.999",
              minDate: "2021-12-23T14:57:49.076",
              maxDate: "2021-12-29T15:16:29.801",
            }
          })
          .then(response => data = response.data)
          .catch(e => console.log(e))
          .finally(() => {
            if(type == "Temp")
              this.mapTempData(data)
            else if (type == "Wind")
              this.mapWindData(data)
            else if (type == "Pressure")
              this.mapPressureData(data)
          })
    },
  },
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