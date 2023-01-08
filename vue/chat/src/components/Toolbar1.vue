<template>
  <v-card
    color="grey lighten-4"
    tile
  >
    <div>
      <v-toolbar class="tb1">
        <v-app-bar-nav-icon></v-app-bar-nav-icon>
        <v-autocomplete
          clearable
          hide-details
          v-model="model"
          :loading="loading"
          :items="users"
          :search-input.sync="search"
          item-text="name"
          item-value="id"
          prepend-icon="mdi-magnify"
          class="search"
          color=#5110e7
          label="Enter user`s name to search"
        >
          <template v-slot:no-data>
            <v-list-item>
              <v-list-item-title>
                No users found
              </v-list-item-title>
            </v-list-item>
          </template>
          <template v-slot:item="{ item }">
            <v-list-item-avatar
              color=#5110e7
              class="text-h5 font-weight-light white--text"
            >
              {{ item.name.charAt(0) }}
            </v-list-item-avatar>
              <v-list-item-content>
                  <v-row>
                  <v-col>
                    <v-list-item-title v-text="item.name"></v-list-item-title>
                    <v-list-item-subtitle v-text="'ID:' + item.id"></v-list-item-subtitle>
                  </v-col>

                  <v-col>
                    <v-row>
                      <v-btn
                        color=#5110e7
                        class="friend"
                        @click="friendUser(item.id)"
                      >
                        <v-icon>mdi-account-plus</v-icon>
                        &nbsp;Add
                      </v-btn>
                      &nbsp;
                      <v-btn
                        class="block"
                        @click="blockUser(item.id)"
                      >
                        <v-icon
                          color=#5110e7
                        >mdi-lock-outline
                        </v-icon>
                        &nbsp;Block
                      </v-btn>
                    </v-row>
                  </v-col>
                </v-row>
              </v-list-item-content>
          </template>
        </v-autocomplete>
      </v-toolbar>
    </div>
  </v-card>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Toolbar1Comp',
  data () {
    return {
      model: null,
      menu: false,
      loading: false,
      search: null,
      users: []
    }
  },

  watch: {
    model (val) {
      if (val != null) this.tab = 0
      else this.tab = null
    },
    search (val) {
      val && val !== this.select && this.findUser(val)
    }
  },

  // eslint-disable-next-line vue/no-dupe-keys
  props: ['jwt'],

  methods: {
    addFriend (v) {
      this.$emit('addFriend', v)
    },
    block (v) {
      this.$emit('block', v)
    },
    async findUser (login) {
      this.loading = true
      await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/find',
        headers: {
          Authorization: 'Bearer ' + this.jwt
        },
        data: {
          name: login
        }
      }).then(res => {
        this.users = res.data
      }).catch(err => {
        console.log(err)
      })
        .finally(() => (this.loading = false))
    },

    async friendUser (id) {
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/friend',
        headers: {
          Authorization: 'Bearer ' + this.jwt
        },
        data: {
          id2: id
        }
      })
        .catch(function (error) {
          if (error.response) {
            alert('This user already your friend/blocked')
          }
        })

      if (response !== undefined) {
        this.addFriend(response.data)
      }
    },

    async blockUser (id) {
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/block',
        headers: {
          Authorization: 'Bearer ' + this.jwt
        },
        data: {
          id2: id
        }
      })
        .catch(function (error) {
          if (error.response) {
            alert('This user already blocked')
          }
        })
      if (response !== undefined) {
        this.block(response.data)
      }
    }
  }
}
</script>

<style scoped>
.tb1, .v-toolbar__content {
  height: 64px !important;
}
.friend, .block {
  flex: 1 0 5%;
  padding: 0 16px;
  margin-top: 5%;
}
.block {
  margin-right: 8%;
}
.row {
  flex: 1 0 4% !important;
}
.v-list-item__avatar {
  padding-left: 25px !important;
  width: 60px !important;
  height: 60px !important;
  font-size: 2rem !important;
}
.v-input__slot {
  color: #FFFFFFB3 !important;
}
.row {
  margin: 0 !important;
}
</style>
