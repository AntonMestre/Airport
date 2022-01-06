<template>
  <div>
    <airport-header v-bind:airport-name="airportName" v-bind:last-update="lastUpdate"></airport-header>
    <value-displayer v-bind:value-object="temperatureValue"></value-displayer>
    <value-displayer v-bind:value-object="windValue"></value-displayer>
    <value-displayer v-bind:value-object="pressureValue"></value-displayer>
    <graph></graph>
    <values-list v-bind:values-list="valuesList"></values-list>
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
  #airport-visualisation{
    display: grid;
    grid-template-rows: repeat(12,1fr);
  }


</style>