<template>
  <main>
      <div id="content">
        <Sidebar @sightings-update="markerUpdate" />
        <SightingMap :sightings="foundSightings" @add-sighting="addSighting" @new-coordinates="updateCoordinates"/>
        <NewSightingBar v-if="addingSighting" @new-sighting="newSighting" :coordinates="coordinates" @close-bar="closeBar"/>
      </div>
  </main>
</template>

<script lang="ts">
import SightingMap from '@/components/SightingMap/SightingMap.vue';
import Sidebar from '../components/SideBar/SideBar.vue';
import NewSightingBar from "@/components/NewSighting/NewSightingBar.vue";
import type { UfoSighting} from "@/types/types";
import { defineComponent } from "vue";
import { insertSighting } from "@/views/HomeView";

export default defineComponent({
  components: {
    Sidebar,
    SightingMap,
    NewSightingBar,
  },
  data() {
    return {
      foundSightings: [] as UfoSighting[],
      addingSighting: false,
      coordinates: {} as { latitude: string, longitude: string },
    };
  },
  methods: {
    markerUpdate(sightings: UfoSighting[]) {
      if (!sightings) {
        this.foundSightings = [];
      } else {
        this.foundSightings = sightings;
      }
    },
    addSighting(addSighting: boolean) {
      this.addingSighting = addSighting;
    },
    async newSighting(sighting: UfoSighting) {
      await insertSighting(sighting);
      this.foundSightings = [];
    },
    closeBar(close: boolean) {
      this.addingSighting = !close;
    },
    updateCoordinates(coordinates: { latitude: string, longitude: string }) {
      this.coordinates = coordinates;
    },
  }
});

</script>
<style>
#content  {
  display: flex;
}
</style>