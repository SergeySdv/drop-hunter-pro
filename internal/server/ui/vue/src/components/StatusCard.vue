<template>
  <v-chip :color="GetStatusColor(status)">{{ GetStatusText(status) }}</v-chip>
</template>

<script lang="ts">

import {defineComponent} from 'vue';

import {GetStatusColor, GetStatusText} from "./helper";

export default defineComponent({
  name: "StatusCard",
  props: {
    status: {
      type: String,
      required: true,
    }
  },
  methods: {
    GetStatusColor,
    GetStatusText

  }
})


</script>


<style>


</style>

