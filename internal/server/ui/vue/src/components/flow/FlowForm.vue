<template>
  <v-row class="mx-5">
    <v-col cols="4">
      <div style="position: fixed; font-size: 18px; width: 30%">
        <ol>
          <li v-for="(item) in tasks" class="my-2">
            <b>{{ item.taskType }}</b> <span>{{ taskDescription(item) }}</span>
          </li>
        </ol>
      </div>
    </v-col>

    <v-col cols="8">
      <v-row>
        <v-col cols="12" sm="6" md="4">
          <v-text-field
            single-line
            ref="label-input"
            v-model="label"
            label="label"
            density="comfortable"
            variant="outlined"
            :rules="labelRules()"
            @input="labelChanged"
            :disabled="disable"
          ></v-text-field>
        </v-col>
      </v-row>

      <draggable
        v-model="tasks"
        @start="drag=true"
        @end="orderChanged"
        item-key="weight"
        :animation="100"
        :componentData="componentData"
        handle=".handle"
      >
        <template #item="{element}">
          <v-row class="d-flex justify-space-between">
            <v-col>
              <v-card>
                <v-card-title v-if="!disable" class="d-flex justify-space-between px-0">
                  <h4>
                    <v-icon icon="mdi-dots-vertical" class="handle" @click="() => {}"/>
                    {{ `${element.weight}) ${element.taskType}` }}
                  </h4>
                  <v-icon icon="mdi-close" @click="taskDeleted(element.weight)"></v-icon>
                </v-card-title>
                <v-card-title v-else class="px-2">
                  <h4>
                    {{ `${element.weight}) ${element.taskType}` }}
                  </h4>
                </v-card-title>

                <component
                  @taskChanged="taskChanged"
                  style="box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;"
                  :is="element.component"
                  :weight="element.weight"
                  :task="element.task"
                  :disabled="disable"
                />
              </v-card>
            </v-col>
          </v-row>
        </template>
      </draggable>


      <v-row v-if="!disable">
        <v-container>
          <v-card cols="12"
                  style="box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;">
            <v-col cols="12">
              <v-btn class="mx-1 my-1" density="compact" variant="tonal" v-for="taskType in stepTypes"
                     @click="addStep(taskType)">
                {{ taskType }}
              </v-btn>
            </v-col>

          </v-card>
        </v-container>
      </v-row>
    </v-col>
  </v-row>

</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {Task, TaskType} from "@/generated/flow";
import {taskComponentMap, TaskArg, taskTypes, taskProps} from '@/components/tasks/tasks'
import draggable from 'vuedraggable'


export default defineComponent({
  name: "FlowForm",
  computed: {
    taskProps() {
      return taskProps
    }
  },
  components: {
    draggable,
  },
  watch: {
    label: {
      handler() {
        this.$emit('flowChanged', this.label, this.tasks)
      }
    },
    tasks: {
      handler() {
        this.$emit('flowChanged', this.label, this.tasks)
      }
    },
  },
  props: {
    disable: {
      type: Boolean as PropType<boolean>,
      required: false,
    },
    labelValue: {
      type: String as PropType<string>,
      required: false,
    },
    tasksValue: {
      type: Array as PropType<Task[]>,
      required: false
    }
  },
  emits: ['flowChanged'],
  data() {
    return {
      selectErrorMessage: "",
      drag: false,
      label: "",
      selectedTask: TaskType.Delay,
      tasks: [] as TaskArg[],
      stepTypes: taskTypes,
      componentData: {
        type: "transition",
        name: "flip-list"
      }
    }
  },
  methods: {
    taskDescription(item: TaskArg) {
      let base = ''
      if (!item.task) {
        return base
      }
      base += taskProps[item.taskType].descFn(item.task)
      return base
    },
    async addStep(task: TaskType) {
      this.tasks.push({
        component: taskComponentMap.get(task),
        weight: this.tasks.length + 1,
        taskType: task
      });
      this.resolveOrder()
    },
    taskDeleted(i: number) {
      this.tasks = this.tasks.filter(t => t.weight !== i)
      this.resolveOrder()
      this.$emit('flowChanged', this.label, this.tasks)
    },
    orderChanged() {
      this.drag = false
      this.resolveOrder()
    },
    resolveOrder() {
      this.tasks = this.tasks.map((t, i) => {
        t.weight = ++i
        return t
      })
    },
    taskChanged(task: Task) {
      console.log('taskChanged')
      this.updateTask(task)
      this.$emit('flowChanged', this.label, this.tasks)
    },
    updateTask(task: Task) {
      for (let i = 0; i < this.tasks.length; i++) {
        if (this.tasks[i].weight.toString() === task.weight) {
          this.tasks[i].task = task
        }
      }
    },
    labelRules() {
      return [
        (v: string) => this.label ? true : 'label is required'
      ]
    },
    tasksRules() {
      return [
        (v: string) => this.tasks.length ? true : 'add at least one task'
      ]
    },
    async labelChanged(e: InputEvent) {
      await this.validate()
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
  },
  created() {
    if (this.labelValue) {
      this.label = this.labelValue
    }
    if (this.tasksValue) {
      this.tasksValue.forEach((v) => {
        this.tasks.push({
          component: taskComponentMap.get(v.taskType),
          weight: Number(v.weight),
          task: v,
          taskType: v.taskType,
        })
      })
    }
  }
})


</script>


<style>


</style>

