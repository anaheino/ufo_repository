<template>
  <v-app>
  <HomeView/>
  </v-app>
</template>

<script lang="ts">
import HomeView from "@/views/HomeView.vue";

export default {
  name: 'App',
  components: {
    HomeView
  }
}
</script>
