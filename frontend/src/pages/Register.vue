<template>
 <div>
   <form class="register">
     <h1>Sign up</h1>
     <label>User name</label>
     <input required v-model="username" type="text" placeholder="Username"/>
     <h1>{{username}}</h1>
     <hr/>
     <label>Name</label>
     <input required v-model="name" type="text" placeholder="Name"/>
     <h1>{{name}}</h1>
     <hr/>
     <label>Password</label>
     <input required v-model="password" type="password" placeholder="Password"/>
     <h1>{{password}}</h1>
     <hr/>
     <button v-on:click.prevent="handlerSubmit">Sign up</button>
   </form>
 </div>
</template>

<script>

export default {
  name: 'Register',
  data(){
    return {
      username: this.username,
      password: this.password,
      name: this.name
    };
  },
 methods: {
    handlerSubmit: function(){
      this.axios.post('http://localhost:8081/register', {  //não aceita endereço docker pois é no browser. Como porta é diferente origina problema de cors
          username: this.username,
          password: this.password,
          name: this.name
        }).then(response => {
          console.log(response);
          
          let status = response.status
        console.log("---REGISTER STATUS--- "+status);
        
        localStorage.setItem('user',JSON.stringify(response.data.user))

          if (localStorage.getItem('jwt') != null){
            
            this.$router.push('login')
        
          } 
          
        }).catch(function (error) {
          console.error(error.response);
        });
    }
  }
}
</script>