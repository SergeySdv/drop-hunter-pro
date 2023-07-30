<template>
  <v-card-actions>
    <br/>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="network"
            :rules="[required]"
            :items="networks"
            v-model="item.network"
            :disabled="disabled"
          />
        </v-col>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="token"
            :rules="[required]"
            :items="tokens"
            v-model="item.token"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
      <AmountInput :coin="item.token" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, OkexDepositTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import AmountInput from "@/components/tasks/AmountInput.vue";

export default defineComponent({
  name: "TaskOkexDeposit",
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
      deep: true,
    }
  },

  watch: {
    item: {
      deep: true,
      handler() {
        this.$emit('taskChanged', this.getTask())
      }
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
      item: {
        network: Network.AVALANCHE,
        token: Token.USDT
      } as OkexDepositTask,
      required: (v: any) => !!v || 'required'
    }
  },
  computed: {
    tokens(): Token[] {
      if (this.item.network === Network.ARBITRUM) {
        return [
          Token.USDT,
          Token.USDC,
          Token.ETH,
        ]
      }
      return [
        Token.USDT,
        Token.USDC,
      ]
    },
    networks(): Network[] {
      if (this.item.token === Token.ETH) {
        return [Network.ARBITRUM]
      }

      return [
        Network.AVALANCHE,
        Network.Etherium,
        Network.POLIGON,
        Network.BinanaceBNB,
        Network.OPTIMISM,
        Network.ARBITRUM,
      ]
    }
  },
  methods: {
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        okexDepositTask: this.item,
        taskType: TaskType.OkexDeposit,
        description: ""
      }
    },
    syncTask() {
      if (this.task) {
        if (this.task.okexDepositTask) {
          this.item = this.task.okexDepositTask
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
