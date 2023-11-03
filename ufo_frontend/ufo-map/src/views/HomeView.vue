<template>
  <main>
      <div id="content">
        <Sidebar @sightings-update="markerUpdate" />
        <Map :sightings="foundSightings"/>
      </div>
  </main>
</template>

<script lang="ts">
import Map from '../components/Map/Map.vue';
import Sidebar from '../components/SideBar/SideBar.vue';
import { UfoSighting } from '@/types/types';

export default {
  components: {
    Sidebar,
    Map,
  },
  data() {
    return {
      foundSightings: [] as UfoSighting[],
    };
  },
  methods: {
    markerUpdate(sightings) {
      if (!sightings) {
        this.foundSightings = [];
      } else {
        this.foundSightings = sightings;
      }

    }
  }
};

</script>
<style>
#content  {
  display: flex;
}
</style>