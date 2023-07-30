<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a :href="link.testnetbridge" target="_blank">TestNetBridge</a> to see
        available swap options
      </div>
      <v-row>
        <v-col>
          <v-select
            ref="test-net-bridge-swap-form"
            density="compact"
            variant="outlined"
            label="from network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.network"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            ref="delay-form"
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="min amount"
            :rules="getMinRules()"
            v-model="item.minAmount"
            :disabled="disabled"
          />
        </v-col>
        <v-col>
          <v-text-field
            ref="delay-form"
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="max amount"
            :rules="getMaxRules()"
            v-model="item.maxAmount"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, Task, TaskType, TestNetBridgeSwapTask, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {link} from "@/components/tasks/links";

export default defineComponent({
  name: "TaskTestNetBridgeSwap",
  computed: {
    link() {
      return link
    }
  },
  emits: ['taskChanged'],
  props: {
    weight: {
      type: Number,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: true,
    },
    task: {
      type: Object as PropType<Task>,
      required: false,
    }
  },

  watch: {
    item: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
      deep: true
    },
    task: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
          this.syncTask()
        }
      },
      deep: true
    }
  },

  data() {
    return {
      networks: [Network.ARBITRUM, Network.OPTIMISM, Network.Etherium] as Network[],
      item: {
        network: Network.ARBITRUM,
        minAmount: "0.0001",
        maxAmount: "0.001"
      } as TestNetBridgeSwapTask,
      required: (v: any) => !!v || 'required'
    }
  },
  methods: {
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        testNetBridgeSwapTask: this.item,
        taskType: TaskType.TestNetBridgeSwap,
        description: "",
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['test-net-bridge-swap-form'].validate()

      return valid
    },
    inputChanged() {
      return this.validateForm()
    },
    getMinRules() {
      return [
        (v: any) => !!this.item.minAmount || 'required',
        (v: any) => {
          if (Number(this.item.minAmount) > Number(this.item.maxAmount)) {
            return 'min > max'
          }
          return true
        }
      ]
    },
    getMaxRules() {
      return [
        (v: any) => !!this.item.maxAmount || 'required',
        (v: any) => {
          if (Number(this.item.minAmount) > Number(this.item.maxAmount)) {
            return 'min > max'
          }
          return true
        }
      ]
    },
    syncTask() {
      if (this.task) {
        if (this.task.testNetBridgeSwapTask) {
          this.item = this.task.testNetBridgeSwapTask
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
