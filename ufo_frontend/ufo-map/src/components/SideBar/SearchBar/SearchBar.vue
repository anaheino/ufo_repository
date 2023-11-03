<template>
  <div id="barcontainer">
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

import {defineComponent, PropType} from 'vue';
import {SearchTerms} from "@/types/types";

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
        const searchTerms = {
          searchTerm: this.searchTerm,
          dates: this.dates,
        };
        const searchParams = this.formSearchString(searchTerms);
        const response = await fetch(`http://localhost:8080/search?search=${searchParams}`);
        const sightings = await response.json();
        this.$emit('search-sightings', sightings);
      }
    },
    formSearchString(searchTerms: SearchTerms) {
      let searchString = `?search=${searchTerms.searchTerm}`;
      if (searchTerms.dates) {
        if (searchTerms.dates.start) {
          searchString = searchString.concat(`&startDate=${searchTerms.dates.start}`)
        }
        if (searchTerms.dates.end) {
          searchString = searchString.concat(`&endDate=${searchTerms.dates.end}`)
        }
      }
      console.log(searchString)
      return searchString;
    }
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