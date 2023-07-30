<template>
  <v-dialog v-model="dialog" width="600px" height="auto" :close-on-back="false" :close-on-content-click="false"
            persistent="true">
    <template v-slot:activator="{ props }">
      <v-btn @click="dialog = true" class="mb-4 mx-4">Generate</v-btn>
    </template>

    <v-card>
      <v-tabs
        v-model="tab"
        bg-color="white"
        mandatory
        fixed-tabs
      >
        <v-tab value="1">Number</v-tab>
        <v-tab value="2">Proxy</v-tab>
        <v-tab value="3">Label</v-tab>
      </v-tabs>

      <v-form ref="form">
        <v-window v-model="tab">
          <v-window-item value="1">
            <v-card>
              <v-card-text>
                <v-text-field
                  label="number of accounts to generate"
                  variant="outlined"
                  type="number"
                  v-model="count"
                  :rules="[accountRules]"
                />
              </v-card-text>
              <v-card-actions class="d-flex justify-end">
                <v-btn @click="reset" variant="outlined">Cancel</v-btn>
                <v-btn @click="tab='2'">Next</v-btn>
              </v-card-actions>
            </v-card>
          </v-window-item>
          <v-window-item value="2">
            <v-card-text>
              <v-checkbox
                label="don't use proxy"
                v-model="noProxy"
              />
              <v-textarea
                :label="`enter ${count} proxies in column`"
                variant="outlined"
                type="text"
                :rules="[validateProxy]"
                v-model="proxyText"
                :messages="proxyMessage"
                :disabled="noProxy"
                auto-grow
              />
            </v-card-text>
            <v-card-actions class="d-flex justify-end">
              <v-btn @click="reset" variant="outlined">Cancel</v-btn>
              <v-btn @click="tab='3'">Next</v-btn>
            </v-card-actions>
          </v-window-item>

          <v-window-item value="3">
            <v-card-text>
              <v-text-field
                type="text"
                label="prefix"
                variant="outlined"
                v-model="labelPrefix"
              />
              <v-text-field
                type="number"
                label="start from"
                variant="outlined"
                v-model="labelStarts"
              />
              <v-textarea
                :label="`enter ${count} labels in column`"
                variant="outlined"
                type="text"
                :rules="[validateLabels]"
                v-model="labelText"
                :messages="labelMessage"
                auto-grow
              />
            </v-card-text>
            <v-card-actions class="d-flex justify-end">
              <v-btn @click="reset" variant="outlined">Cancel</v-btn>
              <v-btn @click="generate" variant="flat">Generate</v-btn>
            </v-card-actions>
          </v-window-item>
        </v-window>
      </v-form>
    </v-card>
  </v-dialog>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
//@ts-ignore
import Papa from 'papaparse';
import {helperService, instance, profileService} from "@/generated/services";


export default defineComponent({
  props: {
    demo: {
      required: false,
      type: Boolean,
      default: false
    }
  },
  watch: {
    tab: {
      handler() {
        if (this.tab.toString() === '3')
          this.labelText = this.makeLabels()
      }
    },
    labelStarts: {
      handler() {
        this.labelText = this.makeLabels()
      }
    },
    labelPrefix: {
      handler() {
        this.labelText = this.makeLabels()
      }
    },
    // count: {
    //   handler() {
    //     this.touch()
    //   }
    // }
  },
  emits: ['generated'],
  name: "GenerationProfiles",
  data() {
    return {
      dialog: false,
      count: 10,
      tab: '1',

      proxies: [] as string[],
      noProxy: true,
      proxyText: '',
      proxyMessage: '',
      filledProxy: 0,

      labels: [] as string[],
      labelText: '',
      labelPrefix: 'gen_',
      labelStarts: 1,
      labelMessage: '',
    }
  },
  methods: {
    accountRules(a: any) {

      console.log(a)
      if (!a) {
        return 'required'
      }

      console.log('accountRules: ', a)
      if (Number.isNaN(a)) {
        return 'must be number'
      }

      if (!Number.isInteger(Number(a))) {
        console.log('accountRules: ', 'must be integer')
        return 'must be integer'
      }

      if (this.demo) {
        if (Number(a) > 5) {
          return 'must be less than 5 in demo mode'
        }
      }

      return true
    },
    makeLabels(): string {
      let out = ""
      let count = this.labelStarts
      for (let i = 0; i < this.count; i++) {
        out += `${this.labelPrefix}${count.toString()} \n`
        count++
      }
      return out
    },
    reset() {
      this.tab = '1'
      this.count = this.demo ? 5 : 10
      this.dialog = false

      this.noProxy = true
      this.proxyText = ''
      this.proxyMessage = ''
      this.filledProxy = 0

      this.labelText = ''
      this.labelMessage = ''
      this.labelPrefix = 'gen_'
      this.labelStarts = 1
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async resetValidation(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].reset()

      return valid
    },
    async touch() {
      if (!await this.validate()) {
        return
      }
      await this.resetValidation()
    },
    validateLabels(v: any) {
      this.labels = []
      const result = Papa.parse<string[]>(this.labelText, {skipEmptyLines: true, delimitersToGuess: [';', ',']})

      for (let i = 0; i < result.data.length; i++) {
        let line = result.data[i]
        if (line.length !== 1) {
          return 'line with label: ' + line + " is invalid"
        }
        this.labels.push(line[0])
      }

      return true
    },
    validateProxy(v: any) {

      if (this.noProxy) {
        return true
      }

      this.filledProxy = 0
      this.proxies = []
      const result = Papa.parse<string[]>(this.proxyText, {skipEmptyLines: true, delimitersToGuess: [';', ',']})

      for (let i = 0; i < result.data.length; i++) {
        let line = result.data[i]
        if (line.length !== 1) {
          return 'line with proxy: ' + line + " is invalid"
        }
        this.proxies.push(line[0])
        this.filledProxy++
      }


      if (this.filledProxy < this.count) {
        return `entered less proxies than required ${this.filledProxy}/${this.count}`
      }

      if (this.filledProxy > this.count) {
        return `entered more proxies than required ${this.filledProxy}/${this.count}`
      }


      this.proxyMessage = `entered ${this.filledProxy}/${this.count} proxies`
      return true
    },
    async generate() {
      if (!await this.validate()) {
        return
      }
      const res = await profileService.profileServiceGenerateProfiles({body: {count: this.count.toString()}})

      const result = Papa.parse<string[]>(res.text, {skipEmptyLines: true, delimitersToGuess: [';', ',']})


      for (let i = 0; i < result.data.length; i++) {
        let line = result.data[i]
        if (line.length !== 3) {
          throw new Error('line: ' + line + " is invalid")
        }

        if (!this.noProxy) {
          line[1] = this.proxies[i]
        }
        line[2] = this.labels[i]
      }

      let text = ''
      result.data.forEach((line) => {
        text += `${line[0]}, ${line[1]}, ${line[2]} \r\n`
      })

      this.$emit('generated', text)
      this.reset()
    },
  },
  created() {
    if (this.demo) {
      this.count = 5
    }
  }
})


</script>


<style>


</style>

