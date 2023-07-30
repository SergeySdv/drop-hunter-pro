<template>
  <v-card-actions>
    <br/>
    <v-container>
      <v-row>
        <v-checkbox
          :disabled="disabled"
          density="compact"
          focused
          name="randomize"
          v-model="random"
          label="randomize"
          @input="CheckboxChanged">
        </v-checkbox>
      </v-row>
      <v-row v-if="!random">
        <v-text-field
          ref="delay-form"
          density="compact"
          variant="outlined"
          v-on:change="inputChanged"
          label="Duration in seconds"
          :rules="[required]"
          v-model="duration"
          :disabled="disabled || random"
          :messages="getDuration"
        />
      </v-row>

      <v-row v-if="random">
        <v-col>
          <v-text-field
            ref="delay-form"
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="min duration in seconds"
            :rules="[required]"
            v-model="minRandom"
            :disabled="disabled"
            :messages="humanDuration(minRandom)"
          />
        </v-col>
        <v-col>
          <v-text-field
            ref="delay-form"
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="max duration in seconds"
            :rules="[required]"
            v-model="maxRandom"
            :disabled="disabled"
            :messages="humanDuration(maxRandom)"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {DelayTask, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import assert from 'assert'
import {humanDuration} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "TaskDelay",
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
    duration: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
    },
    random: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
    },
    minRandom: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
    },
    maxRandom: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
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
      duration: "0",
      random: true,
      minRandom: '180',
      maxRandom: '300',
    }
  },
  methods: {
    required,
    humanDuration,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        delayTask: {
          duration: this.duration ? this.duration : '0',
          random: this.random,
          minRandom: this.minRandom,
          maxRandom: this.maxRandom,
        },
        taskType: TaskType.Delay,
        description: ""
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['delay-form'].validate()

      return valid
    },
    CheckboxChanged() {

    },
    inputChanged() {
      this.duration = this.duration.toString().replace(/^0+/, '');

      return this.validateForm()
    },
    syncTask() {
      if (this.task) {
        if (this.task.delayTask) {
          const t = this.task.delayTask
          this.duration = t.duration
          this.random = t.random
          this.minRandom = t.minRandom ? t.minRandom : this.minRandom
          this.maxRandom = t.maxRandom ? t.maxRandom : this.maxRandom
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  computed: {
    getDuration(): string {
      return humanDuration(this.duration)
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
