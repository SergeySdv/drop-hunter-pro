<template>
  <div class="text-center ma-2">
    <v-snackbar
      variant="flat"
      v-model="getSnackBar"
      :color="isError? 'red' : 'green'"
      position="fixed"
      location="bottom"
    >
      {{ getText }}

      <template v-slot:actions>
        <v-btn
          variant="text"
          @click="close"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script lang="ts">
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";

export default {
  name: "Snackbar",
  computed: {
    ...mapStores(useSysStore),
    getSnackBar() {
      return this.sysStore.snackbar
    },
    getText() {
      return this.sysStore.snackbarText
    },
    isError() {
      return this.sysStore.error
    }
  },
  data() {
    return {}
  },
  methods: {
    close() {
      this.sysStore.$patch((s) => {
        s.snackbar = false
      })
    }
  }
}
</script>

<style scoped>

</style>
