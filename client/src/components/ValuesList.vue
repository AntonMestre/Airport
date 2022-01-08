<template>
  <div id="value-list-container">
    <table>
      <thead>
        <tr>
          <td>Pressure</td>
          <td>Wind</td>
          <td>Temp</td>
          <td>Time</td>
        </tr>
      </thead>
      <tbody>
        <!-- <tr v-for="value in valuesList" :key="value">
          <td>{{value.pressure}} hPa</td>
          <td>{{value.wind}} km/h</td>
          <td>{{value.temperature}} °C</td>
          <td>{{value.time}}</td>
        </tr> -->
        <tr>
          <td>1000 hPa</td>
          <td>10 km/h</td>
          <td>12 °C</td>
          <td>12:32</td>
        </tr>
        <tr>
          <td>1000 hPa</td>
          <td>10 km/h</td>
          <td>12 °C</td>
          <td>12:32</td>
        </tr>
        <tr>
          <td>1000 hPa</td>
          <td>10 km/h</td>
          <td>12 °C</td>
          <td>12:32</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "ValuesList",
  props: {
    valuesList: Array,
  }
}
</script>

<style scoped>

  #value-list-container{
    background-color: white;
    border-radius: 15px;
    height: 100%;
    overflow: scroll;
    scrollbar-width: none;
  }

  table{
    padding-top: 15px;
  }

  thead td{
    color: #4D70F1;
    padding-bottom: 10px;
  }

  td{
    font-weight: lighter;
    color: #4C4C51;
    padding-left: 20px;
    padding-right: 10px;
  }
</style>