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

export default {
  name: "AirportVisualisation",
  components: {ValuesList, GraphContainer, AirportHeader, ValueDisplayer,},
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