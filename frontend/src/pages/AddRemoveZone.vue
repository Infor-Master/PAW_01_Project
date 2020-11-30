<template>
  <div class="myDiv">
    <b-jumbotron bg-variant="dark" text-variant="light" border-variant="success">
      <template slot="lead">
        <h3><b>Manage Zones</b></h3>
        <form class="addZone" @submit.prevent="handlerSubmitAdd">
          <hr />
          <h5 v-for="zone in this.zones" :key="zone.ID">
            <b style="color:powderblue;">ID:</b> {{ zone.ID }}, 
            <b style="color:powderblue;">Name:</b>
             {{ zone.Name}}, <b style="color:powderblue;">Lat:</b> {{ zone.Latitude.toFixed(3) }}, 
             <b style="color:powderblue;">Lng:</b>
            {{ zone.Longitude.toFixed(3) }}
          </h5>

          <h5>
            <b style="color:powderblue;">Localização mapa:</b> {{ mapCoordinates.lat.toFixed(3) }} 
            <b style="color:lightgreen;">Latitude</b>,
            {{ mapCoordinates.lng.toFixed(3) }} <b style="color:lightgreen;">Longitude</b>
          </h5>
          <gmap-map
            :center="{ lat: coordinates.lat, lng: coordinates.lng }"
            :zoom="7"
            style="width: auto; height: 720px; margin: 32px auto"
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
                  <h4 style="color:black;"><b>{{ activeZone.Name }}</b></h4>
                  <h5 style="color:black;">
                    <b>Participantes:</b> {{ activeZone.PplCount }}/{{ activeZone.Limits }}
                  </h5>
                  <form
                    class="RemoveZone"
                    @submit.prevent="handlerSubmitRemove()"
                  >
                    <b-button type="submit" class="btn btn-dark btn-lg btn-block"
                    variant="outline-success" style="width:auto;margin:auto;">Remove
          </b-button>
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
          <div>
            <h5>Name</h5>
            <input required v-model="name" type="text" placeholder="Name" />
            <br><br>

            <h5>Latitude</h5>
            <input
              required
              v-model="mapCoordinates.lat"
              type="number"
              step="0.00000000000000001"
              placeholder="Latitude"
            />
            <br><br>
            <h5>Longitude</h5>
            <input
              required
              v-model="mapCoordinates.lng"
              type="number"
              step="0.00000000000000001"
              placeholder="Longitude"
            />
            <br><br>
            <h5>Limit</h5>
            <input required v-model="limits" type="number" placeholder="Limit" />
            <br>
          </div>
          <br>
          <b-button type="submit" class="btn btn-dark btn-lg btn-block"
              variant="outline-primary" style="width:20%;margin:auto;">Add
          </b-button>
          <br>
          {{ this.message }}
        </form>
      </template>
      <div class="backBtn">
      <b-button v-on:click="handlerOnclickBack" 
      class="btn btn-dark btn-lg btn-block"
      variant="outline-success">Back</b-button>
    </div>
    </b-jumbotron>
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
  .myDiv{
    width: 60%;
    align-content: center;
    margin: auto;
  }
  hr { 
    display: block; height: 2px;
    border: 0; border-top: 2px solid rgb(27, 180, 52);
    margin: 1em 0; padding: 0; 
  }
</style>