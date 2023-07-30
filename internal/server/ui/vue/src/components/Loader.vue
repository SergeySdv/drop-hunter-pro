<template>
  <div class="d-flex justify-center" style="width: 100%; height: 100%">
    <v-img width="200px" height="200px" src="/loader.gif"/>
  </div>
</template>

<script lang="ts">
export default {
  name: "Loader",
  props: {
    id: String,
  },

  data() {
    return {
      value: false
    }
  },
  methods: {
    CheckboxChanged() {
      // @ts-ignore
      this.$emit("CheckboxChanged", this.id, this.value)
    }
  }
}
</script>

<style scoped>

</style>
