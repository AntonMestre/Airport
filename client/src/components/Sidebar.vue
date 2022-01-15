<template>
  <div id="sidebar">
    <div id="logo">
      <img src="@/assets/weatherfly.png" alt="WeatherFly logo">
    </div>
    <div id="airport-container">
     <div v-for="airport in airports" :key="airport">
        <div class="airport" :id="iata == airport.iata ? 'current' : ''" v-on:click="$emit('airport', {name: airport.name, iata: airport.iata}); select(airport.iata)">
          <div class="hover-truc"></div>
          <div>
            <h4>{{airport.iata}}</h4>
            <p>{{airport.name}}</p>
          </div>
        </div>
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
  data() {
    return {
      // airport: {
      //   iata: "TLS",
      //   name: "Toulouse-Blagnac",
      // }
      airports: [
        {
          name: "Nantes",
          iata: "NTE"
        },
        {
          name: "Toulouse-Blagnac",
          iata: "TLS"
        },
        {
          name: "Montréal-Trudeau",
          iata: "YUL"
        },
        {
          name: "Dubai",
          iata: "DXB"
        },
        {
          name: "Lille-Lesquin",
          iata: "LIL"
        },
      ],
      iata: "NTE", // IATA de base
    }
  },

  methods: {
    select(iata){
        this.iata = iata;
    }
  },
}
</script>

<style scoped>

#sidebar{
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: repeat(12, 1fr);
  height: 100vh;
}

#logo {
  grid-row: 1/3;
  padding-top: 3.7vh;
  width: 100%
}

#logo img{
  display: block;
  padding-left: 2.3vw;
  width: 70%;
}

#airport-container{
  grid-row: 4/11;
  display: block;
  overflow: scroll;
  scrollbar-width: none;
  height: 90%;
}

.hover-truc{
    background-color: white;
    width: 0%;
    height: 25%;

}

.airport:hover .hover-truc, #current .hover-truc{
  background-color: #4D70F1;
  width: 15%;
  /* height: 120%; */
  border-radius: 0px 10px 10px 0px;
  align-self: center;
  transition:  0.5s;
}

 #current .hover-truc{
   height: 120%;
 }

.airport{
  display: grid;
  grid-template-columns: 15% 85%;
  color: #B4B4B4;
  margin-top: 1.8vh;
  margin-bottom: 1.8vh;
  padding-top: 0.8vh;
  padding-bottom: 0.8vh;
}

.airport:hover{
  color: #4C4C51;
  transition: color 0.5s;
  cursor: pointer;
}

.airport h4{
  padding: 0;
  margin: 0;
  font-size: 1.4vw;
  font-weight: normal;
}

.airport p{
  padding: 0;
  margin: 0;
  font-size: 1.1vw;
  font-weight: lighter;
}

#search{
  display: grid;
  grid-template-columns: repeat(5,1fr);
  margin-left: 1vw;
  margin-right: 1vw;
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
  font-size: 1vw;
  font-weight: 200;
}

#current{
  color: #4D70F1;
}



</style>
