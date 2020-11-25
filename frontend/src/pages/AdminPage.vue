<template>
 <div class="btns">
    <br>
    <button v-on:click="handlerOnclickZones" class="btn btn-dark btn-lg btn-block">Manage Zones</button>
    <button v-on:click="handlerOnclickUsers" class="btn btn-dark btn-lg btn-block">Manage Users</button>
    <button v-on:click="handlerOnclickZonesToUsers" class="btn btn-dark btn-lg btn-block">Associate Zones</button>
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
    }
  },
}  
</script>

<style scoped>
  .btns{
    width: 50%;
    margin: auto;
  }
</style>