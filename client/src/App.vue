<template>
  <Sidebar v-bind:airports="airports"></Sidebar>
  <AirportVisualisation v-bind:iata-code="iataCode"></AirportVisualisation>
</template>

<script>
import Sidebar from "./components/Sidebar";
import axios from "axios";
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
