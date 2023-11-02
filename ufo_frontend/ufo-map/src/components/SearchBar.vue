<template>
  <div id="barcontainer">
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


<script>

export default {
  data() {
    return {
      searchTerm: '',
    };
  },
  methods: {
    async performSearch() {
      if (this.searchTerm.length > 0) {
        const response = await fetch(`http://localhost:8080/search?search=${this.searchTerm}`);
        const sightings = await response.json();
        this.$emit('search-sightings', sightings);
      }
    },
  },
};
</script>

<style>
#barcontainer {
  margin: 0% auto;
  width: 100%;
  height: auto;
}
</style>