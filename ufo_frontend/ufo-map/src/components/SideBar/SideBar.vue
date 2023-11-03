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
              <h2>Ufo Sighting search</h2>
            </v-list-subheader>
            <div>

              <v-divider></v-divider>

              <v-list-item :class="'sidebar-list-item top-margin'">
                <SearchBar @search-sightings="sightingsSearchedFullText"
                           :dates="dates"
                />
              </v-list-item>
              <v-list-item :class="'sidebar-list-item'">
                <DatePicker @date-update="datesUpdated"></DatePicker>
              </v-list-item>
            </div>
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
  .sidebar {
    width: 40%;
    max-width: 750px;
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