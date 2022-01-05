<template>
  <div>
    <h1>airportName</h1>
    <h4>Last update: {{lastUpdate[0]}} at {{lastUpdate[1]}}</h4>
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

</style>