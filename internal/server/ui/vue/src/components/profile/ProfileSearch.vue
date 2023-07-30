<template>
  <v-autocomplete
    chips="true"
    closable-chips
    return-object
    v-model="selectedProfiles"
    :loading="profilesLoading"
    @update:search="searchProfiles"
    :items="suggestedProfiles"
    item-title="num"
    item-value="id"
    :multiple="multiple"
    :rules="[required]"
    ref="label-input"
    label="Нужно выбрать профили, можно выбрать как один, так и сразу несколько, если ввести интервал. Например 1-10"
    density="comfortable"
    variant="outlined"
    clearable="true"
    :disabled="disabled"
    no-data-text="Петушки не нашлись"
    :hide-no-data="hideNoData"
  >
    <template v-slot:chip="{ props, item }">
      <v-chip
        v-bind="props"
        :text="item.value.num"
      ></v-chip>
    </template>
  </v-autocomplete>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {profileService} from "@/generated/services"
import {Profile} from "@/generated/profile";
import {shuffleArray, Timer} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "ProfileSearch",
  emits: ['update:modelValue'],
  props: {
    multiple: {
      type: Boolean,
      required: false,
      default: true
    },
    disabled: {
      type: Boolean,
      required: false
    },
    modelValue: {
      type: Array as PropType<Profile[]>,
      required: true
    }
  },
  data() {
    return {
      profilesLoading: false,
      suggestedProfile: "",
      suggestedProfiles: [] as Profile[],
      timer: new Timer(),
      hideNoData: false,
    }
  },
  computed: {
    selectedProfiles: {
      set(a: Profile[]) {
        this.$emit('update:modelValue', a)
      },
      get(): Profile[] {
        return this.modelValue
      }
    }
  },
  methods: {
    required,
    async searchProfiles(v: string) {
      this.timer.add(100)
      this.timer.cb(async () => {
        this.hideNoData = false
        if (v === "" || v === null || v === undefined) {
          v = ''
        }
        if (v.split("-").length > 1) {
          try {
            this.hideNoData = true
            this.profilesLoading = true
            const res = await profileService.profileServiceSearchProfile({body: {pattern: v}})

            this.suggestedProfiles = []
            this.suggestedProfiles.push(...res.profiles)

            while (this.selectedProfiles.pop()) {

            }
            this.selectedProfiles.push(...this.suggestedProfiles)
            shuffleArray(this.selectedProfiles)
          } finally {
            this.profilesLoading = false
          }
        } else {
          try {
            this.profilesLoading = true
            const res = await profileService.profileServiceSearchProfile({body: {pattern: v}})
            this.suggestedProfiles = []
            this.suggestedProfiles.push(...res.profiles)
          } finally {
            this.profilesLoading = false
          }
        }
      })
    },
  },
  async mounted() {
    const res = await profileService.profileServiceSearchProfile({body: {pattern: ""}})
    this.suggestedProfiles = res.profiles
  }
})


</script>


<style scoped>
</style>

