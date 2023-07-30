<template>
  <a target="_blank" :href="link.snapshot">Snapshot</a>
  <div><b>Space:</b> {{ item.space }}</div>
  <div><b>Status: </b> <span :style="'color: ' + getStatusColor(getStatus)">{{ getStatus }}</span></div>
  <div v-for="(proposal, key) in getProposals">
    <a :href="proposal.link">{{ `proposal #${key + 1}` }} - {{ proposal.status }}</a>
  </div>
</template>

<script lang="ts">
import {OkexDepositTask, SnapshotVoteTask, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {castGasLimitMap, formatTime, humanDuration} from "@/components/helper";
import {ProcessStatus} from "@/generated/process";
import {getStatusColor} from "@/components/tasks/menu/helper";
import {instance} from "@/generated/services";
import {link} from "@/components/tasks/links";


interface Proposal {
  status: string,
  link: string
}

export default defineComponent({
  name: "MenuSnapshotTask",
  props: {
    task: {
      type: Object as PropType<Task>,
      required: true,
    },
    status: {
      type: String as PropType<ProcessStatus>,
      required: true,
    }
  },
  watch: {
    item: {
      handler() {
        if (this.task?.snapshotVoteTask) {
          this.item = this.task.snapshotVoteTask
        }
      },
      deep: true
    }
  },
  data() {
    return {
      proposals: new Map<string, Proposal>(),
      item: {} as SnapshotVoteTask,
    }
  },
  computed: {
    link() {
      return link
    },
    getProposals(): Proposal[] {
      const res: Proposal[] = []
      this.proposals.forEach(v => res.push(v))
      return res
    },
    getStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }

      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }

      if (this.status == ProcessStatus.StatusRunning) {
        return 'processing'
      }

      return 'not started'
    }
  },
  methods: {
    getStatusColor,
    humanDuration,
    formatTime,
    castProposalMap: (a: Object): Map<string, Proposal> => {
      const result = new Map<string, Proposal>()
      if (!a) {
        return result
      }
      for (const name of Object.getOwnPropertyNames(a)) {
        //@ts-ignore
        result.set(name, a[name])
      }

      return result
    }
  },
  async created() {
    if (this.task?.snapshotVoteTask) {
      this.item = this.task.snapshotVoteTask
      this.proposals = this.castProposalMap(this.item.proposal)
    }
  }
})
</script>

<style scoped>

</style>
