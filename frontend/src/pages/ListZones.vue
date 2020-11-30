<template>
  <div class="myDiv">
    <br />
    <div class="myBtn">
      <b-button variant="success" @click="$bvToast.show('toast')">
        {{ this.worker_name }}
      </b-button>
      <div class="pad">
        <b-toast id="toast" title="Logout?" static no-auto-hide>
          <b-button
            variant="btn btn-dark btn-lg btn-block"
            @click.prevent="logout"
          >
            Yes
          </b-button>
        </b-toast>
      </div>
    </div>
    <br /><br /><br />
    <b-jumbotron>
      <template slot="lead">
        <h3>
          <b> Zones </b>
        </h3>
        <p>{{ message }}</p>
        <hr class="my-2" />
        <b-list-group>
          <b-list-group-item
            v-for="(zone, index) in zones[0]"
            :key="index"
            @click="zoneHandler(zone.ID)"
          >
            {{ zone.Name }}
            <b-progress
              :value="zone.PplCount"
              variant="dark"
              animated
              show-value
              :max="zone.Limits"
              class="mb-3"
            ></b-progress>
            <h6 class="font-weight-bold">Max: {{ zone.Limits }} People</h6>
          </b-list-group-item>
        </b-list-group>
      </template>
    </b-jumbotron>
    <div class="mapBtn">
      <b-button variant="btn btn-dark btn-lg btn-block" @click.prevent="map">
        Map
      </b-button>
    </div>
  </div>
</template>

<script>
import settings from "../settings";
import VueJwtDecode from "vue-jwt-decode";

export default {
  data() {
    return {
      ws: null,
      event: "",
      zones: [],
      message: this.message,
      jwtDecoded: {},
      id_worker: 0,
      worker_name: null,
    };
  },
  methods: {
    logout() {
      localStorage.removeItem("jwt");
      this.$router.push({ name: "login" });
    },
    map() {
      this.$router.push({ name: "map" });
    },
    getWorkerZones() {
      try {
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem("jwt"));
        this.id_worker = this.jwtDecoded.id;
        this.worker_name = this.jwtDecoded.name;
        this.axios({
          method: "get",
          url: `/zones/worker`,
          baseURL: settings.baseURL,
          data: {
            id: this.id_worker,
          },
          headers: {
            Authorization: `Bearer ${localStorage.getItem("jwt")}`,
          },
        })
          .then((response) => {
            this.zones = [];
            for (var i in response.data) {
              this.zones.push(response.data[i]);
            }
          })
          .catch((error) => {
            if (error.response) {
              console.error(error.response);
            }
          });
      } catch (e) {
        console.error(e);
      }
    },
    zoneHandler(id) {
      this.$router.push("/zones/id/" + id);
    },
  },
  mounted() {
    this.getWorkerZones();
    console.log(this.zones);
  },
  created() {
    this.ws = new WebSocket(settings.socketURL);
    this.ws.onmessage = (event) => {
      this.event = event;
    };
    this.ws.onopen = (event) => {
      console.log(event);
      console.log("Successfully connected to the websocket server...");
    };
  },
  watch: {
    event: function (newEvent) {
      if (newEvent.data === "getZone") {
        this.getWorkerZones(); //buscar zonas do worker atualizadas
        this.event = ""; //reset ao evento
      }
    },
  },
};
</script>

<style scoped>
.list-group {
  margin-bottom: 15px;
}
.list-group-item:hover {
  background: rgb(162, 211, 168);
  cursor: pointer;
}
.myBtn {
  position: absolute;
  right: 0;
  width: 38%;
  height: 40%;
  margin: auto;
  padding: auto;
}
.myDiv {
  width: 60%;
  align-content: center;
  margin: auto;
}
.pad {
  padding: 5%;
}
.mapBtn {
  width: 30%;
  margin: auto;
}
</style>
