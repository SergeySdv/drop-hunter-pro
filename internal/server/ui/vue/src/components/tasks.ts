import {TaskType} from "@/generated/flow";


import TaskStargateBridge from "@/components/TaskStargateBridge.vue";
import TaskDelay from "@/components/TaskDelay.vue";
import TaskExchangeWithdraw from "@/components/TaskExchangeWithdraw.vue";

export const stepComponentMap = new Map<TaskType, any>()

stepComponentMap.set(TaskType.Delay, TaskDelay)
stepComponentMap.set(TaskType.StargateBridge, TaskStargateBridge)
stepComponentMap.set(TaskType.WithdrawExchange, TaskExchangeWithdraw)

export const taskTypes: TaskType[] = []
for (let taskType of stepComponentMap.keys()) {
  taskTypes.push(taskType)
}
