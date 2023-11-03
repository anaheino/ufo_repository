<template>
  <div class="sidebar">
    <v-card>
      <v-layout>
        <v-main>
          <v-list
              density="compact"
              nav
          >
            <v-list-subheader :class="'sidebar-list-header'">
              <h1>Search options</h1>
            </v-list-subheader>
            <div>
              <v-divider/>
              <v-list-item :class="'sidebar-list-item top-margin'">
                <SearchBar @search-sightings="sightingsSearchedFullText"
                           :dates="dates"
                />
              </v-list-item>
              <v-list-item :class="'sidebar-list-item'">
                <DatePicker @date-update="datesUpdated"></DatePicker>
              </v-list-item>
            </div>
            <v-list-subheader :class="'sidebar-list-header'">
              <h1>Search results</h1>
            </v-list-subheader>
            <v-divider/>
            <v-list-item :class="'sidebar-list-item top-margin'">
              <SightingAccordion :sightings="searchResults"/>
            </v-list-item>
          </v-list>

        </v-main>

      </v-layout>
    </v-card>
    <v-list>

    </v-list>
  </div>
</template>

<script lang="ts">
import SearchBar from './SearchBar/SearchBar.vue';
import { DateRangeString, UfoSighting } from '@/types/types';
import DatePicker from '@/components/DatePicker/DatePicker.vue';
import SightingAccordion from '@/components/SideBar/SightingAccordion/SightingAccordion.vue';

export default {
  components: {
    DatePicker,
    SearchBar,
    SightingAccordion,
  },
  data() {
    return {
      searchResults: null | [] as UfoSighting[],
      dates: {
        startDate: '',
        endDate: '',
      },
    };
  },
  methods: {
    sightingsSearchedFullText(sightings) {
      this.searchResults = sightings;
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
  .sidebar {
    width: 40%;
    max-width: 1080px;
    min-height: 100%;
  }
  .sidebar-list-header {
    padding-top: 1.5%;
    font-size-adjust: initial;
    padding-bottom: 5%;
  }
  .sidebar-list-item {
    padding: 2.5%;
  }
  .top-margin {
    margin-top:5%;
  }
</style>