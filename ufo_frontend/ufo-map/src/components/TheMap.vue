<template>
  <div id="mapContainer"></div>
</template>

<script lang="ts">
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import { createApp, defineComponent, PropType, ref, watch } from 'vue';
import PopupContent from '@/components/PopupContent.vue';
interface UfoSighting {
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

export default defineComponent({

  props: {
    sightings: {
      type: Array as PropType<UfoSighting[]>,
      required: true,
    },
  },
  watch: {
    sightings(newVal: UfoSighting[], oldVal: UfoSighting[]) {
      if (newVal) {
        this.clearMarkers();
        this.addMarkers(newVal);
      }
    }
  },
  data() {
    return {
      map: null,
      data: null,
      markers: [],
    };
  },
  methods: {
    clearMarkers: function() {

      this.markers.forEach(marker => this.map.removeLayer(marker));
      this.markers = [];
    },
    addMarkers: function(newVal: UfoSighting[]) {
      const sightingsByLocation = this.groupBy(newVal, (sighting) => sighting.latitude.toString() + ' ' + sighting.longitude.toString());
      sightingsByLocation.forEach(multipleSightings => {
        const marker = this.addTooltips(multipleSightings, new L.marker([multipleSightings[0].latitude, multipleSightings[0].longitude]));
        this.markers.push(marker);
        marker.addTo(this.map);
      });
      this.map.setView([newVal[0].latitude, newVal[0].longitude], 5);
    },
    groupBy(array, keyFunc, gimmeDict) {
      const groupedJsons = array.reduce((result, item) => {
        const keyValue = keyFunc(item);
        if (!result[keyValue]) {
          result[keyValue] = [];
        }
        result[keyValue].push(item);
        return result;
      }, {});
      return !gimmeDict ? Object.values(groupedJsons) : groupedJsons;
    },
    addTooltips: function(sightings, marker) {
      const mountEl = document.createElement('div');
      createApp({ extends: PopupContent }, { sightings: sightings, itemsPerPage: 1 }).mount(mountEl);
      const myPopupVueEl = new L.popup().setContent(mountEl);
      marker.bindPopup(myPopupVueEl).openPopup();
      return marker;
    },
  },

  async mounted() {
    this.map = L.map("mapContainer").setView([0, 0], 3);
    L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
      attribution:
          '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(this.map);
    // await this.fetchSightings();
    // this.addMarkers();
  },
  onBeforeUnmount() {
    if (this.map) {
      this.map.remove();
    }
  },
});
</script>

<style scoped>
#mapContainer {
  margin: 2.5% 5% 2.5% 5%;
  width: 90vw;
  height: 90vh;
}
</style>