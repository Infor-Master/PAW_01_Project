<template>
 <div>
   <p>
    <button v-on:click="handlerOnclickBack">Back</button>
    <p>
   <form 
    class="register"
    @submit.prevent="handlerSubmit">
     <h1>Sign up</h1>
     <label>Username</label>
     <input required v-model="username" type="text" placeholder="Username"/>
     <hr/>
     <label>Name</label>
     <input required v-model="name" type="text" placeholder="Name"/>
     <hr/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <hr/>
     <input type="checkbox" id="isAdmin" v-model="isAdmin">
     <label for="checkbox">Admin</label>
     <p>
     <button type="submit">Sign up</button>
     <p>{{message}}</p>
     </form>

     <form 
        class="RemoveUser" @submit.prevent="handlerSubmitRemove()">

        <b-form-select v-model="selected" :plain="true" :options="this.userOptions" />

        <button type="submit">Remove</button>

   </form>
 </div>
</template>

<script>
import settings from '../settings'
import VueJwtDecode from 'vue-jwt-decode'


export default {
  name: 'Register',
  data(){
    return {
      jwtDecoded: {},
      workers: [],
      selected: '',
      userOptions: [],
      username: this.username,
      password: this.password,
      name: this.name,
      isAdmin: this.isAdmin,
      message: this.message,
    };
  },
 methods: {
   handlerOnclickBack(){
     this.$router.go(-1)
   },
    handlerSubmit(){
      this.message = "";
      this.axios({
        method: 'post',
        url: '/admin/users',
        baseURL: settings.baseURL,
        data: {
          username: this.username,
          password: this.password,
          name: this.name,
          admin: this.isAdmin
        },
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then(response => {
          console.log(response);
          this.message = "User Registado!"
        }).catch(error => {
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },handlerSubmitRemove(){

      var workerID = 0;

      for(var j = 0; j < this.workers.length; j++){
      console.log(this.selected)
        if(this.workers[j]["name"]===this.selected){
          workerID=this.workers[j].id;
        }
      }

      this.message = "";
      this.axios({
        method: 'delete',
        url: `admin/users/`+workerID,
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      })
      .then(response => {
        console.log(response);
        this.message = "Worker Deleted!"
      })
      .catch(error => {
        if(error.response){
          console.error(error.response);
          this.message = error.response.status + " - " + error.response.statusText;
        }
      });
    },
  },mounted() {
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
            this.userOptions.push(this.aux[j]["Name"])
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
}
</script>