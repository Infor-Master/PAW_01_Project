<template>
  <div>
    <br />
    <b-container>
      <b-row>
        <b-col>
          <b-jumbotron>
            <!--<template #header>{{zone.Name}}</template>
              <template #lead>{{jwtDecoded.name}}</template>-->
            <b-list-group>
              <div class="item">
                <b-list-group-item>
                  <h3 class="font-weight-bold">{{ zone.Name }}</h3>
                  <b-progress
                    :value="zone.PplCount"
                    variant="dark"
                    animated
                    show-value
                    :max="zone.Limits"
                    class="mb-3"
                  ></b-progress>
                  <h6 class="font-weight-bold">
                    Max: {{ zone.Limits }} People
                  </h6>
                </b-list-group-item>
              </div>
            </b-list-group>
            <div class="btn">
              <b-button
                click="addPerson"
                class="button"
                size="lg"
                variant="success"
                >Add Person</b-button
              >
            </div>
            <div class="btn2">
              <b-button
                @click="removePerson"
                class="button"
                size="lg"
                variant="danger"
                >Remove Person</b-button
              >
            </div>
          </b-jumbotron>
        </b-col>
      </b-row>
    </b-container>
    <div class="backBtn">
      <button
        v-on:click="handlerOnclickBack"
        class="btn btn-dark btn-lg btn-block"
      >
        Back
      </button>
    </div>
  </div>
</template>

<script>
import settings from "../../configs.json";
import VueJwtDecode from "vue-jwt-decode";

export default {
  data() {
    return {
      ws: null,
      event: "",
      zone: {},
      jwtDecoded: {},
    };
  },
  methods: {
    loadPage() {
      try {
        this.jwtDecoded = VueJwtDecode.decode(localStorage.getItem("jwt"));
      } catch (e) {
        console.error(e);
      }
      this.axios({
        method: "get",
        url: `/zones/id/${this.$route.params.id}`,
        baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
        headers: {
          Authorization: `Bearer ${localStorage.getItem("jwt")}`,
        },
      })
        .then((response) => {
          this.zone = response.data.data;
        })
        .catch((error) => {
          if (error.response) {
            this.$router.push({ name: "zones" });
          }
        });
    },
    addPerson() {
      if (this.zone.PplCount < this.zone.Limits){
        this.axios({
          method: "post",
          url: `/zones/id/${this.$route.params.id}/add`,
          baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
          headers: {
            Authorization: `Bearer ${localStorage.getItem("jwt")}`,
          },
        })
          .then(() => {
            try {
              this.ws.send("updZone"); //manda aviso para atualizar
            } catch (error) {
              this.created(); //volta a criar a coneção
              this.ws.send("updZone"); //manda aviso para atualizar
            }
          })
          .catch((error) => {
            if (error.response) {
              console.error(error.response);
              this.message =
                error.response.status + " - " + error.response.statusText;
            }
        });
      }
    },
    removePerson() {
      if (this.zone.PplCount >0){
        this.axios({
          method: "post",
          url: `/zones/id/${this.$route.params.id}/remove`,
          baseURL: settings.backend.protocol + settings.URL + settings.backend.path,
          headers: {
            Authorization: `Bearer ${localStorage.getItem("jwt")}`,
          },
        })
          .then(() => {
            try {
              this.ws.send("updZone"); //manda aviso para atualizar
            } catch (error) {
              this.created(); //volta a criar a coneção
              this.ws.send("updZone"); //manda aviso para atualizar
            }
          })
          .catch((error) => {
            if (error.response) {
              console.error(error.response);
              this.message =
                error.response.status + " - " + error.response.statusText;
            }
          });
      }
    },
    handlerOnclickBack() {
      this.$router.go(-1);
    },
  },
  mounted() {
    this.loadPage();
  },
  created() {
    this.ws = new WebSocket(settings.sockets.protocol + settings.URL + settings.sockets.path);
    this.ws.onmessage = (event) => {
      this.event = event;
    };
    this.ws.onopen = () => {
      console.log("Successfully connected to the websocket server...");
    };
  },
  watch: {
    event: function (newEvent) {
      if (newEvent.data === "getZone") {
        this.loadPage(); //buscar zona atualizada
        this.event = ""; //reset ao evento
      }
    },
  },
};
</script>

<style scoped>
.backBtn {
  margin: auto;
  padding: 2%;
  width: 30%;
}
.item {
  padding: 1%;
}
.btn {
  padding: 1%;
  width: 55%;
  margin: auto;
}
.btn2 {
  padding: 1%;
  width: 55%;
  margin: auto;
}
</style>
