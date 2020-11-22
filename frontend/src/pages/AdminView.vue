<template>
 <div>
   <form 
    class="admin">
     <h1>Admin page</h1>
     
   </form>
   <span class="border border-dark">
       <h1>Users</h1>
   </span>
   <b-list-group>
        <b-list-group-item v-for="(zone,index) in zones[0]"
        :key="index">
          {{zone.Name}}
        </b-list-group-item>
      </b-list-group>
 </div>
</template>

<script>

export default {
  name: 'Register',
  data(){
    return {
      users: [],
      zones: localStorage.getItem("zones"),
    };
  },
}
</script>