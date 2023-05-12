<template>
  <v-checkbox
    style="height: 40px"
    density="compact"
    focused
    v-model="value"
    @input="CheckboxChanged"></v-checkbox>
</template>

<script lang="ts">
export default {
  name: "CheckBox",
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
