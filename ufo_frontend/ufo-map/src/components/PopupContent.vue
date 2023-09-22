<template>
  <h1>{{sighting.shape}}</h1>
  <p>{{new Date(sighting.date).toUTCString()}}</p>
  <div>
    <h4>Report location and duration:</h4>
    <p>{{sighting.city}} {{sighting.state}} {{sighting.country}}</p>
    <p>{{sighting.report_date}}</p>
    <p>{{sighting.duration}}</p>
  </div>
  <div>
    <h4>Short description:</h4>
    <p>{{sighting.description}}</p>
    <b v-if="sighting.has_images">This report includes images.</b>
    <a :href="sighting.link" target="_blank">Link to full sighting</a>
  </div>
</template>

<script>
export default {
  props: ['sighting']
}
</script>