<template>
  <div>
    <div class="header">
      <h4>Filter by date</h4>
    </div>
    <div class="date-container">
      <div class="date">
        <v-text-field type="date" label="Start date" v-model="startDate" v-on:input="dateChanged"/>
      </div>'
    </div>
    <div class="date-container">
      <div class="date">
        <v-text-field type="date" label="End date" v-model="endDate" v-on:input="dateChanged"/>
      </div>
    </div>
  </div>
</template>


<script lang="ts">

import { defineComponent } from "vue";

export default defineComponent({
  watch: {},
  data() {
    return {
      startDate: new Date(0).toISOString().substr(0, 10),
      endDate: new Date().toISOString().substr(0, 10),
    };
  },
  methods: {
    dateChanged() {
      this.$emit('date-update', {
        startDate: this.startDate,
        endDate: this.endDate
      });
    }
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