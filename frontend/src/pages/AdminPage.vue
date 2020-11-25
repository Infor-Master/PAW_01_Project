<template>
 <div>
   <br>
    <div class="btns">
      <button v-on:click="handlerOnclickZones" class="btn btn-dark btn-lg btn-block">Manage Zones</button>
      <button v-on:click="handlerOnclickUsers" class="btn btn-dark btn-lg btn-block">Manage Users</button>
      <button v-on:click="handlerOnclickZonesToUsers" class="btn btn-dark btn-lg btn-block">Associate Zones</button>
    </div>
    <br>
    <div class="myBtn2">
      <b-button variant="danger" @click="$bvToast.show('toast')">
        <h5>Logout</h5>
      </b-button>
      <br>
      <div class="myBtn">
        <b-toast id="toast"  title="Logout?" static no-auto-hide>
          <b-button variant="btn btn-dark btn-lg btn-block" @click.prevent="logout">
            Yes
          </b-button>
        </b-toast>
      </div>
    </div>
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
    width: 30%;
    margin: auto;
    padding: 2%;
  }
  .myBtn{
    padding: 6%;
    margin: auto;
  }
  .myBtn2{
    width: 20%;
    margin: auto;
    margin: auto;
  }
</style>
