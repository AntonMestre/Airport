<template>
  <div>
    <h1>{{airportName}} Airport</h1>
        <h4 v-if="lastUpdate != null">Last update: {{lastUpdate.toISOString().slice(0, 10)}} at {{lastUpdate.toTimeString().slice(0,8)}}</h4>
        <h4 v-else id="no-values"> Aucune valeur a afficher pour aujourd'hui</h4>

  </div>
</template>

<script>
export default {
  name: "AirportHeader",

  props: {
    airportName: String,
    lastUpdate: Date
  }
}
</script>

<style scoped>
  #airport-header{
    text-align: center;
    margin-top: 3vh;
    margin-bottom: 5vh;
  }

  h1{
    margin: 0;
    color: #4D70F1;
    font-size: 3.4vw;
    font-weight: normal;
  }

  h4{
    margin: 0;
    font-size: 1vw;
    color: #817777;
    font-weight: lighter;
  }

  #no-values{
    color: #e31414;
  }
</style>