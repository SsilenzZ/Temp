<template>
  <v-container max-width="100%">
    <v-row>
      <v-col>
        <Toolbar1 :jwt="jwt" @addFriend="addFriend" @block="block" />

        <List :jwt="jwt" :blacklist="blacklist" :friends="friends" @removeFriend="removeFriend" @unblock="unblock"/>
      </v-col>

      <v-divider vertical class="vrow"></v-divider>

      <v-col>

        <Toolbar2 />

        <Chat />

      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { jwtVerify } from '@/utilities/jwt-verify'
import Toolbar1 from '@/components/Toolbar1'
import Toolbar2 from '@/components/Toolbar2'
import List from '@/components/List'
import axios from 'axios'
import Chat from '@/components/Chat'

export default {
  name: 'MainView',
  components: {
    Toolbar2,
    Toolbar1,
    List,
    Chat
  },
  data () {
    return {
      friends: [],
      blacklist: [],
      jwt: null
    }
  },
  created () {
    let jwt = jwtVerify(localStorage.getItem('jwt'))
    if (!jwt) {
      this.$router.push('/')
    } else {
      this.jwt = jwt
    }
  },

  async beforeMount () {
    const res1 = await axios({
      method: 'get',
      url: process.env.VUE_APP_API_URL + '/jwt/friends',
      headers: {
        Authorization: 'Bearer ' + this.jwt
      }
    }).catch(function (error) {
      if (error.response) {
        alert('Error while loading friends')
      }
    })
    this.friends = res1.data
    const res2 = await axios({
      method: 'get',
      url: process.env.VUE_APP_API_URL + '/jwt/blacklist',
      headers: {
        Authorization: 'Bearer ' + this.jwt
      }
    }).catch(function (error) {
      if (error.response) {
        alert('Error while loading blacklist')
      }
    })
    this.blacklist = res2.data
  },

  methods: {
    jwtVerify,
    addFriend (v) {
      this.friends.push(v)
    },
    block (v) {
      this.removeFriend(v)
      this.blacklist.push(v)
    },
    removeFriend (v) {
      for (var i in this.friends) {
        if (this.friends[i].id === v.id) {
          this.friends.splice(i, 1)
          break
        }
      }
    },
    unblock (v) {
      for (var i in this.blacklist) {
        if (this.blacklist[i].id === v.id) {
          this.blacklist.splice(i, 1)
          break
        }
      }
    }
  }
}
</script>

<style scoped>
.container {
  padding: 0 0 !important;
  min-width: 1200px !important;
  max-width: 100% !important;
}
.col {
  padding: 12px 0 !important;
  height: 98% !important;
}
.row {
  margin: 0 !important;
}
.v-divider--vertical {
  margin: 0 !important;
}
.vrow {
  margin-top: 12px !important;
}
</style>
