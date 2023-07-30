<template>
  <v-chip :density="size" variant="tonal" :color="GetStatusColor(status)">
    {{ GetStatusText(status) }}
    <v-menu v-if="msg" v-model="showMsg"
            width="400px" persistent :close-on-content-click="false" :close-on-back="false">
      <template v-slot:activator="{ props }">
        <v-icon v-bind="props" size="20" @click="showMsg = true" icon="mdi-information-outline"/>
      </template>
      <v-card width="300">
        <v-card-title class="d-flex justify-space-between pt-2 pr-2">
          <span>Info</span>
          <v-icon icon="mdi-close" @click="showMsg = false"/>
        </v-card-title>
        <v-card-text>
          {{ `${getMsg}. ${errorRef(getMsg)}` }}
        </v-card-text>
      </v-card>
    </v-menu>
  </v-chip>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';

import {GetStatusColor, GetStatusText} from "./helper";
import {errorRef} from "./error-ref";

export default defineComponent({
  name: "StatusCard",
  data() {
    return {
      showMsg: false
    }
  },
  props: {
    size: {
      type: String as PropType<'comfortable' | 'compact'>,
      required: false,
      default: 'comfortable'
    },
    status: {
      type: String,
      required: true,
    },
    msg: {
      type: String,
      required: false,
      default: undefined,
    },
  },
  computed: {
    getMsg(): string | undefined {
      return this.msg
    }
  },
  methods: {
    errorRef,
    GetStatusColor,
    GetStatusText

  }
})


</script>


<style>


</style>

