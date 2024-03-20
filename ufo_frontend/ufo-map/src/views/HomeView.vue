<template>
  <main>
      <div id="content">
        <Sidebar @sightings-update="markerUpdate" />
        <SightingMap :sightings="foundSightings"/>
      </div>
  </main>
</template>

<script lang="ts">
import SightingMap from '@/components/SightingMap/SightingMap.vue';
import Sidebar from '../components/SideBar/SideBar.vue';

export default {
  components: {
    Sidebar,
    SightingMap,
  },
  data() {
    return {
      foundSightings: [],
    };
  },
  methods: {
    markerUpdate(sightings: []) {
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