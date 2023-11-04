<template>
  <div>
    <div :class="'date-header'">
      <h3>Filter by report date</h3>
      <v-divider></v-divider>
    </div>
    <div class="date-container">
      <div class="date">
        Start date
        <v-text-field type="datetime-local" v-model="startDate" :clearable="true" @update:modelValue="dateChanged"/>
      </div>
    </div>
    <div class="date-container">
      <div class="date">
        End date
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
.date-header {
  margin-bottom: 5%;
}
</style>