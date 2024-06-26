<template>
  <div>
    <div :class="'date-header'">
      <h3>Filter by report date</h3>
      <v-divider></v-divider>
    </div>
    <v-row>
    <div class="date-container">
      <v-col cols="1">
      <div class="date">
        Start date
        <v-text-field type="datetime-local" v-model="startDate" :clearable="true" @update:modelValue="dateChanged"/>
      </div>
      </v-col>
    </div>
    </v-row>
    <v-row>
    <div class="date-container">
      <v-col cols="1">
      <div class="date">
        End date
        <div class="date">
          <v-text-field type="datetime-local" v-model="endDate" :clearable="true" @update:modelValue="dateChanged"/>
        </div>
      </div>
      </v-col>
    </div>
    </v-row>
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