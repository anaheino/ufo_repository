<template>
  <div id="sidebar">
    <v-list>/
    <v-list-item>
      <SearchBar @search-sightings="sightingsSearchedFullText" />
    </v-list-item>
      <v-list-item>
        <DatePicker :start-date="startDate" :end-date="endDate"></DatePicker>
      </v-list-item>
    </v-list>
  </div>
</template>

<script lang="ts">
import SearchBar from './SearchBar/SearchBar.vue';
import { UfoSighting } from '@/types/types';
import DatePicker from "@/components/DatePicker/DatePicker.vue";

export default {
  components: {
    DatePicker,
    SearchBar,
  },
  data() {
    return {
      startDate: new Date(0).toISOString().substr(0, 10),
      endDate: new Date().toISOString().substr(0, 10),
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

<style scoped>
  #sidebar {
    min-width: 30%;
  }
</style>