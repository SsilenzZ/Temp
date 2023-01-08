<template>
  <v-app>
    <v-main>
      <Header />
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

import Header from '@/components/Header'

export default {
  name: 'App',
  components: {
    Header
  },
  mounted () {
    document.title = 'Web-chats'
  }
}

</script>

<style>
::-webkit-scrollbar {display:none;}
.v-main {
  background-color: #1E1E1E;
}
</style>
