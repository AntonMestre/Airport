<template>
  <div>
    <table>
      <thead>
        <tr>
          <td></td>
          <td></td>
          <td></td>
        </tr>
      </thead>
      <tbody>

      </tbody>
        <tr v-for="value in valuesList" :key="value">
          <td>{{value.pressure}} hPa</td>
          <td>{{value.wind}} km/h</td>
          <td>{{value.temperature}} °C</td>
          <td>{{value.time}}</td>
        </tr>
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

</style>