<template>
  <div>
    <div v-for="airport in airports" :key="airport">
      {{airport}}
    </div>
  </div>
</template>

<script>
export default {
  name: 'Sidebar',
  props: {
    airports: Array
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
