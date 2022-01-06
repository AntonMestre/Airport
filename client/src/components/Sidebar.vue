<template>
  <div id="sidebar">
    <div id="logo">
      <img src="@/assets/weatherfly.png" alt="WeatherFly logo">
    </div>
    <div id="airport-container">
<!--      <div v-for="airport in airports" :key="airport">
        {{airport}}
      </div>-->
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
    </div>

    <div id="search">
      <div id="search-field">
        <input type="text" placeholder="iATA code, city name ...">
      </div>
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

#sidebar{
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: repeat(12, 1fr);
}

#logo {
  grid-row: 1/3;
  padding-top: 40px;
}

#logo img{
  display: block;
  margin-left: 45px;
}

#airport-container{
  grid-row: 4/11;
}

.airport{
  color: #B4B4B4;
  margin-left: 45px;
  margin-top: 20px;
  margin-bottom: 20px;
  padding-top: 10px;
  padding-bottom: 10px;
}

.airport:hover{
  color: black;
  transition: color 0.5s;
}

.airport h4{
  padding: 0;
  margin: 0;
  font-size: 25px;
  font-weight: normal;
}

.airport p{
  padding: 0;
  margin: 0;
  font-size: 20px;
  font-weight: lighter;
}

#search{
  grid-row: 11/13;
}

#search-field{
  background-color: #f8f8f8;

}
</style>
