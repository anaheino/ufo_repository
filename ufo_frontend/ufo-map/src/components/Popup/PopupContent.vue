<template>
  <div v-for="(sighting, index) in sightings" :key="index" :class="'tooltip-container'">
    <div v-if="index === (currentPage - 1)">
      <h1>{{sighting.shape}}</h1>
      <p>{{new Date(sighting.date).toUTCString()}}</p>
      <div>
        <h4>Report location and duration:</h4>
        <p>{{sighting.city}} {{sighting.state}} {{sighting.country}}</p>
        <p>{{sighting.report_date}}</p>
        <p>{{sighting.duration}}</p>
      </div>
      <div>
        <h4>Short description:</h4>
        <p>{{sighting.description}}</p>
        <b v-if="sighting.has_images">This report includes images.</b>
        <a :href="sighting.link" target="_blank">Link to full sighting</a>
      </div>
    </div>
  </div>
  <div v-if="totalPages > 1" id="sighting-paginator" style="margin-top: 5%">
    <div>
      <span>Sighting {{ currentPage }} of {{ totalPages }}</span>
    </div>
    <div :class="'display-flex'">
      <div :class="'left'">
        <button @click="previousPage">Previous</button>
      </div>
      <div :class="'right'">
        <button @click="nextPage">Next</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { UfoSighting } from "@/types/types";

export default defineComponent({
  props: {
    sightings: Array as () => UfoSighting[],
    itemsPerPage: Number,
  },
  data() {
    return {
      currentPage: 1,
    };
  },
  computed: {
    totalPages() {
      return this.sightings?.length ?? 0;
    },
  },
  methods: {
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      } else {
        this.currentPage = 1;
      }
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      } else {
        this.currentPage = this.totalPages;
      }
    },
  },
});
</script>
<style>
  .tooltip-container {
    min-width: 15%;
  }
</style>