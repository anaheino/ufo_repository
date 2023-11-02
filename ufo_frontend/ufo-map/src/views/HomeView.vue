<template>
  <main>
      <div id="content">
        <Sidebar @sightings-update="markerUpdate" />
        <TheMap :sightings="foundSightings"/>
      </div>
  </main>
</template>

<script lang="ts">
import TheMap from '../components/Map/Map.vue';
import Sidebar from '../components/SideBar/SideBar.vue';
import {UfoSighting} from "@/types/types";

export default {
  components: {
    Sidebar,
    TheMap,
  },
  data() {
    return {
      foundSightings: [] as UfoSighting[],
    };
  },
  methods: {
    markerUpdate(sightings) {
      this.foundSightings = sightings;
    }
  }
};

</script>
<style>
#content  {
  display: flex;
}
</style>