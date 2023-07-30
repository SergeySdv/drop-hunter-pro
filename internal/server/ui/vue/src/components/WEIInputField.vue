<template>
  <v-slider

    min="1000000000000"
    max="5400000000000000"
    step="1000"
    v-model="gasLimit"
    :messages="message"
    :disabled="disabled"
    thumb-label="always"
    label="Max gas cost limit">

    <template v-slot:thumb-label="{ modelValue }">
      {{ modelValue }}
    </template>

  </v-slider>
</template>

<script lang="ts">
import {AmUni, Network, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ca} from "vuetify/locale";
import {helperService, processService} from "@/generated/services";
import {GetTaskSettingsResponse} from "@/generated/process";
import {mdiPencilOutline} from "@mdi/js";
import {Delay, Timer} from "@/components/helper";

export default defineComponent({
  emits: ['update:modelValue'],
  name: "WEIInputField",
  props: {
    modelValue: {
      type: Number as PropType<number | null>,
      required: false
    },
    disabled: {
      type: Boolean,
      required: true,
    },
    network: {
      type: String as PropType<Network>,
      required: true
    }
  },
  computed: {
    gasLimit: {
      get(): number | null {
        return this.modelValue ? this.modelValue : null
      },
      async set(value: number) {
        await this.valueChanged()
        this.$emit('update:modelValue', value)
      }
    }
  },
  data() {
    return {
      timer: new Timer(),
      message: '',
      errorLoading: null as null | string,
      onlyInteger: (v: number) => {
        if (!this.gasLimit) {
          return 'required'
        }

        if (Number.isInteger(Number(this.gasLimit))) {
          return true
        }
        return 'number must be integer'
      },
    }
  },
  methods: {
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['wei-input-form'].validate()

      return valid
    },
    async valueChanged() {
      this.timer.add(100)
      this.timer.cb(async () => {
        if (!this.gasLimit) {
          return
        }

        try {
          this.message = ''
          const res = await helperService.helperServiceCastWei({
            body: {
              wei: this.gasLimit.toString(),
              network: this.network,
            }
          })
          this.message = 'equals to: ' + this.getBalance(res.am)
        } catch (e) {
          this.message = 'error estimating price'
        }
      })
    },
    getBalance(am?: AmUni): string {
      if (!am) {
        return ''
      }

      return `${Number(am.usd).toPrecision(3)} USD`
    },
  },
  async created() {
    await this.valueChanged()
  }
})
</script>

<style scoped>

</style>
