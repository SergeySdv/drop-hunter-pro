import {Network, ProcessStatus} from "@/generated/process";
import dayjs from "dayjs";

import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import {TaskType, Token} from "@/generated/flow";


dayjs.extend(utc)
dayjs.extend(timezone)

export const GetStatusText = (s: string): string => {
  switch (s) {
    case ProcessStatus.StatusError:
      return "Error"
    case ProcessStatus.StatusRunning:
      return "Running"
    case ProcessStatus.StatusReady:
      return "Ready"
    case ProcessStatus.StatusDone:
      return "Done"
    case ProcessStatus.StatusStop:
      return "Stop"
    case ProcessStatus.StatusRetry:
      return "Retry"
    default:
      return "Unknown"
  }
}

export const GetStatusColor = (s: string): string => {
  switch (s) {
    case ProcessStatus.StatusError:
      return "red"
    case ProcessStatus.StatusRunning:
      return "blue"
    case ProcessStatus.StatusReady:
      return "grey"
    case ProcessStatus.StatusDone:
      return "green"
    case ProcessStatus.StatusRetry:
      return "blue"
    case ProcessStatus.StatusStop:
      return "black"
    default:
      return "black"
  }
}


export class Timer {

  private timer: any = 0
  resolve: any


  fn = () => {

  }

  constructor() {
  }


  cb(fn: () => any) {
    this.fn = fn
  }

  add(seconds: number) {
    if (this.timer) {
      clearTimeout(this.timer)
    }

    this.timer = setTimeout(() => {
      this.fn()
    }, seconds)
  }
}

export const Delay = (seconds: number, cb?: (() => void)) => new Promise((resolve, reject) => {
  const timer = setTimeout(() => {
    if (cb) {
      cb()
    }
    resolve({})
  }, seconds)
})


export const formatTime = (d: Date): string => {
  const tz = dayjs.tz.guess()
  return dayjs(d).utc().tz(tz, false).format('YYYY-MM-DD HH:mm:ss')
}

export const humanDuration = (s?: string): string => {
  let totalSeconds = Number(s)
  const hours = Math.floor(totalSeconds / 3600);
  totalSeconds %= 3600;
  const minutes = Math.floor(totalSeconds / 60);
  const seconds = totalSeconds % 60;

  let result = ''
  if (hours !== 0) {
    result += ` ${hours} hour `
  }
  if (minutes !== 0) {
    result += ` ${minutes} min `
  }
  if (seconds !== 0) {
    result += ` ${seconds} sec `
  }

  if ((hours === 0) && (minutes === 0) && (seconds === 0)) {
    result = '0 sec'
  }

  return result
}


export const castGasLimitMap = (a: Object): Map<TaskType, number> => {

  const result = new Map<TaskType, number>()

  if (!a) {
    return result
  }

  for (const name of Object.getOwnPropertyNames(a)) {
    //@ts-ignore
    result.set(name, a[name])
  }

  return result
}

export const percent = (v: any) => {
  const value = Number(v).valueOf()
  if (value === 0) {
    return 'required'
  }

  if (v > 100) {
    return 'value can not be more than 100'
  }

  if (v < 0) {
    return 'value can not be negative'
  }

  if (!Number.isInteger(value)) {
    return 'number must be integer'
  }
  return true
}

export const onlyInteger = (v: number) => {
  if (!v) {
    return 'required'
  }

  if (Number.isInteger(Number(v))) {
    return true
  }
  return 'number must be integer'
}

export const ratio = (v: number) => {

  if (Number.isNaN(v)) {
    return 'must be number'
  }

  if (!v) {
    return 'required'
  }

  if (v <= 0) {
    return 'must be > 0'
  }

  if (v > 1) {
    return 'must be < 1'
  }
  return true
}


export const shuffleArray = <T = any>(array: T[]) => {
  for (var i = array.length - 1; i > 0; i--) {
    var j = Math.floor(Math.random() * (i + 1));
    var temp = array[i];
    array[i] = array[j];
    array[j] = temp;
  }
}

export interface SwapPair {
  name: string
  from: Token
  to: Token
}

export const tokenSwapPair = (t1: Token, t2: Token): SwapPair => {
  return {
    name: `${t1} -> ${t2}`,
    from: t1,
    to: t2
  }
}

