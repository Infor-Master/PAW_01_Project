<template>
 <div>
   <form 
    class="register"
    @submit.prevent="handlerSubmit">
     <h1>Sign up</h1>
     <label>User name</label>
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
 </div>
</template>

<script>
import settings from '../settings'

export default {
  name: 'Register',
  data(){
    return {
      username: this.username,
      password: this.password,
      name: this.name,
      message: this.message,
      //token: localStorage.getItem('jwt')
    };
  },
 methods: {
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
    }
  }
}
</script>