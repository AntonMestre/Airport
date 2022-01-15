<template>
  <div id="big-graph-container">
    <div id="buttons">
      <p v-on:click="type = 'Temp'" :class="type == 'Temp' ? 'selected' : ''">Temperature</p>
      <p v-on:click="type = 'Wind'" :class="type == 'Wind' ? 'selected' : ''">Wind</p>
      <p v-on:click="type = 'Pressure'" :class="type == 'Pressure' ? 'selected' : ''">Pressure</p>
      <div id="from">
        <p>From </p>
        <input type="date" v-model="startDate" v-on:change="fetchData()">
      </div>
      <div id="to">
        <p>To </p>
        <input type="date"  v-model="endDate" v-on:change="fetchData()">       
      </div>
    </div>
    <graph :chartdata="chartdata" id="graph"></graph>
  </div>
</template>

<script>
import Graph from './Graph.vue';
import axios from 'axios';

let date = new Date();

export default 
{
  components: { Graph },
  name: 'BigGraphContainer',

  props: {
    iataCode: String,
  },
  
  data() {
    return {
      startDate: date.toISOString().slice(0, 10),
      endDate: date.toISOString().slice(0, 10),
      type: "Temp",
      chartdata: null,
    }
  },

  mounted() {
    this.fetchData();
  },

  methods: {
    mapData(data){
      let res = [[],[]];
      if(data == null){
        this.chartdata = res;
        return;
      }
      data.forEach(d => {
        res[0].push([d.pickingDate.slice(11,19), d.pickingDate.slice(0,10)]);
        res[1].push(d.value);
      });
      this.chartdata = res;
    },

    fetchData(){
      let data = null;
      axios
          .get('http://localhost:3000/data',{
            params: {
              sensor: this.type,
              minDate: this.startDate + "T00:00:00.000",
              maxDate: this.endDate + "T23:59:59.999",
              iATA: this.iataCode,
            }
          })
          .then(response => data = response.data)
          .catch(e => console.log(e))
          .finally(() => {
            this.mapData(data);
          })
    },
  },
  watch: {
    type(val, oldval){
      if(val !== oldval){
        this.fetchData(); 
      }
    },
    iataCode(){
      this.fetchData(); 
    },
  },
}
</script>

<style scoped>

  #big-graph-container{
    background-color: white;
    border-radius: 15px;
    box-shadow: 0px 3px 6px #E6E6E6;
    margin: 20px 60px 20px 60px;

  }

  #graph{
    height: 90%;
    padding-left: 28px;
    padding-right: 28px;

  }

  p{
    color: #B4B4B4;
    font-weight: lighter;
    cursor: pointer;
  }

  #buttons{
    margin-left: 30px;
    display: grid;
    grid-template-columns: repeat(10,auto);
  }

  #from{
    grid-column: 8;
  }

  #from, #to{
    display: flex;
    justify-content: space-between;
    margin-right: 20px;
  }

  input{
    height: 30%;
    align-self: center;
    border: none;
    font-family: "Cera Pro";
    font-weight: lighter;
    color: #4C4C51;
  }

  .selected{
    color: #4D70F1;
  }

</style>
