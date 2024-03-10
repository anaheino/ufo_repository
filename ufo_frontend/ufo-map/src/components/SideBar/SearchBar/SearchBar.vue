<template>
  <div id="barcontainer">
    <h3>Free text search</h3>
    <v-card>
      <v-toolbar
          dense
          floating
      >
        <v-text-field
            v-model="searchTerm"
            hide-details
            single-line
        ></v-text-field>
        <v-btn @click="performSearch" icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </v-toolbar>
    </v-card>
  </div>
</template>


<script lang="ts">

import {defineComponent, toRaw} from 'vue';
import type {DateRangeString, SearchTerms} from "@/types/types";

export default defineComponent({
  props: {
    dates: {
      required: false,
    }
  },
  data() {
    return {
      searchTerm: '',
    };
  },
  methods: {
    async performSearch() {
      if (this.searchTerm.length > 0) {
        const searchTerms: SearchTerms = {
          searchTerm: this.searchTerm,
          dates: toRaw(this.dates) as DateRangeString,
        };
        const searchParams = this.formSearchString(searchTerms);
        const response = await fetch(`http://localhost:8080/search?searchTerm=${searchParams}`);
        const sightings = await response.json();
        this.$emit('search-sightings', sightings);
      }
    },
    formSearchString(searchTerms: SearchTerms) {
      let searchString = `${searchTerms.searchTerm}`;
      if (searchTerms.dates) {
        if (searchTerms.dates.startDate) {
          searchString = searchString.concat(`&startDate=${searchTerms.dates.startDate}`)
        }
        if (searchTerms.dates.endDate) {
          searchString = searchString.concat(`&endDate=${searchTerms.dates.endDate}`)
        }
      }
      return searchString;
    },
  },
});
</script>

<style>
#barcontainer {
  margin: 0% auto;
  width: 100%;
  height: auto;
}
</style>