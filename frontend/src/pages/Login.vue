<template>
 <div>
   <br>
   <b-card bg-variant="light" class="bCard">
   <form 
    class="login"
    @submit.prevent="handlerSubmit"
   >
      <h3>Sign in</h3>
      <br>
      <div class="signInForm">
        <div class="form-group">
          <label>Username</label>
          <input required v-model="username" class="form-control form-control-lg" type="text" placeholder="Username"/>
        </div>
        
        <div class="form-group">
          <label>Password</label>
          <input required v-model="password" class="form-control form-control-lg" type="password" placeholder="Password"/>
        </div>
        <button type="submit" class="btn btn-dark btn-lg btn-block">Login</button>
      </div>
      <p>{{message}}</p>
   </form>
   </b-card>
 </div>
</template>

<script>
import settings from '../settings'
import VueJwtDecode from 'vue-jwt-decode'

export default {
  name: "Login",
  data() {
    return {
      username: this.username,
      password: this.password,
      message: this.message,
    };
  },
 methods: {
    handlerSubmit(){
      this.message = "";
      this.axios({
        method: 'post',
        url: '/login',
        baseURL: settings.baseURL,
        data: {
          username: this.username,
          password: this.password,
        }
      }).then(response => {
          console.log(response)
          localStorage.setItem('jwt',response.data.token)
          try{
            let jwtDecoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
            if(jwtDecoded.admin){
              this.$router.push({name: 'admin'})
            }else{
              this.$router.push({name: 'zones'})
            }
          }catch(e){
            console.error(e)
          }
        }).catch(error => {
          // para ter acesso ao this.message, o catch tem de usar função via seta "=>"
          if(error.response){
            console.error(error.response);
            this.message = error.response.status + " - " + error.response.statusText;
          }
        });
    },
  },
};
</script>

<style>
.bCard{
  width: 25%;
  margin: auto;
}
html, 
body {
    margin: 0;
    padding: 0;
    background-color:rgb(131, 130, 130);
}
</style>