<template>
  <main>
    <SearchBar @sightings-searched="sightingsUpdated" />
    <TheMap :sightings="foundSightings"/>
  </main>
</template>

<script lang="ts">
import TheMap from '../components/TheMap.vue';
import SearchBar from '../components/SearchBar.vue';

export default {
  components: {
    SearchBar,
    TheMap,
  },
  data() {
    return {
      foundSightings: [] as UfoSighting[],
    };
  },
  methods: {
    sightingsUpdated(sightings) {
      this.foundSightings = sightings;
    }
  }
};

</script>
