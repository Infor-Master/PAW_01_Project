<template>
 <div>
    <div class="myBtn">
      
      <b-button variant="success" @click="$bvToast.show('toast')">
        <h5>Logout</h5>
      </b-button>
      <b-toast id="toast" title="Logout?" static no-auto-hide>
        <b-button variant="secondary" @click.prevent="logout">
          Yes
        </b-button>
      </b-toast>
    </div>
    <button v-on:click="handlerOnclickZones" >Manage Zones</button>
    <button v-on:click="handlerOnclickUsers">Manage Users</button>
    <button v-on:click="handlerOnclickZonesToUsers">Associate Zones</button>
 </div>
</template>

<script>

export default {
  name: 'Register',
  data(){
    return {
    };
  },
  methods: {
    handlerOnclickZones(){
        this.$router.push({name: 'adminZones'})
    },
    handlerOnclickUsers(){
        this.$router.push({name: 'register'})
    },
    handlerOnclickZonesToUsers(){
        this.$router.push({name: 'zonestousers'})
    },
    logout(){
      localStorage.removeItem('jwt')
      this.$router.push({name: 'login'})
    },
  },
}  
</script>

<style scoped>
  .btns{
    width: 50%;
    margin: auto;
  }
</style>
