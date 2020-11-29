<template>
    <div>
      <br>
      <b-container>

        <b-row >
          <b-col>
            <b-jumbotron>
              <!--<template #header>{{zone.Name}}</template>
              <template #lead>{{jwtDecoded.name}}</template>-->
              <b-list-group>
                <div class="item">
                  <b-list-group-item>
                    <h3 class="font-weight-bold">{{zone.Name}}</h3>
                    <b-progress :value="zone.PplCount" variant="dark" animated show-value :max="zone.Limits" class="mb-3"></b-progress>
                    <h6 class="font-weight-bold">Max: {{zone.Limits}} People</h6>
                  </b-list-group-item>
                </div>
              </b-list-group>
              <div class="btn">
                <b-button @click="addPerson" class="button" size=lg variant="success">Add Person</b-button>
              </div>
              <div class="btn2">
                <b-button @click="removePerson" class="button" size=lg variant="danger">Remove Person</b-button>
              </div>
              <!--<VueTable />-->
            </b-jumbotron>
          </b-col>
        </b-row>
      </b-container>
      <div class="backBtn">
        <button v-on:click="handlerOnclickBack" class="btn btn-dark btn-lg btn-block">Back</button>
      </div>
  </div>
</template>



<script>

import VueTable from '@/components/VueTable'
import settings from '@/settings'
import VueJwtDecode from 'vue-jwt-decode'

export default {
  props:{
    VueTable  
  },
  data(){
    return{
      ws: null,
      event: '',
      zone: {},
      jwtDecoded: {},
    }
  },
  methods: {
    loadPage(){
      try{
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
      }catch(e){
        console.error(e)
      }
      console.log(this.jwtDecoded)
      this.axios({
        method: 'get',
        url: `/zones/id/${this.$route.params.id}`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then((response) =>{
          this.zone = response.data.data
      }).catch(error => {
          if(error.response){
            this.$router.push({'name':'zones'})
          }
        });
    },
    addPerson(){
      this.axios({
        method: 'post',
        url: `/zones/id/${this.$route.params.id}/add`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then((response) =>{
          console.log(response)
          this.ws.send('updZone') //manda aviso para atualizar
      }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },
    removePerson(){
      this.axios({
        method: 'post',
        url: `/zones/id/${this.$route.params.id}/remove`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then((response) =>{
          console.log(response)
          this.ws.send('updZone') //manda aviso para atualizar
      }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },
    handlerOnclickBack(){
      this.$router.go(-1)
    },
  },
  mounted() {
    this.loadPage()
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
    event: function (newEvent) {
      if(newEvent.data === 'getZone'){
        this.loadPage();  //buscar zona atualizada
        this.event = '';  //reset ao evento
      }
    }
  }
}
</script>


<style scoped>
  .backBtn{
    margin: auto;
    padding: 2%;
    width: 30%;
  }
  .item{
    padding: 1%;
  }
  .btn{
    padding: 1%;
    width: 55%;
    margin: auto;
  }
  .btn2{
    padding: 1%;
    width: 55%;
    margin: auto;
  }
</style>