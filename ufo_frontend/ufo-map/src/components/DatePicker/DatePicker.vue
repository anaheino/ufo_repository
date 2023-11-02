<template>
  <div>
    <div class="header">
      <h4>Filter by date</h4>
    </div>
    <div class="date-container">
      <div class="date">
        <v-date-picker
            v-model="startDate"
            @update:modelValue="dateChanged"
        ></v-date-picker>
      </div>
      </div>
    <div class="date-container">
      <div class="date">
        <v-date-picker
            :input-mode="keyboard"
            v-model="endDate"
            @update:modelValue="dateChanged"
        ></v-date-picker>
      </div>
    </div>
  </div>
</template>


<script lang="ts">

import {defineComponent, toRaw} from "vue";

export default defineComponent({
  watch: {},
  data() {
    return {
      startDate: new Date(0).toISOString().substr(0, 10),
      endDate: new Date().toISOString().substr(0, 10),
      date: null,
    };
  },
  methods: {
    dateChanged() {
      this.$emit('date-update', {
        startDate: this.startDate,
        endDate: this.endDate,
      });
    }
  },
});

</script>
<style>
  .date-container {
    margin-left: 5%;
  }
  .header {
    margin-left: 10%;
    margin-bottom: 2.5%;
  }
</style>