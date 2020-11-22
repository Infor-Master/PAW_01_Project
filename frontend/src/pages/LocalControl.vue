<template>
    <div>
      <b-container>
        <b-row >
          <b-col>
            <!--<b-jumbotron header="Local" lead="Person">-->
            <b-jumbotron>
              <template #header>{{zone.Name}}</template>
              <template #lead>{{jwtDecoded}}</template>
              <h1> Number of people: {{zone.PplCount}} </h1>
              <b-button class="button" size=lg variant="success">Add Person</b-button>
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
//import VueJwtDecode from 'vue-jwt-decode'

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
  mounted() {
    try{
      jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
    }catch(e){
      console.error(e)
    }
    console.log(jwtDecoded)
    this.axios({
      method: 'get',
      url: `/zones/${this.$route.params.id}`,
      baseURL: settings.baseURL,
      headers:{
        'Authorization': `Bearer ${localStorage.getItem('jwt')}`
      }
    }).then((response) =>{
        this.zone = response.data.data
    }).catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });
  }
}

</script>

