<template>
 <div>
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
        this.$router.push({name: 'admin/zones'})
    },
    handlerOnclickUsers(){
        this.$router.push({name: 'admin/users'})
    },
    handlerOnclickZonesToUsers(){
        this.$router.push({name: 'admin/zonestousers'})
    }
  },
}  
</script>