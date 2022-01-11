<template>
  <div v-if="valueObject != null">
    <img :src="require(`@/assets/${valueObject.name}.png`)" :alt=valueObject.name>
    <div class="values-container">
      <div v-if="valueObject != null">
        <h3>{{valueObject.name}}</h3>
        <div class="values">
          <p class="current">{{valueObject.currentValue}} {{valueObject.unit}}</p>
          <p class="separator">|</p>
          <p class="avg"> {{valueObject.avgValue}}  {{valueObject.unit}} avg</p>
        </div>
      </div>
    </div>
  </div>
  <div v-else>
    <p id="error">Rien a afficher</p>
  </div>
</template>

<script>
export default {
  name: 'ValueDisplayer',
  props: {
    valueObject: Object
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

  #error{
    font-size: 20px;
    font-weight: lighter;
  }
</style>
