<template>
 <div class="myDiv">
   <br>
   <b-jumbotron bg-variant="dark" text-variant="light" border-variant="success">
      <template slot="lead">
        <form 
          class="register"
          @submit.prevent="handlerSubmit">
          <h3><b>Sign up</b></h3>
          <hr>
          <h5>Username</h5>
          <input required v-model="username" type="text" placeholder="Username"/>
          <br><br>
          <h5>Name</h5>
          <input required v-model="name" type="text" placeholder="Name"/>
          <br><br>
          <h5>Password</h5>
          <input required v-model="password" type="password" placeholder="Password"/>
          <br><br>
          <h5 for="checkbox">Admin</h5>
          <input type="checkbox" id="isAdmin" style="height:25px;width:25px;" v-model="isAdmin">
          <br><br>
          <div class="btn2">
            <b-button type="submit" class="btn btn-dark btn-lg btn-block"
              variant="outline-primary">Sign up</b-button>
          </div>
          {{message}}
          <br><br>
        </form>
        <hr>
        <form 
          class="RemoveUser" @submit.prevent="handlerSubmitRemove()">

          <b-form-select v-model="selected" :plain="true" style="width:33%;margin:auto;" :options="this.userOptions" />
          <br>
          <div class="btn2">
            <b-button type="submit" class="btn btn-dark btn-lg btn-block"
        variant="outline-danger">Remove</b-button>
          </div>
        </form>
      </template>
   </b-jumbotron>
   <br>
   <div class="btn2">
    <b-button v-on:click="handlerOnclickBack"
      class="btn btn-dark btn-lg btn-block"
        variant="outline-success">Back</b-button>
   </div>
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