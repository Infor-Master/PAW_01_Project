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

export default {
  name: 'Register',
  data(){
    return {
      selected: '',
      userOptions: [],
      aux:[],
      username: this.username,
      password: this.password,
      name: this.name,
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
          name: this.name
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

      var temp = 0;
    
      for(var j = 0; j < this.aux.length; j++){
        const aux = this.aux[j]["Username"];
        
        const aux2 = this.selected;

        if(aux===aux2){
          temp=j+1;
        }
       
        console.log(temp);
      }

        this.message = "";
        this.axios({
        method: 'delete',
        url: '/admin/users',
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
  },mounted() {
      this.message = "";
      this.axios({
        method: 'get',
        url: '/admin/users',
        baseURL: settings.baseURL,
        headers:{
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        }
      }).then((response) =>{

        this.aux=response.data["data"];

            for(var j = 0; j < this.aux.length; j++){
            var option = []
            option["username"] = this.aux[j]["Username"]

             this.userOptions.push(option['username'])
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