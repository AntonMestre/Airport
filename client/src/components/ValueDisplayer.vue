<template>
  <div>
    <h1>{{valueObject.name}}</h1>
    <h2>{{valueObject.currentValue}} {{unit}}</h2>
    <h4>{{valueObject.avgValue}} {{unit}}</h4>

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

</style>
