<template>
 <div>
   <form 
    class="login"
    @submit.prevent="handlerSubmit"
   >
     <h1>Sign in</h1>
     <label>User name</label>
     <input required v-model="username" type="text" placeholder="Username"/>
     <hr/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <hr/>
     <button type="submit">Login</button>
     <p>{{message}}</p>
   </form>
 </div>
</template>

<script>
import settings from '../settings'

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
          // descompactar o jwt e guardar no localstorage o nome da criatura
          
          //descompactar jwt e fazer if para zones ou admin
          
          this.$router.push('zones')

          /*
          let is_admin = response.data.user.is_admin
          localStorage.setItem('user',JSON.stringify(response.data.user))
          if (localStorage.getItem('jwt') != null){
            if(is_admin== 1){
              this.$router.push('admin')
            }else {
              this.$router.push('map')
            }
          } 
          */
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
