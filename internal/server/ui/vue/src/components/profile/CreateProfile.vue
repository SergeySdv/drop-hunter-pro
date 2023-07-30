<template>
  <div class="mx-8 my-8" style="min-height: 500px">
    <v-container>
      <div v-if="!demo" class="text-h4 my-5">Create profiles</div>
      <v-btn :loading="loading" @click="validate" class="mb-4 mx-4">Validate</v-btn>
      <v-btn v-if="!demo" :loading="loading" @click="save" class="mb-4 mx-4">Create</v-btn>

      <GenerationProfiles @generated="generated" :demo="demo"/>
      <v-textarea
        auto-grow
        density="compact"
        clearable="true"
        variant="outlined"
        placeholder="<metamask-pk> required, <proxy> optional, <custom label> optional] ..."
        v-model="text"
        :error-messages="errorMessage"
        :disabled="loading"
      />
      <div v-if="valid">Number of profiles: <b>{{ items.length }}</b></div>
      <div v-if="mode">Mode: <b>{{ mode }}</b></div>

      <br/>

      <div v-for="item in items">
        <span :style="`color: ${getStatusColor(item)}`">{{ `[ status: ${item.status} ]` }}</span>
        {{
          `- pk: ${item.pk} proxy: ${item.proxy} label: ${item.label}`
        }}
      </div>
    </v-container>
  </div>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
//@ts-ignore
import Papa from 'papaparse';
import {helperService, instance, profileService} from "@/generated/services";
import GenerationProfiles from "@/components/profile/GenerationProfiles.vue";


interface Item {
  pk: string
  proxy?: string
  label?: string

  status?: string
}

export default defineComponent({
  props: {
    demo: {
      required: false,
      type: Boolean,
      default: false
    }
  },
  components: {GenerationProfiles},
  name: "CreateProfile",
  data() {
    return {

      loading: false,
      mode: '',
      dialog: false,
      items: [] as Item[],
      validError: '',
      valid: false,
      generateCount: 10,
      errorMessage: '',
      text: `
    eee5112gd2433f855896002e9cb9c8c1eeb3d8f8dac388d4901dbbf3dec683aa; 123.456.67.89:login:password; my_label_1
    eee5119x3h243f855896004e9cb9c8c1eeb3d8f8dac388d4901dbbf3dec683aa; ""; my_label_2
    eee51g23s62h3f855896004e9cb9c8c1eeb3d8f8dac388d4901dbbf3dec683aa; 123.456.67.89:login:password; ""`,
    }
  },
  methods: {
    generated(text: string) {
      this.text = text
    },
    getStatusColor(item: Item): string {
      if (item.status === 'ok') {
        return 'green'
      }
      if (item.status === 'validating') {
        return 'blue'
      }
      if (item.status === 'saving') {
        return 'blue'
      }
      if (item.status === 'ready') {
        return 'grey'
      }
      return 'red'
    },
    print(data: string[][]) {
      this.text = this.format(data)
    },
    async generate() {
      const res = await profileService.profileServiceGenerateProfiles({body: {count: this.generateCount.toString()}})
      this.text = res.text
      this.dialog = false
    },
    async save() {
      this.loading = true
      this.mode = 'validate and save'
      this.items = []
      try {
        this.valid = false
        this.validError = ""
        const data = this.preparse(this.text)
        data.forEach((line) => {
          this.items.push({
            pk: line[0],
            proxy: line[1] ? line[1] : '',
            label: line[2] ? line[2] : '',
            status: 'ready'
          })
        })
        this.print(data)
        this.valid = true
      } catch (e) {
        this.validError = "invalid input"
        this.valid = false
        this.loading = false
        return
      }

      for (let i = 0; i < this.items.length; i++) {
        let item = this.items[i]
        try {
          item.status = 'saving'
          const res = await profileService.profileServiceCreateProfile({
            body: {
              proxy: item.proxy,
              mmskPk: item.pk,
              label: item.label ? item.label : ''
            }
          })
          item.status = 'ok'
        } catch (e) {
          item.status = 'error saving: ' + e
        }
      }
      this.loading = false
    },
    async validate() {
      this.loading = true
      this.mode = 'validate only'
      this.items = []
      try {
        this.valid = false
        this.validError = ""
        const data = this.preparse(this.text)
        data.forEach((line) => {
          this.items.push({
            pk: line[0],
            proxy: line[1] ? line[1] : '',
            label: line[2] ? line[2] : '',
            status: 'ready'
          })
        })
        this.print(data)
        this.valid = true
      } catch (e) {
        this.validError = "invalid input"
        this.valid = false
        this.loading = false
        return
      }

      for (let i = 0; i < this.items.length; i++) {
        let item = this.items[i]
        this.items[i] = await this.validateItem(item)
      }

      this.loading = false
    },
    async validateItem(item: Item): Promise<Item> {
      item.status = 'validating'

      if (item.proxy) {
        try {
          const stat = await helperService.helperServiceValidateProxy({body: {proxy: item.proxy}})
          if (!stat.valid) {
            item.status = stat.errorMessage
            return item
          } else {
          }
        } catch (e) {
          item.status = 'proxy error'
          return item
        }
      }

      try {
        const pkStat = await helperService.helperServiceValidatePk({body: {pk: item.pk}})
        if (!pkStat.valid) {
          item.status = 'pk invalid'
          return item
        }
      } catch (e) {
        item.status = 'pk invalid'
        return item
      }

      item.status = 'ok'

      return item
    },
    preparse(text: string): string[][] {
      text = text.replaceAll(" ", "")
      this.errorMessage = ''
      const result = Papa.parse<string[]>(text, {skipEmptyLines: true, delimitersToGuess: [';', ',']})
      result.data.forEach((line: string[]) => {
        if (line.length !== 3) {
          throw new Error('line length is not 3 words')
        }
      })
      return result.data
    },
    format(data: string[][]): string {
      let formatted = ""
      data.forEach((line: string[]) => {
        formatted += line.join(";")
        formatted += '\r'
      })
      return formatted
    }
  }
})


</script>


<style>


</style>

