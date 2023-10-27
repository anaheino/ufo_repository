<template>
  <main>
    <SearchBar @sightings-searched="sightingsUpdated" />
    <TheMap :sightings="foundSightings"/>
  </main>
</template>

<script lang="ts">
import TheMap from '../components/TheMap.vue';
import SearchBar from '../components/SearchBar.vue';


export interface UfoSighting {
  city: string;
  country: string;
  date: string;
  description: string;
  duration: string;
  has_images: boolean;
  latitude: number;
  link: string;
  longitude: number;
  report_date: string;
  shape: string;
  state: string;
  _id: string;
}

export default {
  components: {
    SearchBar,
    TheMap,
  },
  data() {
    return {
      foundSightings: [] as UfoSighting[],
    };
  },
  methods: {
    sightingsUpdated(sightings) {
      this.foundSightings = sightings;
    }
  }
};

</script>
