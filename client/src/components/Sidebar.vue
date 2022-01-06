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
        <h4>CDG</h4>
        <p>Paris Charles De Gaulle</p>
      </div>
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
      <div class="airport">
        <h4>CDG</h4>
        <p>Paris Charles De Gaulle</p>
      </div>
      <div class="airport">
        <h4>LYS</h4>
        <p>Lyon St Exupéry</p>
      </div>
      <div class="airport">
        <h4>CDG</h4>
        <p>Paris Charles De Gaulle</p>
      </div>

    </div>

    <div id="search">
      <img src="@/assets/search.png" alt="Search">
      <div id="search-field">
        <input type="text" placeholder="iATA code, city . . .">
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
  height: 100vh;
}

#logo {
  grid-row: 1/3;
  padding-top: 30px;
  width: 100%
}

#logo img{
  display: block;
  padding-left: 35px;
  width: 70%;
}

#airport-container{
  grid-row: 4/11;
  display: block;
  overflow: scroll;
  scrollbar-width: none;
  height: 90%;
}

.airport{
  color: #B4B4B4;
  margin-left: 35px;
  margin-top: 15px;
  margin-bottom: 15px;
  padding-top: 6px;
  padding-bottom: 6px;
}

.airport:hover{
  color: #4C4C51;
  transition: color 0.5s;
}

.airport h4{
  padding: 0;
  margin: 0;
  font-size: 22px;
  font-weight: normal;
}

.airport p{
  padding: 0;
  margin: 0;
  font-size: 17px;
  font-weight: lighter;
}

#search{
  display: grid;
  grid-template-columns: repeat(5,1fr);
  margin-left: 15px;
  margin-right: 15px;
  padding: 8px 8px 8px 8px;
  border-radius: 12px;
  grid-row: 12/13;
  background-color: #f8f8f8;
  height: fit-content;
}

#search img{
  grid-column: 1;
  width: 55%;
}

#search-field{
  grid-column: 2/5;
  align-self: center;
}

#search input{
  border: none;
  outline: none;
  background-color: transparent;
}

#seach input:focus{
  border: none;
  outline: none;
  color: #4C4C51;
}

#search input::placeholder{
  color: #B4B4B4;
  font-size: 15px;
  font-weight: 200;
}


</style>
