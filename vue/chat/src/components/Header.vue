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
    <v-divider vertical></v-divider>
    <span class="text-bar">Web-chat</span>
  </v-app-bar>
</template>

<script>
export default {
  name: 'HeaderComponent'
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
  margin-left: 10px;
  color: #FFFFFFB3;
}
</style>
