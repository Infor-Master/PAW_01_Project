<template>
    <div>
      <b-container>
        <b-row >
          <b-col>
            <b-jumbotron header="Local" lead="Person">
              <h1>Number of people: </h1>
              <b-button class="button" size=lg variant="success">Add Person</b-button>
              <p></p>
              <b-button class="button" size=lg variant="danger">Remove Person</b-button>
              <VueTable />
            </b-jumbotron>
          </b-col>
        </b-row>
      </b-container>



      <b-jumbotron header="Local" lead="Person">
        <h1>Number of people: </h1>
        <b-container class="bv-example-row">
          <b-row>
            <b-col>
              <b-button class="button" size=lg variant="success">Add Person</b-button>
           </b-col>
          </b-row>
          <b-row>
            <b-col>
              <b-button class="button" size=lg variant="danger">Remove Person</b-button>
            </b-col>
          </b-row>
        </b-container>
      </b-jumbotron>
</div>
</template>



<script>

import VueTable from '@/components/VueTable'

export default {
  props: VueTable    
}

</script>

