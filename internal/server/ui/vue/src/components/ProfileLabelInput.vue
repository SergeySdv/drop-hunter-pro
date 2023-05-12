<template>
  <v-text-field
    ref="label-input"
    v-model="labelName"
    label="label"
    density="comfortable"
    variant="outlined"
    :rules="labelRules()"
    @input="labelChanged"
    :loading="loading"
    :disabled="loading"
  ></v-text-field>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, profileService} from "@/generated/services"
import {Delay} from "@/components/helper";

export default defineComponent({
  name: "ProfileLabelInput",
  props: ['modelValue', 'profileId'],
  emits: ['update:modelValue'],
  data() {
    return {
      loading: false,
    }
  },
  computed: {
    labelName: {
      get(): string {
        return this.modelValue
      },
      async set(value: string) {
        await this.labelChanged()
        this.$emit('update:modelValue', value)
      }
    }
  },
  methods: {
    async labelChanged() {
      const before = this.labelName
      await Delay(1000)
      if (this.labelName !== before) {
        return
      }

      this.loading = true
      if (!await this.validate()) {
        this.loading = false
        return
      }

      this.loading = false
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
    labelRules() {
      return [
        async (v: string) => {
          if (!this.labelName) {
            return 'Label required'
          }

          const res = await profileService.profileServiceValidateLabel({
            body: {
              label: this.labelName,
              profileId: this.profileId
            }
          })

          if (res.valid) {
            return true
          } else {
            return "label already exist. choose another label"
          }
        }
      ]
    },
  }
})


</script>


<style>


</style>

