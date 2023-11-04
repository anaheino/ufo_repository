<template>
  <div>
    <div id="accordion-container">

    <v-container class="bg-surface">
      <v-row no-gutters v-for="(sighting, index) in slicedSightings" :key="index">
        <v-col cols="2" sm="1">
          <v-checkbox @update:modelValue="(val) => checkboxChanged(val, sighting)"></v-checkbox>
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
      <v-row>
        <p>Showing {{start}} to {{end}} of {{allSightings.length}}</p>
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
      allSightings: [],
      currentPage: 1,
      start: 1,
      end: 20,
      perPage: 20,
      paginationLength: 0,
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
        this.paginationLength = Math.ceil(updatedSightings.length / this.perPage);
      }
      this.slicedSightings = updatedSightings.slice(0, this.perPage);
    }
  },
  methods: {
    pageChanged: function(newPage) {
      this.currentPage = newPage;
      const arrayStart = newPage > 1 ?  (newPage - 1) * this.perPage : 0;
      const end = (arrayStart + this.perPage);
      this.slicedSightings = this.allSightings.slice(arrayStart, end);
      this.end = arrayStart + this.slicedSightings.length;
      this.start = arrayStart + 1;
    },
    checkboxChanged: function(checked: boolean, sighting: UfoSighting) {
    },
  }
});
</script>

<style>
#accordion-container > .v-list-item__content {
    background-color: blue;
}

</style>