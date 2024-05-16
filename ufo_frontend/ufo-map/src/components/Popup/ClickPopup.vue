<template>
  <div>
    <div>
      <h4>Probability of sighting:</h4>
      <p>{{probability}}</p>
    </div>
    <div>
      <h5>Did you see something?</h5>
      <div class="report-div">
        <button class="report-button" @click="reportNow">Report Now</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { UfoSighting } from "@/types/types";

export default defineComponent({
  props: {
    probability: String as () => String,
  },
  methods: {
    reportNow() {
      this.$emit('sighting', true);
    },
  },
});
</script>
<style>
.report-button {
  padding: 1px;

  border: 1px solid black;
}
.report-div {
  padding-top: 1.5vw;
}
</style>