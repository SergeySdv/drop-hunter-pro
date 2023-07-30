<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a :href="link.stargatefinance" target="_blank">stargate.finance</a> to see
        available swap
        options
      </div>
      <v-row>
        <v-col>
          <v-select
            ref="stargate-bridge-form"
            density="compact"
            variant="outlined"
            label="from network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.fromNetwork"
            :disabled="disabled"
          />
          <v-select
            density="compact"
            variant="outlined"
            label="from token"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="tokens"
            v-model="item.fromToken"
            :disabled="disabled"
          />
        </v-col>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="to network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="GetToNetworks"
            v-model="item.toNetwork"
            :disabled="disabled"
          />
          <v-select
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="to token"
            :items="tokens"
            :rules="[required]"
            v-model="item.toToken"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
      <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, StargateBridgeTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";

export default defineComponent({
  name: "TaskStargateBridge",
  components: {AmountInput},
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
      networks: [
        Network.BinanaceBNB,
        Network.ARBITRUM,
        Network.OPTIMISM,
        Network.POLIGON,
        Network.AVALANCHE,
      ] as Network[],
      tokens: [Token.USDC, Token.USDT, Token.ETH, Token.STG],
      item: {
        fromNetwork: Network.ARBITRUM,
        toNetwork: Network.OPTIMISM,
        fromToken: Token.USDC,
        toToken: Token.USDC,
        amount: {
          sendAll: true,
        }
      } as StargateBridgeTask,
      required: (v: any) => !!v || 'required'
    }
  },
  methods: {
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        stargateBridgeTask: this.item,
        taskType: TaskType.StargateBridge,
        description: JSON.stringify(this.item),
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['stargate-bridge-form'].validate()

      return valid
    },
    inputChanged() {
      return this.validateForm()
    },
    syncTask() {
      if (this.task) {
        if (this.task.stargateBridgeTask) {
          this.item = this.task.stargateBridgeTask
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  computed: {
    link() {
      return link
    },
    GetToNetworks(): Network[] {
      return this.networks.filter((n) => n !== this.item.fromNetwork);
    },
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
