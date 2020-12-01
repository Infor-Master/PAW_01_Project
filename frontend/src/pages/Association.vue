<template>
 <div class="myDiv">
   <br>
   <b-jumbotron bg-variant="dark" text-variant="light" border-variant="success">
      <template slot="lead">
        <h3><b>Association</b></h3>
        <hr>
        <form class="Associate" @submit.prevent="handlerSubmitAssociate()">
        <br>
        <b-form-select v-model="selectedWorker" :plain="true" style="width:33%;margin:auto;" :options="this.userOptions">
          <b-form-select-option :value="null" v-model="selectedWorker" :plain="true" style="width:33%;margin:auto;" :options="this.userOptions" disabled>Select a Worker</b-form-select-option>
        </b-form-select>
        <br>
        <b-form-select v-model="selectedZone" :plain="true" style="width:33%;margin:auto;" :options="this.zoneOptions">
          <b-form-select-option :value="null" v-model="selectedZone" disabled>Select a Zone</b-form-select-option>
        </b-form-select>
        <br>
        <div class="btn2">
          <b-button type="submit" class="btn btn-dark btn-lg btn-block"
          variant="outline-primary">Associate</b-button>
        </div>
        </form>
        <form class="Desassociate" @submit.prevent="handlerRemoveAssociate()">
          <div class="btn2">
            <b-button type="remove" class="btn btn-dark btn-lg btn-block"
          variant="outline-danger">Remove</b-button>
          </div>
        </form>
      </template>
    </b-jumbotron>
    <div class="btn2">
      <b-button v-on:click="handlerOnclickBack" class="btn btn-dark btn-lg btn-block"
          variant="outline-success">Back</b-button>
    </div>
 </div>
</template>

<script>
  
import settings from '../../configs.json';
import VueJwtDecode from 'vue-jwt-decode'

  export default {
  name: 'association',
  data(){
    return {
      ws: null,
      event: '',
      zones: [],
      selectedZone: '',
      selectedWorker: '',
      zoneOptions: [],
      userOptions: [],
      workers: []
    };
  },
  methods: {
    handlerOnclickBack(){
      this.$router.go(-1)
    },
    handlerSubmitAssociate(){
        this.message = "";
        this.axios({
        method: 'post',
        url: '/admin/associate',
        baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
        data: {
          idWorker: this.selectedWorker,
          idZone: this.selectedZone
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`        }
      })
      .then(() => {
        this.message = " Registado!"
        this.ws.send('updZone') //manda aviso para atualizar
        this.getAllZones();
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });

    },
    handlerRemoveAssociate(){
        this.message = "";
        this.axios({
        method: 'delete',
        url: '/admin/associate',
        baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
        data: {
          idWorker: this.selectedWorker,
          idZone: this.selectedZone
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`        }
      })
      .then(() => {

        this.message = " Registado!"
        this.ws.send('updZone') //manda aviso para atualizar
        this.getAllZones();
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });
      
    },
    getZones(){
    this.message = "";
      this.axios({
        method: 'get',
        url: '/zones/all',
        baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      })
      .then((response) =>{
        this.aux=response.data["data"];
            
        for(var j = 0; j < this.aux.length; j++){

            this.zoneOptions.push({
                value: this.aux[j]["ID"],
                text: this.aux[j]["Name"],
            })

          this.zones.push({
            "id": this.aux[j]["ID"],
            "name": this.aux[j]["Name"]
          })
        }
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      })
   },
   getWorkers(){
      this.message = "";
      this.axios({
        method: 'get',
        url: 'admin/users',
        baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      })
      .then((response) =>{
        this.aux=response.data["data"];
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
        for(var j = 0; j < this.aux.length; j++){
            
          this.workers.push({
            "name": this.aux[j]["Name"],
            "id": this.aux[j]["ID"]
          })
          if(this.aux[j]["ID"]!=this.jwtDecoded.id){
            this.userOptions.push({
                value: this.aux[j]["ID"],
                text: this.aux[j]["Name"],
            })
          }
        }
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      })
   }
  },
  mounted(){
    this.getZones()
    this.getWorkers()
  },
  created() {
    this.ws = new WebSocket(settings.sockets.protocol + settings.URL + settings.sockets.path)
    this.ws.onmessage = (event => {
      this.event = event;
    });
    this.ws.onopen = (() => {
        console.log("Successfully connected to the websocket server...")
    })
  },
  watch: {
    event: function () {
      //se for preciso para receber eventos
    }
  }
}
</script>


<style scoped>
  .btn2 {
    margin: auto;
    padding: 2%;
    width: 30%;
  }
  .myDiv{
    width: 60%;
    align-content: center;
    margin: auto;
  }
  hr { 
    display: block; height: 2px;
    border: 0; border-top: 2px solid rgb(27, 180, 52);
    margin: 1em 0; padding: 0; 
  }
</style>