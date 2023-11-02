<template>
  <div id="sidebar">
    <SearchBar @search-sightings="sightingsSearchedFullText" />
<!--<v-container><v-date-picker
        locale="en-in"
        v-model="fromDateVal"
        no-title
        @input="fromDateMenu = false"
        :min="minDate"
    ></v-date-picker>
    </v-container>-->
  </div>
</template>

<script lang="ts">
import SearchBar from '../components/SearchBar.vue';
// import { VDatePicker } from 'vuetify';
export default {
  components: {
    SearchBar,
   // VDatePicker,
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