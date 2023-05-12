import {Network, ProcessStatus} from "@/generated/process";

interface Task {

}


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
    case ProcessStatus.StatusStop:
      return "orang"
    default:
      return "black"
  }
}

export const Delay = (seconds: number) => new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve({})
  }, seconds)
})

