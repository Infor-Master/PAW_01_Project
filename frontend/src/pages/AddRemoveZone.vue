<template>
  <div>
    <h1>Manage Zones</h1>
    
    <form class="addZone" @submit.prevent="handlerSubmitAdd">
      <hr />
      <li v-for="zone in this.zones" :key="zone.ID">
        ID: {{ zone.ID }}, Name: {{ zone.Name }}, Lat: {{ zone.Latitude }}, Lng:
        {{ zone.Longitude }}
      </li>
      <hr />

      <p>
        Localização mapa: {{ mapCoordinates.lat }} Latitude,
        {{ mapCoordinates.lng }} Longitude
      </p>
      <gmap-map
        :center="{ lat: coordinates.lat, lng: coordinates.lng }"
        :zoom="7"
        style="width: 1280px; height: 720px; margin: 32px auto"
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
              <h5>
                Participantes: {{ activeZone.PplCount }}/{{ activeZone.Limits }}
              </h5>
              <form
                class="RemoveZone"
                @submit.prevent="handlerSubmitRemove()"
              >
                <button type="submit">Remove</button>
              </form>
            </div>
          </gmap-info-window>
        </div>
        <gmap-marker
          v-for="zone in zones"
          :key="zone.ID"
          :position="getPosition(zone)"
          :clickable="true"
          :draggable="false"
          @click="handleMarkerClicked(zone)"
        >
        </gmap-marker>
      </gmap-map>

      <label>Name</label>
      <input required v-model="name" type="text" placeholder="Name" />
      <p />

      <label>Latitude</label>
      <input
        required
        v-model="mapCoordinates.lat"
        type="number"
        step="0.00000000000000001"
        placeholder="Latitude"
      />
      <p />
      <label>Longitude</label>
      <input
        required
        v-model="mapCoordinates.lng"
        type="number"
        step="0.00000000000000001"
        placeholder="Longitude"
      />
      <p />
      <label>Limit</label>
      <input required v-model="limits" type="number" placeholder="Limit" />
      <p />
      <button type="submit">Add</button>
      {{ this.message }}
    </form>
    <div class="backBtn">
        <button v-on:click="handlerOnclickBack" 
        class="btn btn-dark btn-lg btn-block">Back</button>
      </div>
  </div>
</template>

<script>
import settings from "../settings";

export default {
  name: "addRemoveZone",
  data() {
    return {
      ws: null,
      event: '',
      map: null,
      coordinates: {
        lat: 0,
        lng: 0,
      },
      activeZone: {},
      infoWindowOptions: {
        pixelOffset: {
          width: 0,
          height: -35,
        },
      },
      infoWindowOpened: false,
      zones: [],
      zoneOptions: [],
      selected: "",
      message: this.message,
      name: this.name,
      limits: this.limits,
    };
  },
  methods: {
    getAllZones() {
      this.axios({
        method: "get",
        url: "/zones/all",
        baseURL: settings.baseURL,
        headers: {
          Authorization: `Bearer ${localStorage.getItem("jwt")}`,
        },
      })
        .then((response) => {
          this.zones = [];
          this.zoneOptions = [];
          for (var i in response.data.data) {
            this.zones.push(response.data.data[i]);
            this.zoneOptions.push(response.data.data[i]["Name"]);
          }

          /*  this.aux = response.data["data"];

          for (var j = 0; j < this.aux.length; j++) {
            this.zoneOptions.push(this.aux[j]["Name"]);
            this.zones.push({
              name: this.aux[j]["Name"],
              id: this.aux[j]["ID"],
            });
          } */
        })
        .catch((error) => {
          if (error.response) {
            console.error(error.response);
            this.message =
              error.response.status + " - " + error.response.statusText;
          }
        });
    },
    getPosition(zone) {
      return {
        lat: parseFloat(zone.Latitude),
        lng: parseFloat(zone.Longitude),
      };
    },
    infoWindowPosition() {
      return {
        lat: parseFloat(this.activeZone.Latitude),
        lng: parseFloat(this.activeZone.Longitude),
      };
    },
    handleMarkerClicked(zone) {
      this.activeZone = zone;
      this.infoWindowOpened = true;
    },
    handleInfoWindowClose() {
      this.activeZone = {};
      this.infoWindowOpened = false;
    },
    handlerOnclickBack() {
      this.$router.go(-1);
    },
    handlerSubmitAdd() {
      this.message = "";
      this.axios({
        method: "post",
        url: "/admin/zones",
        baseURL: settings.baseURL,
        data: {
          name: this.name,
          latitude: parseFloat(this.mapCoordinates.lat),
          longitude: parseFloat(this.mapCoordinates.lng),
          limits: parseInt(this.limits),
        },
        headers: {
          Authorization: `Bearer ${localStorage.getItem("jwt")}`,
        },
      })
        .then((response) => {
          console.log(response)
          this.message = "Zone Registado!";
          try {
            this.ws.send('updZone') //manda aviso para atualizar
          } catch (error) {
            this.created()          //volta a criar a coneção
            this.ws.send('updZone') //manda aviso para atualizar
          }
        })
        .catch((error) => {
          if (error.response) {
            console.error(error.response);
            this.message =
              error.response.status + " - " + error.response.statusText;
          }
        });
    },
    handlerSubmitRemove() {
      console.log(this.activeZone);
      var zoneID = this.activeZone["ID"];

      this.message = "";
      this.axios({
        method: "delete",
        url: `admin/zones/` + zoneID,
        baseURL: settings.baseURL,
        headers: {
          Authorization: `Bearer ${localStorage.getItem("jwt")}`,
        },
      })
        .then((response) => {
          console.log(response)
          this.message = "Zone Deleted!";
          this.activeZone = {}
          this.infoWindowOpened = false;
          try {
            this.ws.send('updZone') //manda aviso para atualizar
          } catch (error) {
            this.created()          //volta a criar a coneção
            this.ws.send('updZone') //manda aviso para atualizar
          }
        })
        .catch((error) => {
          if (error.response) {
            console.error(error.response);
            this.message =
              error.response.status + " - " + error.response.statusText;
          }
        });
    },
  },
  mounted() {
    this.$refs.mapRef.$mapPromise.then((map) => (this.map = map));
    this.message = "";
    this.getAllZones();
  },
  created() {
    // get current coords from browser
    this.$getLocation({})
      .then((coordinates) => {
        this.coordinates = coordinates;
      })
      .catch((error) => {
        alert(error);
      });

    this.ws = new WebSocket(settings.socketURL)
    this.ws.onmessage = (event => {
      this.event = event;
    });
    this.ws.onopen = (event => {
        console.log(event)
        console.log("Successfully connected to the websocket server...")
    })
  },
  computed: {
    mapCoordinates() {
      if (!this.map) {
        return {
          lat: 0,
          lng: 0,
        };
      }
      return {
        lat: this.map.getCenter().lat(),
        lng: this.map.getCenter().lng(),
      };
    },
  },
  watch: {
    event: function (newEvent) {
      if(newEvent.data === 'getZone'){
        this.getAllZones();  //buscar zonas atualizadas
        this.event = '';  //reset ao evento
      }
    }
  }
};
</script>


<style scoped>
  .map {
    margin: auto;
    width: 50%;
    padding: 20px;
  }
  .backBtn {
    margin: auto;
    padding: 2%;
    width: 30%;
  }
</style>