<template>
  <div>
    <div class="header">
      <h4>Filter by date</h4>
    </div>
    <div class="date-container">
      <div class="date">
        <v-text-field type="date" label="Start date" v-model="startDate"/>
      </div>'
    </div>
    <div class="date-container">
      <div class="date">
        <v-text-field type="date" label="End date" v-model="endDate"/>
      </div>
    </div>
  </div>
</template>


<script lang="ts">

import {defineComponent, PropType} from "vue";
import {UfoSighting} from "@/types/types";

export default defineComponent({
  props: {
    startDate: {
      type: Array as PropType<UfoSighting[]>,
      required: true,
    },
    endDate: {
      type: Array as PropType<UfoSighting[]>,
      required: true,
    },
  },
  watch: {},
  data() {
    return {
      map: null,
      data: null,
      markers: [],
    };
  },
});
</script>
<style>
  .date {
    min-width: 50%;
  }
  .date-container {
    display: flex;
    margin-left: 5%;
  }
  .header {
    margin-left: 10%;
    margin-bottom: 2.5%;
  }
</style>