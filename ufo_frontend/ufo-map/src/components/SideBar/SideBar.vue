<template>
  <div id="sidebar">
    <v-list>/
    <v-list-item>
      <SearchBar @search-sightings="sightingsSearchedFullText"
                 :dates="dates"
      />
    </v-list-item>
      <v-list-item>
        <DatePicker @date-update="datesUpdated"></DatePicker>
      </v-list-item>
    </v-list>
  </div>
</template>

<script lang="ts">
import SearchBar from './SearchBar/SearchBar.vue';
import { DateRangeString, UfoSighting } from '@/types/types';
import DatePicker from "@/components/DatePicker/DatePicker.vue";

export default {
  components: {
    DatePicker,
    SearchBar,
  },
  data() {
    return {
      fullTextSearchResults: null | [] as UfoSighting[],
      dates: {
        startDate: '',
        endDate: '',
      },
    };
  },
  methods: {
    sightingsSearchedFullText(sightings) {
      this.fullTextSearchResults = sightings;
      this.$emit('sightings-update', sightings);
    },
    datesUpdated(dateObject: DateRangeString) {
      this.dates = {
        startDate: dateObject.startDate ?? '',
        endDate: dateObject.endDate ?? '',
      };
    }
  }
};
</script>

<style scoped>
  #sidebar {
    min-width: 40%;
  }
</style>