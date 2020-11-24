<template>
    <div>
      <b-container>
        <b-row >
          <b-col>
            <b-jumbotron>
              <template #header>{{zone.Name}}</template>
              <template #lead>{{jwtDecoded.name}}</template>
              <h1> Number of people: {{zone.PplCount}} </h1>
              <b-button @click="addPerson" class="button" size=lg variant="success">Add Person</b-button>
              <p></p>
              <b-button class="button" size=lg variant="danger">Remove Person</b-button>
              <!--<VueTable />-->
            </b-jumbotron>
          </b-col>
        </b-row>
      </b-container>
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
      zone: {},
      jwtDecoded: {}
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
        url: `/zones/${this.$route.params.id}`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        },
        data: {
          teste: 'ola'
        }
      }).then((response) =>{
          this.zone = response.data.data
      }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },
    addPerson(){
      this.axios({
        method: 'post',
        url: `/zones/${this.$route.params.id}/add`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        },
        data: {
            zoneId : this.$route.params.id,
            workerId: VueJwtDecode.decode(localStorage.getItem('jwt')).id
        }
      }).then((response) =>{
          console.log(response)
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
        url: `/zones/${this.$route.params.id}/remove`,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        },
        data: {
            zoneId : this.$route.params.id,
            //workerId: VueJwtDecode.decode(localStorage.getItem('jwt')).id
        }
      }).then((response) =>{
          console.log(response)
      }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    }
  },
  mounted() {
    this.loadPage()
  }
}
</script>