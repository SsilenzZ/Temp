<template @>
  <v-container max-width="100%">
    <template>
      <v-card
        color="grey lighten-4"
        flat
        height="200px"
        tile
      >
        <v-toolbar>
          <v-app-bar-nav-icon></v-app-bar-nav-icon>
          <v-autocomplete
            auto-select-first
            clearable
            hide-details
            hide-selected
            v-model="model"
            :loading="loading"
            :items="items"
            :search-input.sync="search"
            item-text="name"
            item-value="id"
            prepend-icon="mdi-magnify"
            class="search"
            color=#6200EE
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
                color=#6200EE
                class="text-h5 font-weight-light white--text"
              >
                {{ item.name.charAt(0) }}
              </v-list-item-avatar>
              <v-list-item-content>
                <v-list-item-title v-text="item.name"></v-list-item-title>
                <v-list-item-subtitle v-text="'ID:' + item.id"></v-list-item-subtitle>
              </v-list-item-content>
            </template>
          </v-autocomplete>
          <v-divider vertical></v-divider>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </v-toolbar>
      </v-card>
    </template>
  </v-container>
</template>

<script>
import axios from 'axios'
import { jwtVerify } from '@/utilities/jwt-verify'

export default {
  name: 'HomeView',
  data () {
    return {
      loading: false,
      items: [],
      search: null,
      select: null,
      model: null,
      tab: null,
      login: ''
    }
  },
  created () {
    if (jwtVerify(localStorage.getItem('jwt'))) {
      this.$router.push('/home')
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
  methods: {
    jwtVerify,
    async findUser (login) {
      this.loading = true
      let token = localStorage.getItem('jwt')
      await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/jwt/find',
        headers: {
          Authorization: 'Bearer' + ' ' + token.replace(/"/g, '')
        },
        data: {
          name: login
        }
      }).then(res => {
        this.items = res.data
      }).catch(err => {
        console.log(err)
      })
        .finally(() => (this.loading = false))
    }
  }
}
</script>

<style scoped>
.search {
  max-width: 45%;
  padding-right: 2%;
}
.v-input__slot {
  color: #FFFFFFB3 !important;
}
.v-icon {
  color: #FFFFFFB3 !important;
}
.container {
  min-width: 960px !important;
  max-width: 100% !important;
  padding: 8px;
}
.v-list-item__avatar {
  padding-left: 13px !important;
}
</style>
