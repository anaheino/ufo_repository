<template>
  <div>
    <div class="date-container">
      <div class="date">
        <h5>Start date</h5>
        <v-text-field type="datetime-local" v-model="startDate" :clearable="true" @update:modelValue="dateChanged"/>
      </div>
    </div>
    <div class="date-container">
      <div class="date">
        <h5>End date</h5>
        <div class="date">
          <v-text-field type="datetime-local" v-model="endDate" :clearable="true" @update:modelValue="dateChanged"/>
        </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts">

import { defineComponent } from 'vue';

export default defineComponent({
  watch: {},
  data() {
    return {
      startDate: null,
      endDate: null,
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
</style>