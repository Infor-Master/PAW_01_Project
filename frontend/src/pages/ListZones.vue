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
        <b-list-group >
          <b-list-group-item v-for="(zone,index) in zones[0]"
            :key="index">
            {{zone.Name}}
          <b-progress :value="zone.PplCount" show-value :max="zone.Limits" class="mb-3"></b-progress>
        </b-list-group-item>
          
      </b-list-group>
        
      </template>
    </b-jumbotron>
    <b-button variant="success" @click.prevent="map">
      Map
    </b-button>
  </div>
</template>

<script>
import settings from '../settings'
import VueJwtDecode from 'vue-jwt-decode'

export default {
  data(){
    return{
      zones: [],
      message: this.message,
      jwtDecoded: {},
      id_worker: 0
    }
  },
  methods: {
    logout(){
      localStorage.removeItem('jwt')
      this.$router.push({name: 'login'})
    },
    map(){
      this.$router.push({name: 'map'})
    },
    /*getAllZones(){
      this.message = "";
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
    },*/
    getWorkerZones(){
      try{
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
        this.id_worker=this.jwtDecoded.id;
        this.axios({
          method: 'get',
          url: `/zones`,
          baseURL: settings.baseURL,
          data: {
          id: this.id_worker,
          },
          headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
          }
          }).then((response) =>{
              console.log(response.data)
              
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
      this.getWorkerZones();
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
