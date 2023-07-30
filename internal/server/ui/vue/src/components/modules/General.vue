<template>
  <Exchange/>
  <Zksync/>
  <LayerZero/>
  <Snapshot/>
  <Orbiter/>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Zksync from "@/components/modules/Zksync.vue";
import Snapshot from "@/components/modules/Snapshot.vue";
import Orbiter from "@/components/modules/Orbiter.vue";
import LayerZero from "@/components/modules/LayerZero.vue";
import Exchange from "@/components/modules/Exchange.vue";


export default defineComponent({
  name: "General",
  components: {Exchange, LayerZero, Orbiter, Snapshot, Zksync},
  created() {
  },
})


</script>


<style>
</style>

