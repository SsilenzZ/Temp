<template>
  <v-app-bar
    color=#1E1E1E
    elevation="7"
    rounded
  ><v-img
    :src="require('../assets/logo-32x32.png')"
    contain
    height="32"
    class="logo-icon"
  />
    <b class="text-bar">Web-chat</b>
  </v-app-bar>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Header'
}
</script>

<style scoped>

.logo-icon {
  max-width: 32px;
  margin-right: 15px;
}
.text-bar {
  font-size: x-large;
  padding: 5px;
  color: #FFFFFFB3;
}
</style>
