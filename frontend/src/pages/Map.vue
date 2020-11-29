<template>
 <div>
   <h1>Coords:</h1>
   <p>Localização atual: {{ coordinates.lat }} Latitude, {{ coordinates.lng }} Longitude</p>
   <p>Localização mapa: {{ mapCoordinates.lat }} Latitude, {{ mapCoordinates.lng }} Longitude</p>
   <gmap-map 
        :center="{lat:coordinates.lat, lng:coordinates.lng}"
        :zoom="7"
        style="width:1280px; height:720px; margin: 32px auto;"
        ref="mapRef"
    >
        <div class="map">
        <gmap-info-window
            :options="infoWindowOptions"
            :position="infoWindowPosition()"
            :opened="infoWindowOpened"
            @closeclick="handleInfoWindowClose()"
        >
            <div>
                <h2>{{ activeZone.Name }}</h2>
                <h5>Participantes: {{ activeZone.PplCount }}/{{ activeZone.Limits }} </h5>
            </div>
        </gmap-info-window>
        </div>
        <gmap-marker
            v-for="(zone) in zones"
            :key="zone.ID"
            :position="getPosition(zone)"
            :clickable="true"
            :draggable="false"
            @click="handleMarkerClicked(zone)"
        >
        </gmap-marker>
   </gmap-map>
    <div class="backBtn">
        <button v-on:click="handlerOnclickBack" class="btn btn-dark btn-lg btn-block">Back</button>
    </div>
 </div>
</template>

<script>
import settings from '../settings'

export default {
  data() {
    return {
        ws: null,
        event: '',
        map : null,
        zones: [],
        activeZone: {},
        infoWindowOptions: {
            pixelOffset: {
                width: 0,
                height: -35
            }
        },
        infoWindowOpened: false,
        coordinates: {
            lat: 0,
            lng: 0
        }
    };
  },
  mounted(){
      this.loadZones();
      this.$refs.mapRef.$mapPromise.then(map => this.map = map);
  },
  created() {
    // get current coords from browser
    this.$getLocation({})
        .then(coordinates => {
            this.coordinates = coordinates
        }).catch(error => {
            alert(error)
        })

    this.ws = new WebSocket(settings.socketURL)
    this.ws.onmessage = (event => {
      this.event = event;
    });
    this.ws.onopen = (event => {
        console.log(event)
        console.log("Successfully connected to the websocket server...")
    })
  },
  watch: {
    event: function (newEvent) {
      if(newEvent.data === 'getZone'){
        this.loadZones();  //buscar zonas atualizadas do worker
        this.event = '';  //reset ao evento
      }
    }
  },
  computed: {
      mapCoordinates() {
          if(!this.map){
              return {
                  lat: 0,
                  lng: 0
              }
          }
          return {
              lat: this.map.getCenter().lat(),
              lng: this.map.getCenter().lng()
          }
      }
  },
  methods: {
    loadZones(){
        this.axios({
            method: 'get',
            url: '/zones/worker',
            baseURL: settings.baseURL,
            headers:{
                'Authorization': `Bearer ${localStorage.getItem('jwt')}`
            }
        }).then((response) =>{
            this.zones = [];
            for(var i in response.data.data){
                this.zones.push(response.data.data[i])
            }
        }).catch(error => {
            if(error.response){
                console.error(error.response);
            }
        });
    },
    handlerOnclickBack(){
        this.$router.go(-1)
    },
    getPosition(zone) {
        return {
            lat: parseFloat(zone.Latitude),
            lng: parseFloat(zone.Longitude)
        }
    },
    infoWindowPosition(){
        return {
            lat: parseFloat(this.activeZone.Latitude),
            lng: parseFloat(this.activeZone.Longitude)
        }
    },
    handleMarkerClicked(zone){
        this.activeZone = zone;
        this.infoWindowOpened = true
    },
    handleInfoWindowClose(){
        this.activeZone = {};
        this.infoWindowOpened = false
    }
  }
}
</script>


<style scoped>
    .map{
        margin: auto;
        width: 50%;
        padding: 20px;
    }
    .backBtn{
        margin: auto;
        padding: 2%;
        width: 30%;
    }
</style>