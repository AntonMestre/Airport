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
        <tr v-for="(value,key) in valuesList" :key="value">
          <td v-if="value.pressure != null">{{value.pressure}} hPa</td>
          <td v-else></td>
          <td v-if="value.wind != null">{{value.wind}} km/h</td>
          <td v-else></td>          
          <td v-if="value.temp != null">{{value.temp}} °C</td>
          <td v-else></td>  
          <td>{{key.slice(11)}}</td>
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

  mounted() {
    this.sortValuesList();
  },

  methods: {
    sortValuesList(){
      let valuesListArray = [];
      for(a in this.valuesList){
        valuesListArray.push([new Date().setTime(e.slice(11)),tihs.valuesList[a]])
      }
    }
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