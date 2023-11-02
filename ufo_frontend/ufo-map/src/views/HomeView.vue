<template>
  <main>
    <div id="content">
      <Sidebar @sightings-update="markerUpdate" />
      <TheMap :sightings="foundSightings"/>
    </div>
  </main>
</template>

<script lang="ts">
import TheMap from '../components/TheMap.vue';
import Sidebar from '../components/Sidebar.vue';

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