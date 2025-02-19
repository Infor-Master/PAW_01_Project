import Vue from 'vue'
import VueRouter from 'vue-router'
import VueGeolocation from 'vue-browser-geolocation'
import * as VueGoogleMaps from 'vue2-google-maps'
import axios from 'axios'
import VueAxios from 'vue-axios'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import App from './App.vue'
import Login from './pages/Login.vue'
import Map from './pages/Map.vue'
import ListZones from './pages/ListZones.vue'
import Register from './pages/Register.vue'
import LocalControl from './pages/LocalControl.vue'
import AddRemoveZone from './pages/AddRemoveZone.vue'
import AdminPage from './pages/AdminPage.vue'
import VueJwtDecode from 'vue-jwt-decode'
import Association from './pages/Association.vue'

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueAxios, axios)
Vue.use(BootstrapVue)
Vue.use(VueGeolocation)
Vue.use(VueJwtDecode)
Vue.use(VueGoogleMaps,  {
  load: {
    key: 'AIzaSyDvzBG1YC-EqqOau_4BMMAgK7p-t9nrYjE',
    autobindAllEvents: true,
  },
  installComponents: true,
})


const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/map',
      name: 'map',
      component: Map,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else{
          next()
        }
      }
    },
    {
      path: '/zones',
      name: 'zones',
      component: ListZones,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else{
          next()
        }
      }
    },
    {
      path: '/admin/users',
      name: 'register',
      component: Register,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else {
          try {
            let decoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
            if(!decoded.admin){
              next({ name: 'login' })
            }else{
              next()
            }
          } catch (error) {
            next({ name: 'login' })
          }
        }
      }
    },
    {
      path: '/zones/id/:id',
      name: 'localControl',
      component:LocalControl,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else{
          next()
        }
      }
    },
    {
      path: '/admin/zones',
      name: 'adminZones',
      component: AddRemoveZone,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else {
          try {
            let decoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
            if(!decoded.admin){
              next({ name: 'login' })
            }else{
              next()
            }
          } catch (error) {
            next({ name: 'login' })
          }
        }
      }
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminPage,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else {
          try {
            let decoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
            if(!decoded.admin){
              next({ name: 'login' })
            }else{
              next()
            }
          } catch (error) {
            next({ name: 'login' })
          }
        }
      }
    },
    {
      path: '/admin/zonestousers',
      name: 'zonestousers',
      component: Association,
      beforeEnter: (to, from, next) => {
        if (localStorage.getItem('jwt') === null){
          next({ name: 'login' })
        }else {
          try {
            let decoded = VueJwtDecode.decode(localStorage.getItem('jwt'))
            if(!decoded.admin){
              next({ name: 'login' })
            }else{
              next()
            }
          } catch (error) {
            next({ name: 'login' })
          }
        }
      }
    }
  ]
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
