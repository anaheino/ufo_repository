<template>
  <div id="barcontainer">
    <h3>Free text search</h3>
    <v-card>
      <v-toolbar
          dense
          floating
      >
        <v-text-field
            v-model="searchTerm"
            hide-details
            single-line
        ></v-text-field>
        <v-btn @click="performSearch" icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </v-toolbar>
    </v-card>
  </div>
</template>


<script lang="ts">

import { defineComponent } from 'vue';
import type { PropType } from 'vue';
import type { DateRangeString } from "@/types/types";
import { querySightings } from "@/components/SideBar/SearchBar/SearchBar";

export default defineComponent({
  props: {
    dates: {
      type: Object as PropType<DateRangeString>,
      required: false,
    }
  },
  data() {
    return {
      searchTerm: '',
    };
  },
  methods: {
    async performSearch() {
      this.$emit('search-sightings', await querySightings(this.searchTerm, this.dates ?? {}));
      }
    }
});
</script>

<style>
#barcontainer {
  margin: 0% auto;
  width: 100%;
  height: auto;
}
</style>