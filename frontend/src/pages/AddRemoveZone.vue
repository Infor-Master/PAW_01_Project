<template>
 <div>
     <h1>Manage Zones</h1>
     <p>
    <button v-on:click="handlerOnclickBack">Back</button>
    <p>
   <form 
    class="addZone"
    @submit.prevent="handlerSubmitAdd">
    <hr>
    {{this.zones}}
    <hr>
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

    <b-form-select v-model="selected" :plain="true" :options="this.zoneOptions" />

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
      zones: [],
      zoneOptions: [],
      selected: '',
      message: this.message,
      name: this.name,
      latitude: this.latitude,
      longitude: this.longitude,
      limits: this.limits
    };
  },
  methods: {
    handlerOnclickBack(){
      this.$router.go(-1)
    },
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
      })
      .then(response => {
        console.log("ADD ZONE "+response);
        this.message = "Zone Registado!"
        this.getAllZones();
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });
    },
    handlerSubmitRemove(){

      var zoneID = 0;

      for(var j = 0; j < this.zones.length; j++){
      console.log(this.selected)
        if(this.zones[j]["name"]===this.selected){
          zoneID=this.zones[j].id;
        }
      }

      this.message = "";
      this.axios({
        method: 'delete',
        url: `admin/zones/`+zoneID,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      })
      .then(response => {
        console.log(response);
        this.message = "Zone Deleted!"
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });
    }
  },
  mounted(){
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

          this.zoneOptions.push(this.aux[j]["Name"])
          this.zones.push({
            "name": this.aux[j]["Name"],
            "id": this.aux[j]["ID"]
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
}

</script>