<script>
import { Line } from 'vue3-chart-v2'
export default {
  name: 'Graph',
  extends: Line,
  props: {
    chartdata: Array,
  },

  mounted() {
    this.renderGraph();
  },

  data() {
    return {
      chartOptions: {
        maintainAspectRatio: false, 
        legend: {
          display: false
        },
        plugins: {
            legend: {
                display: false,
            }
        },
        elements:{
          line: {
            backgroundColor: 'rgba(0, 0, 0, 0)',
            borderColor: 'rgba(77, 112, 241,0.5)',
            borderWidth: 2,
          },
          point: {
            backgroundColor: 'rgba(77, 112, 241, 1)',
            radius: 3,
          },
        },
        scales: {
          xAxes: [{
            gridLines: {
              color: 'rgba(0, 0, 0, 0.05)',
            },
            ticks: {
                fontSize: 10,
                fontColor: '#B4B4B4'
            }
          }],
          yAxes: [{
            gridLines: {
              color: 'rgba(0, 0, 0, 0.05)',
            },
            ticks: {
                fontSize: 10,
                fontColor: '#B4B4B4',
            }
          }]
        }
      }
    }
  },
  methods: {
    renderGraph(){
      if(this.chartdata != null){
        this.renderChart(
        {
          labels: this.chartdata[0],
          datasets: [{
            data: this.chartdata[1],
          }]
        }, this.chartOptions
        )
      }
    },

  },

  watch: {
    chartdata(val, oldval){
      console.log("Re rendering graph ...");
      this.renderGraph();
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
