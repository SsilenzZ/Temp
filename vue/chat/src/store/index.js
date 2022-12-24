import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    socket: null
  },
  getters: {
  },
  mutations: {
    updateConnection (state, socket) {
      state.socket = socket
    }
  },
  actions: {
    connect (ctx, jwt) {
      if (ctx.getters.socket == null) {
        const socket = new WebSocket(process.env.VUE_APP_WEBSOCKET + '/' + jwt)
        ctx.commit('updateConnection', socket)
      }
    },
    disconnect (ctx) {
      ctx.commit('updateConnection', null)
    }
  },
  modules: {
    socket (state) {
      return state.socket
    }
  }
})
