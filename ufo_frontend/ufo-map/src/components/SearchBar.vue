<template>
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
</template>


<script>

export default {
  data() {
    return {
      searchTerm: '',
    };
  },
  methods: {
    performSearch() {
      this.$emit('sighting-search-click', this.searchTerm);
    },
  },
};
</script>