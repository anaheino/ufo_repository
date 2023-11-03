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
import { DateRange, UfoSighting } from '@/types/types';
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
        start: '',
        end: '',
      },
    };
  },
  methods: {
    sightingsSearchedFullText(sightings) {
      this.fullTextSearchResults = sightings;
      this.$emit('sightings-update', sightings);
    },
    datesUpdated(dateObject: DateRange) {
      this.dates = {
        start: dateObject.startDate?.toISOString() ?? '',
        end: dateObject.endDate?.toISOString() ?? '',
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