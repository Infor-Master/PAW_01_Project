<template>
    <div>
      <b-container>
        <b-row >
          <b-col>
            <!--<b-jumbotron header="Local" lead="Person">-->
            <b-jumbotron lead="Person">
              <template #header>{{local.name}}</template>
              <h1>{{local}} Number of people: </h1>
              <b-button class="button" size=lg variant="success">Add Person</b-button>
              <p></p>
              <b-button class="button" size=lg variant="danger">Remove Person</b-button>
              <VueTable />
            </b-jumbotron>
          </b-col>
        </b-row>
      </b-container>



      <b-jumbotron header="Local" lead="Person">
        <h1>Number of people: </h1>
        <b-container class="bv-example-row">
          <b-row>
            <b-col>
              <b-button class="button" size=lg variant="success">Add Person</b-button>
           </b-col>
          </b-row>
          <b-row>
            <b-col>
              <b-button class="button" size=lg variant="danger">Remove Person</b-button>
            </b-col>
          </b-row>
        </b-container>
      </b-jumbotron>
</div>
</template>



<script>

import VueTable from '@/components/VueTable'
import settings from '@/settings'

export default {
  props:{
    VueTable    
  },
  data(){
    return{
      zone: []
    }
  },
  mounted() {
      this.axios({
        method: 'get',
        url: '/zones/1',
        baseURL: settings.baseURL
      }).then((response) =>{
          for(var i in response.data){
            this.zones.push(response.data[i])
           }
           console.log(this.zones)
        })
        .catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
  }
}

</script>

