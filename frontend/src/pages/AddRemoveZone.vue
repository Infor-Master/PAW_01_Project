<template>
 <div>
     <h1>Manage Zones</h1>
   <form 
    class="addZone"
    @submit.prevent="handlerSubmitAdd">
     <label>Name</label>
     <input required v-model="name" type="text" placeholder="Name"/>
     <p/>
     <label>Latitude</label>
     <input required v-model="latitude" type="number" step="0.00000000000000001" placeholder="Latitude"/>
     <p/>
     <label>Longitude</label>
     <input required v-model="longitude" type="number" step="0.00000000000000001" placeholder="Longitude"/>
     <p/>
     <label>Limit</label>
     <input required v-model="limits" type="number" placeholder="Limit"/>
     <p/>
     <button type="submit">Add</button>
     {{this.message}}
   </form>

   <form 
    class="RemoveZone" @submit.prevent="handlerSubmitRemove()">

    <b-form-select v-model="selected" :plain="true" :options="this.zonesOptions" />

    <button type="submit">Remove</button>
   </form>
 </div>
</template>

<script>
import settings from '../settings';

export default {
  name: 'addRemoveZone',
  data(){
    return {
      selected: '',
      zonesOptions: [],
      aux:[],
      message: this.message,
      name: this.name,
      latitude: this.latitude,
      longitude: this.longitude,
      limits: this.limits
    };
  },
 methods: {
    handlerSubmitAdd(){
      this.message = "";
      this.axios({
        method: 'post',
        url: '/admin/zones',
        baseURL: settings.baseURL,
        data: {
            name: this.name,
            latitude: parseFloat(this.latitude),
            longitude: parseFloat(this.longitude),
            limits: parseInt(this.limits)
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then(response => {
          console.log("ADD ZONE "+response);
          this.message = "Zone Registado!"
        }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },handlerSubmitRemove(){

      var temp = 0;
    
      for(var j = 0; j < this.aux.length; j++){
        const aux = this.aux[j]["Name"];
        
        const aux2 = this.selected;

        if(aux===aux2){
          temp=j+1;
        }
       
        console.log(temp);

      }

        this.message = "";
        this.axios({
        method: 'delete',
        url: '/admin/zones',
        baseURL: settings.baseURL,
        data: {
          id: temp,
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then(response => {
          console.log(response);
          this.message = "Zone Deleted!"
        }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },
  },
  mounted() {
      this.message = "";
      this.axios({
        method: 'get',
        url: '/zones',
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then((response) =>{

        this.aux=response.data["data"];

            for(var j = 0; j < this.aux.length; j++){
            var option = []
            option["name"] = this.aux[j]["Name"]

             this.zonesOptions.push(option['name'])

          }
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