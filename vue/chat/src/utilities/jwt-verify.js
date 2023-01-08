export function jwtVerify (jwt) {
  if (jwt !== undefined) {
    const item = JSON.parse(jwt)
    const now = new Date()
    if (now.getTime() > item.expire) {
      localStorage.removeItem('jwt')
      alert('The session has been ended up')
      return false
    }
    return item.value.replace(/"/g, '')
  }
  return false
}
