<template>
  <main>
      <div id="content">
        <Sidebar @sightings-update="markerUpdate" />
        <SightingMap :sightings="foundSightings" @add-sighting="addSighting"/>
        <div>
          {{addingSighting}}
        </div>
      </div>
  </main>
</template>

<script lang="ts">
import SightingMap from '@/components/SightingMap/SightingMap.vue';
import Sidebar from '../components/SideBar/SideBar.vue';
import type { UfoSighting} from "@/types/types";
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    Sidebar,
    SightingMap,
  },
  data() {
    return {
      foundSightings: [] as UfoSighting[],
      addingSighting: false,
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
    }
  }
});

</script>
<style>
#content  {
  display: flex;
}
</style>