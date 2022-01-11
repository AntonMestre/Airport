<template>
  <div>
    <h1>{{airportName}} Airport</h1>
    <!-- <h4>Last update: {{lastUpdate[0]}} at {{lastUpdate[1]}}</h4> -->
        <h4 v-if="lastUpdate != null">Last update: {{lastUpdate[0]}} at {{lastUpdate[1]}}</h4>

  </div>
</template>

<script>
export default {
  name: "AirportHeader",

  props: {
    airportName: String,
    lastUpdate: Array // [date, time]
  }
}
</script>

<style scoped>
  #airport-header{
    text-align: center;
    margin-top: 8vh;
    margin-bottom: 8vh;
  }

  h1{
    margin: 0;
    color: #4D70F1;
    font-size: 50px;
    font-weight: normal;
  }

  h4{
    margin: 0;
    color: #817777;
    font-weight: lighter;

  }
</style>