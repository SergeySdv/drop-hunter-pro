export const errorRef = (text: string): string => {
  let base = "("
  switch (true) {
    case text.search('Withdrawal address is not allowlisted for verification exemption') !== -1:
      base = addReason(base, "address is not in whitelist on okex exchange for withdrawal")
    case text.search('transaction failed') !== -1:
      base = addReason(base, "try to increase gas multiplayer value in /settings")
    case text.search('is higher than max') !== -1:
      base = addReason(base, "try to increase max gas value in /settings")
  }

  base += ")"

  if (base == "()") {
    return ""
  }
  base = "possible reasons: " + base
  return base
}

const addReason = (base, reason: string): string => {
  return base + reason + ", "
}
