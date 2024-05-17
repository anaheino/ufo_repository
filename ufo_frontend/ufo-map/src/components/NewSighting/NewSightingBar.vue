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
                    v-model="currentSighting.shape"
                    :counter="10"
                    :rules="shapeRules"
                    label="Shape"
                    required
                ></v-text-field>

                <!--<v-select
                    v-model="select"
                    :items="items"
                    :rules="[v => !!v || 'Item is required']"
                    label="Item"
                    required
                ></v-select>-->

                <v-checkbox
                    v-model="checkbox"
                    :rules="[v => !!v || 'You must agree to continue!']"
                    label="Do you agree?"
                    required
                ></v-checkbox>

                <div class="d-flex flex-column">
                  <v-btn
                      class="mt-4"
                      color="success"
                      block
                      @click="validate"
                  >
                    Validate
                  </v-btn>

                  <v-btn
                      class="mt-4"
                      color="error"
                      block
                      @click="reset"
                  >
                    Reset Form
                  </v-btn>

                  <v-btn
                      class="mt-4"
                      color="warning"
                      block
                      @click="resetValidation"
                  >
                    Reset Validation
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
      const { valid } = await this.$refs.form.validate()

      if (valid) alert('Form is valid')
    },
    reset () {
      this.$refs.form.reset()
    },
    resetValidation () {
      this.$refs.form.resetValidation()
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