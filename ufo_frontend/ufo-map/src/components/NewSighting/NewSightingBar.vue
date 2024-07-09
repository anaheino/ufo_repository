<template>
  <div id="barcontainer">
    <v-card>
      <v-layout>
        <v-main>
          <div>Add sighting</div>
          <button @click="closeBar">Close</button>
            <v-sheet class="mx-auto" width="300">
              <v-form ref="form">
                <v-text-field
                    v-model="coordinates.latitude"
                    label="Latitude"
                    readonly
                ></v-text-field>
                <v-text-field
                    v-model="coordinates.longitude"
                    label="Longitude"
                    readonly
                ></v-text-field>
                <v-text-field
                    v-model="currentSighting.city"
                    :counter="10"
                    :rules="textFieldRules"
                    label="City"
                    required
                ></v-text-field>
                <v-text-field
                    v-model="currentSighting.country"
                    :counter="10"
                    :rules="textFieldRules"
                    label="Country"
                    required
                ></v-text-field>
                <v-text-field
                    v-model="currentSighting.state"
                    :counter="10"
                    label="State"
                ></v-text-field>
                <v-text-field
                    type="datetime-local"
                    v-model="currentSighting.date"
                    required
                />
                <v-text-field
                    v-model="currentSighting.duration"
                    label="Event duration"
                ></v-text-field>
                <v-text-field
                    v-model="currentSighting.shape"
                    :counter="10"
                    :rules="shapeRules"
                    label="Shape"
                    required
                ></v-text-field>
                <v-textarea
                    v-model="currentSighting.description"
                    :rules="textFieldRules"
                    label="Description"
                    required
                ></v-textarea>
                <div class="d-flex flex-column">
                  <v-btn
                      class="mt-4"
                      color="success"
                      block
                      @click="validate"
                  >
                    Submit
                  </v-btn>
                  <v-btn
                      class="mt-4"
                      color="error"
                      block
                      @click="reset"
                  >
                    Reset Form
                  </v-btn>
                </div>
              </v-form>
            </v-sheet>
        </v-main>
      </v-layout>
    </v-card>
    <v-list>
    </v-list>
  </div>
</template>

<script lang="ts">
import type { PropType } from "vue";
import type { UfoSighting } from "@/types/types";

export default {
  props: {
    coordinates: {
      type: {} as PropType<{ latitude: string, longitude: string }>,
      required: true,
    },
  },
  data: () => ({
      currentSighting: {} as UfoSighting,
      shapeRules: [
        v => !!v || 'Shape is required',
      ],
      textFieldRules: [
          v => !!v || 'Item is required',
      ],
      select: null,
      checkbox: false,
    }),
  methods: {
    newSighting(sighting: UfoSighting) {
      this.$emit('new-sighting', sighting);
    },
    closeBar() {
      this.$emit('close-bar', true);
    },
    async validate () {
      const { valid } = await this.$refs.form.validate();
      if (valid) {
        this.newSighting({
          ...this.currentSighting,
          report_date: new Date().toISOString().slice(0, 16),
          latitude: Number(this.coordinates.latitude),
          longitude: Number(this.coordinates.longitude),
          link: '',
          has_images: false,
      });
        this.closeBar();
      }
    },
    reset () {
      this.$refs.form.reset()
    },
  }
};


</script>

<style scoped>

#barcontainer {
  margin: 0% auto;
  width: 50%;
  height: auto;
}
</style>