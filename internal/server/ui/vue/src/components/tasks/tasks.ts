import {Network, Task, TaskType} from "@/generated/flow";
import {flow_Flow} from "@/generated/process";


import TaskStargateBridge from "@/components/tasks/TaskStargateBridge.vue";
import TaskDelay from "@/components/tasks/TaskDelay.vue";
import TaskExchangeWithdraw from "@/components/tasks/TaskExchangeWithdraw.vue";
import TaskOkexDeposit from "@/components/tasks/TaskOkexDeposit.vue";
import TaskTestNetBridgeSwap from "@/components/tasks/TaskTestNetBridgeSwap.vue";
import TaskSnapshotVote from "@/components/tasks/TaskSnapshotVote.vue";
import TaskSyncSwap from "@/components/tasks/TaskSyncSwap.vue";
import MenuTaskSyncSwap from "@/components/tasks/menu/MenuTaskSyncSwap.vue";
import MenuDelayTask from "@/components/tasks/menu/MenuDelayTask.vue";
import MenuExchangeWithdraw from "@/components/tasks/menu/MenuExchangeWithdraw.vue";
import MenuOkexDeposit from "@/components/tasks/menu/MenuOkexDeposit.vue";
import MenuSnapshotTask from "@/components/tasks/menu/MenuSnapshotTask.vue";
import MenuTaskStargateBridge from "@/components/tasks/menu/MenuTaskStargateBridge.vue";
import MenuTaskTestNetBridge from "@/components/tasks/menu/MenuTaskTestNetBridge.vue";
import TaskZkSyncOfficialBridgeToEth from "@/components/tasks/TaskZkSyncOfficialBridgeToEth.vue";
import TaskOrbiterBridge from "@/components/tasks/TaskOrbiterBridge.vue";
import MenuTaskOrbiterBridge from "@/components/tasks/menu/MenuTaskOrbiterBridge.vue";
import MenuZkSyncOfficialBridgeToEth from "@/components/tasks/menu/MenuZkSyncOfficialBridgeToEth.vue";
import TaskZkSyncOfficialBridgeFromEth from "@/components/tasks/TaskZkSyncOfficialBridgeFromEth.vue";
import MenuZkSyncOfficialBridgeFromEth from "@/components/tasks/menu/MenuZkSyncOfficialBridgeFromEth.vue";
import TaskWethSwap from "@/components/tasks/TaskWethSwap.vue";
import MenuTaskWethSwap from "@/components/tasks/menu/MenuTaskWethSwap.vue";
import TaskMuteIoSwap from "@/components/tasks/TaskMuteIoSwap.vue";
import MenuTaskMuteioSwap from "@/components/tasks/menu/MenuTaskMuteioSwap.vue";
import {Process} from "@/generated/process";
import {humanDuration} from "@/components/helper";
import TaskSyncSwapLP from "@/components/tasks/TaskSyncSwapLP.vue";
import MenuTaskSyncSwapLP from "@/components/tasks/menu/MenuTaskSyncSwapLP.vue";
import TaskMaverickSwap from "@/components/tasks/TaskMaverickSwap.vue";
import MenuTaskMaverickSwap from "@/components/tasks/menu/MenuTaskMaverickSwap.vue";
import TaskSpaceFiSwap from "@/components/tasks/TaskSpaceFiSwap.vue";
import MenuTaskSpaceFiSwap from "@/components/tasks/menu/MenuTaskSpaceFiSwap.vue";
import TaskVelocoreSwap from "@/components/tasks/TaskVelocoreSwap.vue";
import MenuTaskVelocoreSwap from "@/components/tasks/menu/MenuTaskVelocoreSwap.vue";
import TaskIzumiSwap from "@/components/tasks/TaskIzumiSwap.vue";
import MenuTaskIzumiSwap from "@/components/tasks/menu/MenuTaskIzumiSwap.vue";
import TaskVeSyncSwap from "@/components/tasks/TaskVeSyncSwap.vue";
import MenuTaskVeSyncSwap from "@/components/tasks/menu/MenuTaskVeSyncSwap.vue";
import TaskEzkaliburSwap from "@/components/tasks/TaskEzkaliburSwap.vue";
import MenuTaskEzkaliburSwap from "@/components/tasks/menu/MenuTaskEzkaliburSwap.vue";
import TaskZkSwap from "@/components/tasks/TaskZkSwap.vue";
import MenuTaskZkSwap from "@/components/tasks/menu/MenuTaskZkSwap.vue";


export interface TaskArg {
  task?: Task
  component: any
  weight: number

  taskType: TaskType
}

interface TaskSpec {
  menu: any
  component: any

  canBeEstimated: boolean
  deprecated: boolean

  descFn: DescFn

  service?: Service
}

interface Service {
  link: string
  img: string

  name: string

  implemented: string
}

export const serviceList: Service[] = [
  {
    name: 'Binance',
    img: '/icons/binance.ico',
    link: 'https://www.binance.com/',
    implemented: 'Withdraw'
  },
  {
    name: 'Okex',
    img: '/icons/okx.png',
    link: 'https://www.okx.com/',
    implemented: 'Withdraw, Deposit'
  },
]

interface DescFn {
  (task: Task): string;
}


export const taskProps: Record<TaskType, TaskSpec> = {
  Delay: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuDelayTask,
    component: TaskDelay,
    descFn(task: Task): string {
      let p = task.delayTask
      return ` (${p?.random ? `${humanDuration(p?.minRandom)}:${humanDuration(p?.maxRandom)}` : humanDuration(p?.duration)})`
    }
  },
  StargateBridge: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuTaskStargateBridge,
    component: TaskStargateBridge,
    descFn(task) {
      let p = task.stargateBridgeTask
      return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
    },
    service: {
      name: 'Startgate.finance',
      img: '/icons/stg.svg',
      link: 'https://stargate.finance/',
      implemented: 'Bridge'
    }
  },
  WithdrawExchange: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuExchangeWithdraw,
    component: TaskExchangeWithdraw,
    descFn(task) {
      let p = task.withdrawExchangeTask
      return ` (${p?.network} am: ${p?.sendAllCoins ? 'all' : `[${p?.amountMin}:${p?.amountMax}]`} ${p?.token})`
    }
  },
  OkexDeposit: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuOkexDeposit,
    component: TaskOkexDeposit,
    descFn(task) {
      let p = task.okexDepositTask
      return ` (${p?.network} ${p?.token})`
    }
  },
  TestNetBridgeSwap: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuTaskTestNetBridge,
    component: TaskTestNetBridgeSwap,
    descFn(task) {
      return `[${task.testNetBridgeSwapTask?.minAmount}:${task.testNetBridgeSwapTask?.maxAmount}]`
    },
    service: {
      name: 'TestNet-bridge',
      img: '/icons/testnetbridge.svg',
      link: 'https://testnetbridge.com/',
      implemented: 'Bridge'
    }
  },
  SnapshotVote: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuSnapshotTask,
    component: TaskSnapshotVote,
    descFn(task) {
      let p = task.snapshotVoteTask
      return ` (space: ${p?.space})`
    }
  },
  SyncSwap: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuTaskSyncSwap,
    component: TaskSyncSwap,
    descFn(task) {
      let p = task.syncSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    service: {
      name: 'SyncSwap',
      img: '/icons/syncswap.svg',
      link: 'https://syncswap.xyz/',
      implemented: 'Swap'
    }
  },
  ZkSyncOfficialBridgeToEthereum: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuZkSyncOfficialBridgeToEth,
    component: TaskZkSyncOfficialBridgeToEth,
    descFn(task) {
      let p = task.zkSyncOfficialBridgeToEthereumTask
      return ` (from ${Network.ZKSYNCERA} to ${Network.Etherium} ETH)`
    },
    service: {
      name: 'zksync bridge',
      img: '/icons/era.svg',
      link: 'https://portal.zksync.io/bridge/',
      implemented: 'L1 <-> L0'
    }
  },
  Mock: {
    deprecated: true,
    canBeEstimated: false,
    menu: null,
    component: null,
    descFn(task) {
      return ''
    }
  },
  OkexBinance: {
    deprecated: true,
    canBeEstimated: false,
    menu: null,
    component: null,
    descFn(task) {
      return ''
    }
  },
  Swap1inch: {
    deprecated: true,
    canBeEstimated: false,
    menu: null,
    component: null,
    descFn(task) {
      return ''
    }
  },
  OrbiterBridge: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuTaskOrbiterBridge,
    component: TaskOrbiterBridge,
    descFn(task) {
      let p = task.orbiterBridgeTask
      return ` (from ${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
    },
    service: {
      name: 'Orbiter.finance',
      img: '/icons/orbiter.ico',
      link: 'https://www.orbiter.finance/',
      implemented: 'Bridge'
    }
  },
  ZkSyncOfficialBridgeFromEthereum: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuZkSyncOfficialBridgeFromEth,
    component: TaskZkSyncOfficialBridgeFromEth,
    descFn(task) {
      let p = task.zkSyncOfficialBridgeFromEthereumTask
      return ` (from ${Network.Etherium} to ${Network.ZKSYNCERA} ETH)`
    },
    service: {
      name: 'zksync bridge',
      img: '/icons/era.svg',
      link: 'https://portal.zksync.io/bridge/',
      implemented: 'L1 <-> L0'
    }
  },
  WETH: {
    deprecated: false,
    menu: MenuTaskWethSwap,
    component: TaskWethSwap,
    canBeEstimated: true,
    descFn(task) {
      return task.wETHTask?.wrap ? ' (wrap)' : ' (unwrap)'
    }
  },
  MuteioSwap: {
    deprecated: false,
    component: TaskMuteIoSwap,
    menu: MenuTaskMuteioSwap,
    canBeEstimated: true,
    descFn(task) {
      const p = task.muteioSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    service: {
      name: 'Mute.io',
      img: '/icons/muteio.png',
      link: 'https://app.mute.io/swap',
      implemented: 'Swap'
    }
  },
  SyncSwapLP: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      let p = task.syncSwapLPTask
      return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p?.a}/${p?.b})`
    },
    menu: MenuTaskSyncSwapLP,
    component: TaskSyncSwapLP,
  },
  MaverickSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.maverickSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskMaverickSwap,
    menu: MenuTaskMaverickSwap,
    service: {
      name: 'Maverick',
      img: '/icons/maverick.ico',
      link: 'https://app.mav.xyz/?chain=324',
      implemented: ''
    }
  },
  SpaceFISwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.spaceFiSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskSpaceFiSwap,
    menu: MenuTaskSpaceFiSwap,
    service: {
      name: 'SpaceFi',
      link: 'https://swap-zksync.spacefi.io/#/swap',
      img: '/icons/spacefi.png',
      implemented: '',
    }
  },
  VelocoreSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.velocoreSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskVelocoreSwap,
    menu: MenuTaskVelocoreSwap,
    service: {
      name: 'Velocore',
      link: 'https://app.velocore.xyz/swap',
      img: '/icons/velocore.png',
      implemented: '',
    }
  },
  IzumiSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.izumiSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskIzumiSwap,
    menu: MenuTaskIzumiSwap,
    service: {
      name: 'IZUMI',
      link: 'https://izumi.finance/trade/swap',
      img: '/icons/izumi.svg',
      implemented: '',
    }
  },
  VeSyncSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.veSyncSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskVeSyncSwap,
    menu: MenuTaskVeSyncSwap,
    service: {
      name: 'VeSync',
      link: 'https://app.vesync.finance/swap',
      img: '/icons/vesync.svg',
      implemented: '',
    }
  },
  EzkaliburSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.izumiSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskEzkaliburSwap,
    menu: MenuTaskEzkaliburSwap,
    service: {
      name: 'Ezkalibur',
      link: 'https://dapp.ezkalibur.com/',
      img: '/icons/ezkalibur.ico',
      implemented: '',
    }
  },
  ZkSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.zkSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskZkSwap,
    menu: MenuTaskZkSwap,
    service: {
      name: 'ZkSwap',
      link: 'https://zkswap.finance/swap',
      img: '/icons/zkswap.ico',
      implemented: '',
    }
  }

}

export const getFlow = (flow: flow_Flow): string[] => {
  const result: string[] = []
  if (!flow || !flow.tasks) {
    return result
  }

  flow.tasks.forEach((task) => {

    if (!task || !task.taskType) {
      return
    }

    const out = task.taskType + " " + taskProps[task.taskType].descFn(task)
    result.push(out)
  })
  return result
}

export const taskComponentMap = new Map<TaskType, any>()
export const menuTaskComponentMap = new Map<TaskType, any>()
export const estimatedTaskMap = new Map<TaskType, boolean>()
export const taskTypes: TaskType[] = []

for (let name of Object.getOwnPropertyNames(taskProps)) {
  const n = name as TaskType

  const v = taskProps[n]
  if (v) {
    if (v.deprecated) {
      continue
    }

    taskComponentMap.set(n, v.component)
    menuTaskComponentMap.set(n, v.menu)
    taskTypes.push(n)
    estimatedTaskMap.set(n, v.canBeEstimated)
  }
}


