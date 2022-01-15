<template>
  <div>
    <img :src="require(`@/assets/${sensor}.png`)" :alt=sensor>
    <div class="values-container">
      <div>
        <h3>{{sensor}}</h3>
        <div v-if="valueObject != null" class="values">
          <p class="current">{{valueObject.currentValue}} {{valueObject.unit}}</p>
          <p class="separator">|</p>
          <p class="avg"> {{valueObject.avgValue}}  {{valueObject.unit}} avg</p>
        </div>
        <div v-else>
          <p class="current">Rien à afficher</p>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
export default {
  name: 'ValueDisplayer',
  props: {
    valueObject: Object,
    sensor: String
  }
}
</script>

<style scoped>
  .value-displayer{
    grid-row: 1;
    display: grid;
    grid-template-columns: repeat(5,1fr);
    background-color: white;
    border-radius: 15px;
    box-shadow: 0px 3px 6px #E6E6E6;
  }

  img{
    grid-column: 1;
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 80%;
    padding-top: 5px;
  }

  .values-container{
    grid-column: 2/6;
    padding: 10px 0px;

  }

  h3{
    font-weight: lighter;
    color: #817777;
    font-size: 14px;
    margin: 0;
  }

  p{
    margin: 0;
  }

  .values p{
    display: inline;
  }

  .current{
    color: #4D70F1;
    font-size: 23px;
  }

  .avg{
    color: #B4B4B4;
    font-weight: lighter;
    font-size: 14px;
  }
  
  .separator{
    color: #B4B4B4;
    font-weight: lighter;
    margin-left: 15px;
    margin-right: 15px;
  }


</style>
