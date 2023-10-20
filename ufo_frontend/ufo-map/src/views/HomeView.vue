<template>
  <main>
    <SearchBar @sighting-search-click="sightingSearch" />
    <TheMap />
  </main>
</template>

<script>
import TheMap from '../components/TheMap.vue';
import SearchBar from '../components/SearchBar.vue';

export default {
  components: {
    SearchBar,
    TheMap,
  },
  data() {
    return {
      searchData: null,
    };
  },
  methods: {
    sightingSearch(searchTerm) {
      this.searchData = searchTerm;
      console.log(this.searchData);
    }
  }
};

</script>
