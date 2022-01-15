<template>
  <Sidebar v-bind:airports="airports" id="sidebar" @airport="getAirport"></Sidebar>
  <div id="switch-view">
    <p v-on:click="isOverview = !isOverview">Switch view →</p>
  </div>
  <AirportVisualisation v-bind:airport=airport id="airportVisualisation" v-if="isOverview"></AirportVisualisation>
  <big-graph-container :iataCode="airport.iata" v-if="!isOverview" id="big-graph-container"></big-graph-container>
</template>

<script>
import Sidebar from "./components/Sidebar";
import AirportVisualisation from "./components/AirportVisualisation";
import BigGraphContainer from "./components/BigGraphContainer";

export default {
  name: 'App',
  components: {
    AirportVisualisation,
    Sidebar,
    BigGraphContainer
  },
  data: function() {
    return {
      airports: [
          //TODO Récuperer les airports via l'api
      ],
      airport: {
        iata: "NTE",
        name: "Nantes",
      },
      isOverview: true,
    }
  },
  
  methods: {
    getAirport(airport){
      this.airport = airport;
    }
  },
}
</script>

<style>

body{
  margin: 0;
  padding: 0;
  overflow: hidden;
}

html, body, #app{
  width: 100%;
  height: 100%;
}
#app {
  display: grid;
  grid-template-columns: repeat(6,1fr);
  grid-template-rows: 7% 93%;
  grid-column-gap: 1vw;
  font-family: Cera Pro,sans-serif;
  background-color: #f8f8f8;
}

#sidebar{
  grid-column: 1;
  grid-row: 1;
  height: 100%;
  background-color: white;
}

#switch-view{
  grid-row: 1;
  grid-column: 6/7;
  text-align: center;
  cursor: pointer;
  color: #817777;
  margin-top: 2.5vh;
  margin-left: 2vh;
  font-size: 1vw;
}


#airportVisualisation, #big-graph-container{
  grid-column: 2/7;
  grid-row: 2;
}

::-webkit-scrollbar{
  width: 0;
}

</style>
