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
    hide-details
  ></v-text-field>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, profileService} from "@/generated/services"
import {Delay, Timer} from "@/components/helper";

export default defineComponent({
  name: "ProfileLabelInput",
  props: ['modelValue', 'profileId'],
  emits: ['update:modelValue'],
  data() {
    return {
      loading: false,
      timer: new Timer(),
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
      this.timer.add(1000)

      this.timer.cb(async () => {
        if (!this.labelName) {
          return
        }

        try {
          this.loading = true
          if (!await this.validate()) {
            this.loading = false
            return
          }
        } finally {
          this.loading = false
        }
      })
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
            return true
          }

          try {
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

          } catch (e: any) {
            return "error happened"
          }

        }
      ]
    },
  }
})


</script>


<style>


</style>

