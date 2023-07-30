<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
  >

    <template v-slot:activator="{ props }">
      <v-icon v-bind="props"
              @click="click" icon="mdi-check-network-outline"/>

    </template>

    <template v-slot:default="{ props }">

      <v-card>
        <v-card-text>
          <span v-if="loading">stat:
            <v-progress-circular
              indeterminate
              color="primary"
            ></v-progress-circular>
          </span>
          <span v-else>
            {{ stat }}
          </span>
        </v-card-text>
      </v-card>
    </template>
  </v-menu>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {helperService,} from "@/generated/services"

export default defineComponent({
  name: "BtnCheckProxy",
  props: {
    proxy: {
      type: String,
      required: true,
    },
  },

  data() {
    return {
      loading: false,
      menu: false,
      stat: "",
    }
  },
  computed: {},
  methods: {
    async click() {
      try {
        this.loading = true
        const res = await helperService.helperServiceValidateProxy({body: {proxy: this.proxy}})
        this.stat = `${res.ip} ${res.countryName} [${res.countryCode}]`
      } finally {
        this.loading = false
      }
    },
  }
})


</script>


<style>


</style>

