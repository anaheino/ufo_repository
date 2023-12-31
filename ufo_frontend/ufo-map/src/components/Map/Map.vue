<template>
  <div id="mapContainer"></div>
</template>

<script lang="ts">
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import { createApp, defineComponent, PropType, toRaw } from 'vue';
import PopupContent from '@/components/Popup/PopupContent.vue';
import { UfoSighting, CoordPair } from "@/types/types";

export default defineComponent({

  props: {
    sightings: {
      type: Array as PropType<UfoSighting[]>,
      required: true,
    },
  },
  watch: {
    sightings(newVal: UfoSighting[]) {
      if (newVal) {
        this.clearMarkers();
        if (newVal.length) {
          this.addMarkers(newVal);
        }
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
      if (this.markers && this.markers.length) {
        toRaw(this.markers).forEach(marker => toRaw(this.map).removeLayer(marker));
        this.tooltips = [];
        this.markers = [];
      }
    },
    addMarkers: function(newVal: UfoSighting[]) {
      const sightingsByLocation = this.groupBy(newVal, (sighting) => sighting.latitude.toString() + ' ' + sighting.longitude.toString());
      const defaultCoords = this.findMaxAmountOfSightings(sightingsByLocation);
      sightingsByLocation.forEach(multipleSightings => {
        const marker = new L.marker([multipleSightings[0].latitude, multipleSightings[0].longitude]);
        const markerWithTooltip = this.addTooltips(multipleSightings, marker);
        this.markers.push(markerWithTooltip);
        if (multipleSightings.length > 1) {
          const tooltip = L.tooltip({ content: `${multipleSightings.length}`, permanent: true, direction: 'top' })
          marker.bindTooltip(tooltip).openTooltip();
          markerWithTooltip.on('click', () => markerWithTooltip.getTooltip().setOpacity(0));
          markerWithTooltip.getPopup().on('remove', () => markerWithTooltip.getTooltip().setOpacity(.9))
        }
        markerWithTooltip.addTo(toRaw(this.map));
      });
      if (newVal?.length) {
        this.map.setView([defaultCoords.latitude, defaultCoords.longitude], 8);
      }
    },
    groupBy(array, keyFunc, gimmeDict?) {
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
    findMaxAmountOfSightings: function(sightingsByLocation): CoordPair {
      let zoomedCords: CoordPair = { latitude: 0, longitude: 0 };
      sightingsByLocation.reduce((max, current) => {
        if (current.length > max) {
          max = current.length;
          zoomedCords = {
            latitude: current[0].latitude,
            longitude: current[0].longitude,
          }
        }
        return max;
      }, 0);
      return zoomedCords;
    }
  },
  async mounted() {
    this.map = L.map('mapContainer').setView([0, 0], 3);
    L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
      attribution:
          '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(toRaw(this.map));
  },
});
</script>

<style scoped>


#mapContainer {
  width: 90vw;
  height: 90vh;
}
</style>