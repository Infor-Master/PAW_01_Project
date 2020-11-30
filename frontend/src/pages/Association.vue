<template>
 <div>
    <h1>Association</h1>
    <button v-on:click="handlerOnclickBack">Back</button>
    <p>
    <h5>Worker: {{this.selectedWorker}}</h5>
    <h5>Zones: {{this.zoneOptions}}</h5>
    <p>
    <form class="Associate" @submit.prevent="handlerSubmitAssociate()">
    
    <b-form-select v-model="selectedWorker" :plain="true" :options="this.userOptions" />
    <b-form-select v-model="selectedZone" :plain="true" :options="this.zoneOptions"/>
     
    <button type="submit">Associate</button>
    </form>

    <form class="Desassociate" @submit.prevent="handlerRemoveAssociate()">
    <button type="remove">Remove</button>
    </form>

 </div>
</template>

<script>
  
import settings from '../settings';
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
        baseURL: settings.baseURL,
        data: {
          idWorker: this.selectedWorker,
          idZone: this.selectedZone
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`        }
      })
      .then(response => {
        console.log(response)
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
        baseURL: settings.baseURL,
        data: {
          idWorker: this.selectedWorker,
          idZone: this.selectedZone
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`        }
      })
      .then(response => {
        console.log(response)

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
        baseURL: settings.baseURL,
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
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      })
      .then((response) =>{
        this.aux=response.data["data"];
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
        console.log()
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
    this.ws = new WebSocket(settings.socketURL)
    this.ws.onmessage = (event => {
      this.event = event;
    });
    this.ws.onopen = (event => {
        console.log(event)
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