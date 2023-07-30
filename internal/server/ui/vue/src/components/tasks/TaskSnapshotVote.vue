<template>
  <v-card-actions>
    <br/>
    <v-container>
      <v-row>
        <a target="_blank" :href="link.snapshot">Snapshot.org</a>
      </v-row>
      <v-row>
        <v-text-field
          ref="task-snapshot-vote-form"
          density="compact"
          variant="outlined"
          v-on:change="inputChanged"
          label="space name (ex: stgdao.eth)"
          :rules="[required]"
          v-model="local.space"
          :disabled="disabled"
        />
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {DelayTask, SnapshotVoteTask, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import assert from 'assert'
import {link} from "@/components/tasks/links";

export default defineComponent({
  name: "TaskSnapshotVote",
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
      deep: true,
    }
  },
  watch: {
    local: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
        }
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
      local: {} as SnapshotVoteTask,
      required: (v: any) => !!v || 'required'
    }
  },
  methods: {
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        snapshotVoteTask: this.local,
        taskType: TaskType.SnapshotVote,
        description: ""
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['task-snapshot-vote-form'].validate()

      return valid
    },

    inputChanged() {
      this.$emit('taskChanged', this.getTask())
      return this.validateForm()
    },
    syncTask() {
      if (this.task) {
        if (this.task.snapshotVoteTask) {
          this.local = this.task.snapshotVoteTask
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
