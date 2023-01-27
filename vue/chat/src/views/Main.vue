<template>
  <v-container fluid class="fill-height">
    <v-row>
      <v-col width="49%">
        <Toolbar1 :jwt="jwt" @addFriend="addFriend" @block="block" />

        <List
          :jwt="jwt"
          :chatID="currentChatID"
          :blacklist="blacklist"
          :friends="friends"
          :chatsPreview="chatsPreview"
          @removeFriend="removeFriend"
          @unblock="unblock"
          @openChatBox="openChatBox"
          @getChatMessages="getChatMessages"
        />
      </v-col>

      <v-divider vertical class="vrow fill-height"></v-divider>

      <v-col width="49%">
        <Toolbar2
          v-if="open"
          :name="chatWith.login"
          @closeChatBox="closeChatBox"
        />

        <Chat
          v-if="open"
          :chat="chat"
          :chatWith="chatWith"
        />

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
    Toolbar1,
    Toolbar2,
    List,
    Chat
  },
  data () {
    return {
      friends: [],
      blacklist: [],
      currentChatID: 0,
      currentSender: '',
      chatsPreview: [],
      chat: [],
      chatWith: {},
      message: {},
      pingMsg: JSON.stringify({
        senderID: 0,
        receiverID: 0,
        text: 'ping'
      }),
      open: false,
      jwt: null,
      socket: null
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

    const res3 = await axios({
      method: 'get',
      url: process.env.VUE_APP_API_URL + '/jwt/chatspreview',
      headers: {
        Authorization: 'Bearer ' + this.jwt
      }
    }).catch(function (error) {
      if (error.response) {
        alert('Error while loading chats preview')
      }
    })
    for (let i = 0; i < res3.data.length; i++) {
      let rawTime = new Date(res3.data[i].time)
      let time = rawTime.toLocaleString()
      if (time === 'Invalid Date') {
        time = ''
      }
      res3.data[i].time = time
    }
    this.chatsPreview = res3.data

    this.$store.dispatch('connect', this.jwt)
    this.socket = this.$store.getters.socket
    this.socket.onopen = () => {
      const msg = this.pingMsg
      this.socket.send(msg)
    }
    this.socket.onmessage = (event) => {
      this.readWsMessage(JSON.parse(event.data))
    }
    this.socket.onclose = () => {
      alert('Reload this page to make it active again')
    }
    this.socket.onerror = () => {
      alert('error')
      this.$router.go()
    }
    if (this.$route.query.id !== undefined) {
      this.getChatMessages(this.$route.query.id)
      alert('OK')
    }
    await this.$router.replace('/chat')
  },

  methods: {
    jwtVerify,
    async getChatMessages (id) {
      const response = await axios({
        method: 'get',
        url: process.env.VUE_APP_API_URL + '/jwt/chat/' + id,
        headers: {
          Authorization: 'Bearer ' + this.jwt
        }
      }).catch(function (error) {
        if (error.response) {
          alert('Error while getting messages')
        }
      })
      for (let i = 0; i < response.data.chat.length; i++) {
        let rawTime = new Date(response.data.chat[i].CreatedAt)
        response.data.chat[i].CreatedAt = rawTime.toLocaleString()
      }
      this.chat = response.data.chat
      this.chatWith = response.data.user
      this.currentChatID = response.data.id
    },

    send (text) {
      const obj = {
        SenderID: 0,
        ReceiverID: this.chatWith.id,
        Text: text
      }
      this.socket.send(JSON.stringify(obj))
      console.log(this.socket)
    },

    async ping () {
      const socket = this.socket
      const pingMsg = this.pingMsg
      setTimeout(function () {
        socket.send(pingMsg)
      }, 30000)
    },

    readWsMessage (data) {
      if (data.text === 'pong' || data.code === 1) {
        this.ping()
      } else if (data.code === 0) {
        alert('You can`t send a message to blocked user')
      } else {
        let id = this.chatsPreview.findIndex((x) => x.chatID === data.chatID)
        let rawTime = new Date(data.time)
        let time = rawTime.toLocaleString()
        if (id === -1) {
          this.chatsPreview.push({
            chatID: data.id,
            messageID: data.messageID,
            senderID: data.senderID,
            text: data.text,
            time: time
          })
        } else {
          this.chatsPreview[id].messageID = data.messageID
          this.chatsPreview[id].text = data.text
          this.chatsPreview[id].time = time
        }
        if (data.chatID === this.currentChatID) {
          this.chat.push({
            id: data.id,
            chatID: data.chatID,
            senderID: data.userID,
            text: data.text,
            createdAt: time
          })
        }
      }
    },

    addFriend (v) {
      this.friends.push(v)
    },

    block (v) {
      this.removeFriend(v)
      this.blacklist.push(v)
    },

    removeFriend (v) {
      for (let i in this.friends) {
        if (this.friends[i].id === v.id) {
          this.friends.splice(i, 1)
          break
        }
      }
    },

    unblock (v) {
      for (let i in this.blacklist) {
        if (this.blacklist[i].id === v.id) {
          this.blacklist.splice(i, 1)
          break
        }
      }
    },

    openChatBox (id) {
      this.currentChatID = id
      this.chatWith = {
        id: this.chatsPreview[id - 1].senderID,
        login: this.chatsPreview[id - 1].senderLogin
      }
      if (this.open === false) {
        this.open = true
      }
    },

    closeChatBox () {
      this.currentChatID = 0
      this.chatWith = {}
      this.open = false
    }
  }
}
</script>

<style scoped>
.container {
  padding: 0 0 !important;
  min-width: 1200px !important;
  max-width: 100% !important;
  align-items: inherit !important;
}
.col {
  padding: 9px 0 !important;
}
.row {
  margin: 0 !important;
  max-height: 98% !important;
}
.v-divider--vertical {
  margin: 0 !important;
}
.vrow {
  margin-top: 74px !important;
}
.fill-height {
  height: 88.8% !important;
}
</style>
