<template>
  <div>
    <!-- <h1>{{valueObject.name}}</h1>
    <h2>{{valueObject.currentValue}} {{unit}}</h2>
    <h4>{{valueObject.avgValue}} {{unit}}</h4> -->
    <img src="@/assets/temperature.png" alt="">
    <div class="values-container">
      <h3>Temperature</h3>
      <div class="values">
        <p class="current">12 °C</p>
        <p class="separator">|</p>
        <p class="avg"> 10 °C avg</p>
      </div>
    </div>
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
    width: 30%;
  }

  .values-container{
    grid-column: 2/5;
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
