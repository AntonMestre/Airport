<template>
  <div>
    <div class="container">
      <div id="airport-header-container">
        <airport-header id="airport-header" v-bind:airport-name="airportName" v-bind:last-update="lastUpdate"></airport-header>
      </div>
    </div>
    <div class="container">
      <h2>
        Currently
      </h2>
      <div id="value-displayer-container">
        <value-displayer class="value-displayer" v-bind:value-object="temperatureValue"></value-displayer>
        <value-displayer class="value-displayer" v-bind:value-object="windValue"></value-displayer>
        <value-displayer class="value-displayer" v-bind:value-object="pressureValue"></value-displayer>
      </div>
    </div>
    <div class="container">
      <div id="graph-container">
        <graph></graph>
      </div>
    </div>
    <div class="container">
      <div id="values-list-container">
        <values-list v-bind:values-list="valuesList"></values-list>
      </div>
    </div>
  </div>
</template>

<script>
import ValueDisplayer from "./ValueDisplayer";
import AirportHeader from "./AirportHeader";
import Graph from "./Graph";
import ValuesList from "./ValuesList";
export default {
  name: "AirportVisualisation",
  components: {ValuesList, Graph, AirportHeader, ValueDisplayer},
  props: {
    iataCode: String
  },

  data: function () {
    return {
      temperatureValue: {

      },

      windValue: {

      },
      pressureValue: {

      },
      airportName: {

      },
      lastUpdate: {

      },
      valuesList: [],
    }
  }
}
</script>

<style scoped>
  h2{
    margin: 0;
    padding-bottom: 20px;
    font-weight: normal;
    color: #4C4C51;
  }

  #airportVisualisation{
    display: grid;
    grid-template-rows: repeat(10,1fr);
    margin: 0 60px 10px 60px;
    height: 100vh;
  }

  #airport-header-container{
    grid-row: 1/3;
  }

  #value-displayer-container{
    display: grid;
    grid-row: 4;
    grid-template-columns: repeat(3,1fr);
    column-gap: 50px;
  }
  


</style>