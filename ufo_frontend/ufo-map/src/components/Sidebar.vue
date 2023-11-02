<template>
  <div id="sidebar">
    <SearchBar @search-sightings="sightingsSearchedFullText" />
  </div>
</template>

<script lang="ts">
import SearchBar from '../components/SearchBar.vue';

export default {
  components: {
    SearchBar,
  },
  data() {
    return {
      fullTextSearchResults: [] as UfoSighting[],
    };
  },
  methods: {
    sightingsSearchedFullText(sightings) {
      this.fullTextSearchResults = sightings;
      this.$emit('sightings-update', sightings);
    }
  }
};
</script>

<style>
#sidebar {
  width: 25%;
}
</style>