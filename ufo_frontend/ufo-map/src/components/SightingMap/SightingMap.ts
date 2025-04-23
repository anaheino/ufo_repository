import type { UfoSighting, CoordPair } from "@/types/types";
import L from "leaflet";
import { createApp } from "vue";
import PopupContent from "@/components/Popup/SightingPopup.vue";
import {backendUrl} from "@/constants/contants";

export const addMarkers = (sightingsByLocation: UfoSighting[][], markers: L.Marker[]) => {
    return sightingsByLocation.map((multipleSightings: UfoSighting[]) => {
        const marker = L.marker([multipleSightings[0].latitude, multipleSightings[0].longitude]);
        const markerWithTooltip = addTooltips(multipleSightings, marker);
        markers.push(markerWithTooltip);
        if (multipleSightings.length > 1) {
            const tooltip = L.tooltip({ content: `${multipleSightings.length}`, permanent: true, direction: 'top' })
            marker.bindTooltip(tooltip).openTooltip();
            markerWithTooltip.on('click', () => markerWithTooltip?.getTooltip()?.setOpacity(0));
            markerWithTooltip?.getPopup()?.on('remove', () => markerWithTooltip?.getTooltip()?.setOpacity(.9))
        }
        return markerWithTooltip;
    });
};

const addTooltips = (sightings: UfoSighting[], marker: L.Marker): L.Marker => {
    const mountEl = document.createElement('div');
    createApp({ extends: PopupContent }, { sightings: sightings, itemsPerPage: 1 }).mount(mountEl);
    const myPopupVueEl = L.popup().setContent(mountEl);
    marker.bindPopup(myPopupVueEl).openPopup();
    return marker;
};

export const findMaxAmountOfSightings = (sightingsByLocation: UfoSighting[][]): CoordPair => {
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
};

export const queryProbability = async (latitude: string, longitude: string)=> {
    const requestData = {
        'latitude':latitude,
        'longitude': longitude
    };
    try {
        const response = await fetch(`${backendUrl}/probability`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData),
        });
        return await response.json();
    } catch (error) {
        console.error('Error:', error);
    }
}