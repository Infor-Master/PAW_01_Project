<template>
  <div class="myDiv">
    <br>
    <div class="myBtn">
      
      <b-button variant="success" @click="$bvToast.show('example-toast')">
        {{this.worker_name}}
      </b-button>
      <b-toast id="example-toast" title="Logout?" static no-auto-hide>
        <b-button variant="secondary" @click.prevent="logout">
          Yes
        </b-button>
      </b-toast>
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
            :key="index"
            @click="zoneHandler(zone.ID)">
            {{zone.Name}}
          <b-progress :value="zone.PplCount" variant="dark" show-progress animated show-value :max="zone.Limits" class="mb-3"></b-progress>
          <h6>Max: {{zone.Limits}} People</h6>
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
      id_worker: 0,
      worker_name: null
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
    getWorkerZones(){
      try{
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
        this.id_worker=this.jwtDecoded.id;
        this.worker_name=this.jwtDecoded.name;
        this.axios({
          method: 'get',
          url: `/zones/worker`,
          baseURL: settings.baseURL,
          data: {
          id: this.id_worker,
          },
          headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
          }
          }).then((response) =>{
              for(var i in response.data){
                this.zones.push(response.data[i])
              }
          }).catch(error => {
              if(error.response){
                console.error(error.response);
              }
            });
          }catch(e){
            console.error(e)
          }
      },
      zoneHandler(id){
        this.$router.push('/zones/id/'+id)
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
    height: 40%;
    margin: auto;
  }
  .myDiv{
    width: 70%;
    align-content: center;
    margin: auto;
  }
</style>
