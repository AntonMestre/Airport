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
        <tr v-for="value in valuesList" :key="value">
          <td v-if="value[1] != null">{{value[1]}} hPa</td>
          <td v-else></td>
          <td v-if="value[2] != null">{{value[2]}} km/h</td>
          <td v-else></td>          
          <td v-if="value[3] != null">{{value[3]}} °C</td>
          <td v-else></td>  
          <td>{{value[0].toTimeString().slice(0,8)}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "ValuesList",
  props: {
    valuesList: Object,
  },
}

</script>

<style scoped>

  #value-list-container{
    background-color: white;
    border-radius: 15px;
    height: 100%;
    overflow: scroll;
    scrollbar-width: none;
    box-shadow: 0px 3px 6px #E6E6E6;
  }

  table{
    margin-top: 15px;
  }

  thead td{
    color: #4D70F1;
    padding-bottom: 10px;
  }

  td{
    font-weight: lighter;
    color: #4C4C51;
    padding-left: 15px;
    padding-right: 10px;
  }
</style>