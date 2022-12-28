<template>
  <v-app-bar
    color="grey-lighten-4"
    dense
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
  max-width: 24px;
  margin-right: 15px;
}
.text-bar {
  font-size: large;
  padding: 5px;
  margin-left: 10px;
  color: #FFFFFFB3;
}
</style>
