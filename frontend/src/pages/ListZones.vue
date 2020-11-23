<template>
  <div class="myDiv">
    <br>
    <div class="myBtn">
      <b-button variant="success" @click.prevent="logout">
        Logout
      </b-button>
    </div>
    <br><br><br>
    <b-jumbotron>
      <template slot="lead">
        <h3>
          <b>
            Zones
          </b>
        </h3>
        <p>{{message}}</p>
        <hr class="my-2" />
        <b-list-group>
        <b-list-group-item v-for="(zone,index) in zones[0]"
        :key="index">
          {{zone.Name}}
          <b-progress :value="value+5" show-value :max="max" class="mb-3"></b-progress>
        </b-list-group-item>
      </b-list-group>
        
      </template>
    </b-jumbotron>
  </div>
</template>

<script>
import settings from '../settings'
import VueJwtDecode from 'vue-jwt-decode'

export default {
  data(){
    return{
      zones: [],
      value: 33.333,
      max:100,
      message: this.message,
      jwtDecoded: {},
      id_zone: 0
    }
  },
  methods: {
    logout(){
      localStorage.removeItem('jwt')
      this.$router.push({name: 'login'})
    },
    getAllZones(){
      this.message = "";
      //this.axios.get('http://localhost:5000/api/zones/') 
      this.axios({
        method: 'get',
        url: `/zones/`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      })
      .then((response) =>{
          for(var i in response.data){
            this.zones.push(response.data[i])
         }
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });
    },
    getZoneOfWorker(){
      try{
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
        this.axios({
          method: 'get',
          url: `/workers/${this.jwtDecoded.id}`,
          baseURL: settings.baseURL,
          headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
      }
      }).then((response) =>{
          this.id_zone=response.data
          

      }).catch(error => {
          if(error.response){
            console.error(error.response);
          }
        });
        }catch(e){
          console.error(e)
        }
      }
  },
  mounted() {
      this.getAllZones();
      this.getZoneOfWorker();
  }
}
</script>

<style scoped>
.list-group {
  margin-bottom: 15px;
}
.list-group-item:hover {
  background: #EEE;
  cursor: pointer;
}
.myBtn{
  position: absolute; 
  right: 0;
  width: 38%;
  margin: auto;
}
.myDiv{
  width: 70%;
  align-content: center;
  margin: auto;
}
</style>