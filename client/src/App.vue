<template>
  <Sidebar v-bind:airports="airports" id="sidebar" @airport="getAirport"></Sidebar>
  <AirportVisualisation v-bind:airport=airport id="airportVisualisation"></AirportVisualisation>
</template>

<script>
import Sidebar from "./components/Sidebar";
import AirportVisualisation from "./components/AirportVisualisation";

export default {
  name: 'App',
  components: {
    AirportVisualisation,
    Sidebar,
  },
  data: function() {
    return {
      airports: [
          //TODO Récuperer les airports via l'api
      ],
      airport: {
        iata: "TLS",
        name: "Toulouse-Blagnac",
      },
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
  grid-column-gap: 10px;
  font-family: Cera Pro,sans-serif;
  background-color: #f8f8f8;
}

#sidebar{
  grid-column: 1;
  grid-row: 1;
  height: 100%;
  background-color: white;
}

#airportVisualisation{
  grid-column: 2/7;
  grid-row: 1;
  height: 100%;
}
</style>
