import {Amount} from "@/generated/flow";

export const getStatusColor = (status: string) => {
  switch (status) {
    case 'error':
      return 'red'
    case 'completed':
      return 'green'
    case 'not started':
      return 'grey'
    default:
      return 'blue'
  }
}

export const getAmountSend = (am?: Amount): string => {
  if (!am) {
    return 'all coins'
  }

  switch (true) {
    case am.sendAll:
      return 'all coins'
    case am.sendPercent !== undefined:
      return `${String(am.sendPercent)}%`
    case am.sendValue !== undefined:
      return `${String(am.sendAmount)} `
    default:
      return "invalid value"
  }
}

export const onlyInteger = (v: any) => {
  if (!v) {
    return 'required'
  }

  if (Number.isInteger(Number(v))) {
    return true
  }
  return 'number must be integer'
}

export const required = (v: any) => !!v || 'required'
