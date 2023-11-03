<template>
  <div class="container">
    <ul>
      <table>
      <tr v-for="(sighting, index) in sightings" :key="index">
        <td>{{sighting.city}}</td>
        <td>{{sighting.date}}</td>
      </tr>
      </table>

    </ul>
  </div>
</template>


<script lang="ts">

import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    sightings: {
      required: false,
    }
  },
});
</script>

<style>
.container {

}
</style>