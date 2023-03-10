<template>
  <v-card tile>
    <v-tabs
      v-model="tab"
      background-color=#1E1E1E
      color=#5110e7
      centered
      dark
      icons-and-text
    >
      <v-tabs-slider></v-tabs-slider>
      <v-tab href="#tab-1">
        Chats
        <v-icon color=#5110e7>mdi-forum-outline</v-icon>
      </v-tab>

      <v-tab href="#tab-2">
        Friends
        <v-icon color=#5110e7>mdi-account-box-outline</v-icon>
      </v-tab>

      <v-tab href="#tab-3">
        Blacklist
        <v-icon color=#5110e7>mdi-account-lock-outline</v-icon>
      </v-tab>
    </v-tabs>

    <v-divider></v-divider>

    <v-tabs-items v-model="tab">
      <v-tab-item :value="'tab-1'">
        <v-list two-line>
          <template v-for="chatsPreview in chatsPreview">
            <v-list-item
              :key="chatsPreview.id"
              :class="chatsPreview.id === chatID ? 'active-chat' : null"
              @click="openChatBox(chatsPreview.id); getChatMessages(chatsPreview.senderID)"
            >
              <v-list-item-avatar
                color=#5110e7
                class="text-h5 font-weight-light white--text"
              >
                {{ chatsPreview.senderLogin.charAt(0) }}
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title v-text="chatsPreview.senderLogin"></v-list-item-title>
                <v-list-item-subtitle v-text="chatsPreview.text"></v-list-item-subtitle>
                <v-list-item-subtitle v-text="chatsPreview.time"></v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>

            <v-divider
              :key="chatsPreview.chatID"
              dark
            >
            </v-divider>

          </template>
        </v-list>
      </v-tab-item>

      <v-tab-item :value="'tab-2'">
        <v-list>
          <template v-for="(item, id) in friends">
            <v-list-item
              :key="item.title"
            >
              <v-list-item-avatar
                color=#5110e7
                class="text-h5 font-weight-light white--text"
              >
                {{ item.name.charAt(0) }}
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title v-text="item.name"></v-list-item-title>
                <v-list-item-subtitle v-text="'ID:' + item.id"></v-list-item-subtitle>
                <v-btn
                  class="friend"
                  @click="unfriendUser(item.id)"
                >
                  <v-icon
                    color=#5110e7
                  >mdi-account-minus-outline
                  </v-icon>
                  &nbsp;Remove friend
                </v-btn>
              </v-list-item-content>
            </v-list-item>

            <v-divider
              :key="id"
              dark
            >
            </v-divider>

          </template>
        </v-list>
      </v-tab-item>

      <v-tab-item
        :value="'tab-3'"
      >
        <v-list>
          <template v-for="(item, id) in blacklist">
            <v-list-item
              :key="item.title"
            >
              <v-list-item-avatar
                color=#5110e7
                class="text-h5 font-weight-light white--text"
              >
                {{ item.name.charAt(0) }}
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title v-text="item.name"></v-list-item-title>
                <v-list-item-subtitle v-text="'ID:' + item.id"></v-list-item-subtitle>
                <v-btn
                  class="block"
                  @click="unblockUser(item.id)"
                >
                  <v-icon
                    color=#5110e7
                  >mdi-lock-open-variant-outline
                  </v-icon>
                  &nbsp;Unblock User
                </v-btn>
              </v-list-item-content>
            </v-list-item>

            <v-divider
              :key="id"
              dark
            >
            </v-divider>

          </template>
        </v-list>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ListComp',

  data () {
    return {
      open: false,
      tab: null
    }
  },

  props: ['jwt', 'blacklist', 'friends', 'chatsPreview', 'chatID'],

  methods: {
    removeFriend (v) {
      this.$emit('removeFriend', v)
    },

    unblock (v) {
      this.$emit('unblock', v)
    },

    openChatBox (v) {
      this.$emit('openChatBox', v)
    },

    getChatMessages (v) {
      this.$emit('getChatMessages', v)
    },

    async unfriendUser (id) {
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/unfriend',
        headers: {
          Authorization: 'Bearer ' + this.jwt
        },
        data: {
          id2: id
        }
      })
        .catch(function (error) {
          if (error.response) {
            alert('This user is not your friend')
          }
        })

      if (response !== undefined) {
        alert('Success')
        this.removeFriend(response.data)
      }
    },

    async unblockUser (id) {
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/unblock',
        headers: {
          Authorization: 'Bearer ' + this.jwt
        },
        data: {
          id2: id
        }
      })
        .catch(function (error) {
          if (error.response) {
            alert('This user is not in your blacklist')
          }
        })

      if (response !== undefined) {
        alert('Success')
        this.unblock(response.data)
      }
    }
  }
}
</script>

<style scoped>
.friend, .block {
  flex: 1 0 49%;
  padding: 0 16px;
}
.active-chat {
  background-color: #272727;
}
.v-list-item__avatar {
  padding-left: 20px !important;
  width: 60px !important;
  height: 60px !important;
  font-size: 2rem !important;
}
.v-list-item__title, .v-list-item__subtitle {
  flex: 1 1 51% !important;
}
</style>
