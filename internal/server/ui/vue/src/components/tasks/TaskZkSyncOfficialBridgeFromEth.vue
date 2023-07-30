<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a :href="link.officialZkSyncBridge" target="_blank">Official zksync bridge</a> to see
        available options
        options
      </div>
      <v-row>
        <v-col cols="6">
          <v-select
            density="compact"
            variant="outlined"
            label="from"
            :rules="[required]"
            v-model="Network.Etherium"
            :disabled="true"
            hide-details
          />
        </v-col>
        <v-col cols="6">
          <v-select
            density="compact"
            variant="outlined"
            label="to"
            :rules="[required]"
            v-model="Network.ZKSYNCERA"
            :disabled="true"
            hide-details
          />
        </v-col>
      </v-row>
      <AmountInput :coin="Token.ETH" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {
  Network,
  SyncSwapTask,
  Task,
  TaskType,
  Token,
  ZkSyncOfficialBridgeFromEthereumTask,
  ZkSyncOfficialBridgeToEthereumTask
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {required} from "@/components/tasks/menu/helper";
import {link} from "@/components/tasks/links";

export default defineComponent({
  name: "TaskZkSyncOfficialBridgeFromEth",
  computed: {
    link() {
      return link
    },
    Network() {
      return Network
    },
    Token() {
      return Token
    }
  },
  components: {AmountInput, WEIInputField},
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
      networks: [Network.ZKSYNCERA] as Network[],
      item: {
        network: Network.ZKSYNCERA,
        amount: {
          sendAll: true
        },
      } as ZkSyncOfficialBridgeFromEthereumTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        zkSyncOfficialBridgeFromEthereumTask: this.item,
        taskType: TaskType.ZkSyncOfficialBridgeFromEthereum,
        description: "",
      }
    },
    syncTask() {
      if (this.task) {
        if (this.task.zkSyncOfficialBridgeFromEthereumTask) {
          this.item = this.task.zkSyncOfficialBridgeFromEthereumTask
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
