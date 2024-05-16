<template>
  <div id="mapContainer"></div>
</template>

<script lang="ts">
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import { defineComponent, toRaw } from 'vue';
import type { PropType } from 'vue';
import type { UfoSighting, CoordPair } from "@/types/types";
import { groupSightings } from "@/services/commonService";
import { addMarkers, findMaxAmountOfSightings, queryProbability } from "@/components/SightingMap/SightingMap";
import { createApp } from "vue";
import ClickPopup from "@/components/Popup/ClickPopup.vue";

export default defineComponent({
  props: {
    sightings: {
      type: Array as PropType<UfoSighting[]>,
      required: true,
    },
  },
  watch: {
    sightings(updatedSightings: UfoSighting[]) {
      if (updatedSightings) {
        this.clearMarkers();
        if (updatedSightings.length) {
          this.addMarkers(updatedSightings);
        }
      }
    }
  },
  data() {
    return {
      map: {} as L.Map,
      data: null,
      markers: [] as L.Marker[],
      popUpText: 'default value',
      tooltips: [] as L.Tooltip[],
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
    addMarkers(updatedSightings: UfoSighting[]) {
      const sightingsByLocation = groupSightings(updatedSightings, (sighting: UfoSighting) => sighting?.latitude?.toString() + ' ' + sighting?.longitude?.toString()) as UfoSighting[][];
      const defaultCoords = findMaxAmountOfSightings(sightingsByLocation);
      const markers = addMarkers(sightingsByLocation, this.markers as L.Marker[]);
      markers.forEach(markerWithToolTip => markerWithToolTip.addTo(toRaw(this.map)));
      if (updatedSightings?.length) {
        this.map.setView([defaultCoords.latitude, defaultCoords.longitude], 8);
      }
    },
    async getProbability(latitude: string, longitude: string) {
      const data = await queryProbability(latitude, longitude);
      const mountEl = document.createElement('div');
      createApp({ extends: ClickPopup }, { probability: data['probability'] }).mount(mountEl);
      const myPopupVueEl = L.popup().setContent(mountEl);
      myPopupVueEl
            .setLatLng([data.latitude, data.longitude])
            .openOn(toRaw(this.map));
    },
  },
  async mounted() {
    this.map = L.map('mapContainer').setView([0, 0], 3);
    L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
      attribution:
          '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(toRaw(this.map));
    this.map.on('click', (e) => this.getProbability(e.latlng.lat.toFixed(4), e.latlng.lng.toFixed(4)));
  },
});
</script>

<style>
.red-tooltip {
  padding: 0px;
  max-width: 20vw;
  .leaflet-popup-content-wrapper {
    .leaflet-popup-content {
      font-weight: bold;
    }
  }
}

#mapContainer {
  width: 90vw;
  height: 90vh;
}
</style>