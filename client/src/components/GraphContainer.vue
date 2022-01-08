<template>
  <div id="graph-container">
    <div id="buttons">
      <p>Temperature</p>
      <p>Wind</p>
      <p>Pressure</p>
      <div id="from">
        <p>From </p>
        <input type="date">
      </div>
      <div id="to">
        <p>To </p>
        <input type="date">       
      </div>
    </div>
    <graph></graph>
  </div>
</template>

<script>
import Graph from "./Graph.vue"

export default 
{
  components: { Graph },
  name: 'GraphContainer',
  props: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

  #graph-container{
    background-color: white;
    border-radius: 15px;
  }

  p{
    color: #B4B4B4;
    font-weight: lighter;
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
  }

</style>
