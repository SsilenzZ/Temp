export function jwtVerify (jwt) {
  if (jwt !== undefined) {
    const item = JSON.parse(jwt)
    const now = new Date()
    if (now.getTime() > item.expire) {
      localStorage.removeItem('jwt')
      alert('Token of session is expired')
      return false
    }
    return item.value
  }
  return false
}
