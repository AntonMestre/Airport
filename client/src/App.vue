<template>
  <Sidebar v-bind:airports="airports" id="sidebar"></Sidebar>
  <AirportVisualisation v-bind:iata-code="iataCode" id="airportVisualisation"></AirportVisualisation>
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
      iataCode: "NTE" //TODO Récuperer le IATA Code au clic
    }
  },
}
</script>

<style>

body{
  margin: 0;
  padding: 0;
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
