<template>
  <div>
    <div>
      <v-container class="bg-surface">
        <v-row no-gutters v-for="(sighting, index) in slicedSightings" :key="index">
          <v-col cols="2" sm="1">
            <v-checkbox @update:modelValue="checkboxChanged"></v-checkbox>
          </v-col>
          <v-col cols="10" sm="6">
            <v-sheet class="ma-2 pa-2">
              {{`${sighting.city},${sighting.state ? ' ' + sighting.state +', ' : ' '}${sighting.country}`}}
            </v-sheet>
          </v-col>
          <v-col cols="12" sm="4">
            <v-sheet class="ma-2 pa-2">
              {{sighting.date}}
            </v-sheet>
          </v-col>
          <v-col cols="10" sm="2">
          </v-col>
        </v-row>
        <v-pagination v-if="paginationLength > 0"
                      :length="paginationLength"
                      @update:modelValue="pageChanged"
                      :model-value="currentPage">
        </v-pagination>
      </v-container>
    </div>
  </div>
</template>


<script lang="ts">

import { defineComponent } from 'vue';
import { UfoSighting } from '@/types/types';

export default defineComponent({
  data() {
    return {
      slicedSightings: [],
      perPage: 20,
      paginationLength: 0,
      currentPage: 1,
      allSightings: [],
    };
  },
  props: {
    sightings: {
      required: false,
    }
  },
  watch: {
    sightings(updatedSightings: UfoSighting[]) {
      this.allSightings = updatedSightings;
      if (updatedSightings.length > 0) {
        this.paginationLength = Math.round(updatedSightings.length / this.perPage);
      }
      this.slicedSightings = updatedSightings.slice(0, this.perPage);
    }
  },
  methods: {
    pageChanged: function(newPage) {
      this.currentPage = newPage;
      const start = newPage > 1 ?  (newPage - 1) * this.perPage : 0;
      const end = (start + this.perPage);
      this.slicedSightings = this.allSightings.slice(start, end);
    },
    checkboxChanged: function(value) {
      console.log(value);
    },
  }
});
</script>

<style>
</style>