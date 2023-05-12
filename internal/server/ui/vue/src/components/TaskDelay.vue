<template>
  <div class="pa-3">
    <h3>{{ `${weight}) Delay` }}</h3>
    <div>
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
              :messages="HumanDuration(minRandom)"
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
              :messages="HumanDuration(maxRandom)"
            />
          </v-col>

        </v-row>
      </v-container>
    </div>
  </div>
</template>

<script lang="ts">
import {DelayTask, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";

export default defineComponent({
  name: "TaskDelay",
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
        this.$emit('stepChanged', this.getTask())
      },
    },
    random: {
      handler() {
        this.$emit('stepChanged', this.getTask())
      },
    },
    minRandom: {
      handler() {
        this.$emit('stepChanged', this.getTask())
      },
    },
    maxRandom: {
      handler() {
        this.$emit('stepChanged', this.getTask())
      },
    }
  },

  data() {
    return {
      duration: "0",
      random: false,
      minRandom: '0',
      maxRandom: '3600',
      required: (v: any) => !!v || 'required'
    }
  },
  methods: {
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
        description: this.random ? JSON.stringify({
          min: this.minRandom,
          max: this.maxRandom
        }) : JSON.stringify(this.duration)
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['delay-form'].validate()

      return valid
    },
    CheckboxChanged() {
      console.log(this.random)
    },
    HumanDuration(s: string): string {
      let totalSeconds = Number(s)
      const hours = Math.floor(totalSeconds / 3600);
      totalSeconds %= 3600;
      const minutes = Math.floor(totalSeconds / 60);
      const seconds = totalSeconds % 60;
      return `${hours} hours ${minutes} minutes ${seconds} seconds`
    },
    inputChanged() {
      this.duration = this.duration.toString().replace(/^0+/, '');

      return this.validateForm()
    },
  },
  computed: {
    getDuration(): string {
      return this.HumanDuration(this.duration)
    }
  },
  created() {
    this.$emit('stepChanged', this.getTask())
    if (this.task) {
      if (this.task.delayTask) {
        const t = this.task.delayTask
        this.duration = t.duration
        this.random = t.random
        this.minRandom = t.minRandom ? t.minRandom : this.minRandom
        this.maxRandom = t.maxRandom ? t.maxRandom : this.maxRandom
      }

    }
  }
})
</script>

<style scoped>

</style>
