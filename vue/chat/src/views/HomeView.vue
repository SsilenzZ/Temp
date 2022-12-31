<template>
  <v-container max-width="100%">
    <v-row>
      <v-col>
        <toolbar1 />
        <v-card tile>
          <v-tabs
            v-model="tab"
            background-color=#1E1E1E
            color=#6200EE
            centered
            dark
            icons-and-text
          >
            <v-tabs-slider></v-tabs-slider>
            <v-tab href="#tab-1">
              Chats
              <v-icon color=#6200EE>mdi-forum-outline</v-icon>
            </v-tab>

            <v-tab href="#tab-2">
              Friends
              <v-icon color=#6200EE>mdi-account-box-outline</v-icon>
            </v-tab>

            <v-tab href="#tab-3">
              Blacklist
              <v-icon color=#6200EE>mdi-account-lock-outline</v-icon>
            </v-tab>
            </v-tabs>

            <v-tabs-items v-model="tab">
              <v-tab-item
                :value="'tab-2'"
              >
                <v-list>
                  <template v-for="(item, id) in friends">
                    <v-subheader
                      v-if="item.header"
                      :key="item.header"
                      v-text="item.header"
                    ></v-subheader>

                    <v-divider
                      v-else-if="item.divider"
                      :key="id"
                      :inset="item.inset"
                    ></v-divider>

                    <v-list-item
                      v-else
                      :key="item.title"
                    >
                      <v-list-item-avatar
                        color=#6200EE
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
                            color=#6200EE
                            >mdi-trash-can-outline
                          </v-icon>
                          &nbsp;Remove friend
                        </v-btn>
                      </v-list-item-content>
                    </v-list-item>

                    <v-divider
                      :key="id"
                      inset
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
                    <v-subheader
                      v-if="item.header"
                      :key="item.header"
                      v-text="item.header"
                    ></v-subheader>

                    <v-divider
                      v-else-if="item.divider"
                      :key="id"
                      :inset="item.inset"
                    ></v-divider>

                    <v-list-item
                      v-else
                      :key="item.title"
                    >
                      <v-list-item-avatar
                        color=#6200EE
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
                            color=#6200EE
                          >mdi-trash-can-outline
                          </v-icon>
                          &nbsp;Unblock User
                        </v-btn>
                      </v-list-item-content>
                    </v-list-item>

                    <v-divider
                      :key="id"
                      inset
                      dark
                    >
                    </v-divider>

                  </template>
                </v-list>
              </v-tab-item>
            </v-tabs-items>
          </v-card>
        </v-col>
        <v-col>
          <toolbar2 />
          <v-col cols="12">
          <v-img
            :src="require('../assets/logo.png')"
            contain
            height="200"
          />
          </v-col>
        </v-col>
      </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'
import { jwtVerify } from '@/utilities/jwt-verify'
import Toolbar1 from '@/components/Toolbar1'
import Toolbar2 from '@/components/Toolbar2'

export default {
  name: 'HomeView',
  components: {
    Toolbar2,
    Toolbar1
  },
  data () {
    return {
      friends: [],
      blacklist: [],
      select: null,
      tab: null,
      login: ''
    }
  },
  created () {
    if (jwtVerify(localStorage.getItem('jwt'))) {
      this.$router.push('/')
    }
  },

  async beforeMount () {
    const token = localStorage.getItem('jwt')
    const res1 = await axios({
      method: 'get',
      url: process.env.VUE_APP_API_URL + '/jwt/friends',
      headers: {
        Authorization: 'Bearer ' + token.replace(/"/g, '')
      }
    }).catch(function (error) {
      if (error.response) {
        alert('error')
      }
    })
    this.friends = res1.data
    const res2 = await axios({
      method: 'get',
      url: process.env.VUE_APP_API_URL + '/jwt/blacklist',
      headers: {
        Authorization: 'Bearer ' + token.replace(/"/g, '')
      }
    }).catch(function (error) {
      if (error.response) {
        alert('error')
      }
    })
    this.blacklist = res2.data
  },
  methods: {
    jwtVerify,
    async unfriendUser (id) {
      const token = localStorage.getItem('jwt')
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/unfriend',
        headers: {
          Authorization: 'Bearer ' + token.replace(/"/g, '')
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
        window.location.reload()
      }
    },

    async unblockUser (id) {
      const token = localStorage.getItem('jwt')
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/unblock',
        headers: {
          Authorization: 'Bearer ' + token.replace(/"/g, '')
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
        window.location.reload()
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
.v-container {
  padding-top: 15px !important;
}
.col {
  padding: 12px 0 !important;
}
.v-input__slot {
  color: #FFFFFFB3 !important;
}
.container {
  min-width: 960px !important;
  max-width: 100% !important;
  padding: 8px;
}
.v-list-item__avatar {
  padding-left: 25px !important;
  width: 60px !important;
  height: 60px !important;
  font-size: 2rem !important;
}
.v-list-item__title, .v-list-item__subtitle {
  flex: 1 1 51% !important;
}
</style>
