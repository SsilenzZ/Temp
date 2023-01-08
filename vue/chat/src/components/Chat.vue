<template>
  <v-col>
    <v-row align="end" class="fill-height">
      <v-col>
        <div v-for="(item, index) in chat" :key="index"
             :class="['d-flex flex-row align-center my-2', item.from === 'user' ? 'justify-end': null]">
          <span v-if="item.from === 'user'">{{ item.msg }}</span>
          <v-avatar :color="item.from === 'user' ? 'indigo': 'red'" size="36">
            <span class="white--text">{{ item.from[0] }}</span>
          </v-avatar>
          <span v-if="item.from !== 'user'" class="blue--text ml-3">{{ item.msg }}</span>
        </div>
      </v-col>
    </v-row>

    <v-row class="send" align="bottom">
      <v-text-field
        v-model="message"
        background-color=#272727
        color=#5110e7
        elevation="4"
        solo
        clear-icon="mdi-close-circle"
        clearable
        label="Message"
        type="text"
        @keypress.enter="send"
      ></v-text-field>
      &nbsp;
      <v-btn
        color=#5110e7
        elevation="4"
        fab
        icon
        @click="send"
      >
        <v-icon>mdi-send</v-icon>
      </v-btn>
    </v-row>
  </v-col>
</template>

<script>
export default {
  name: 'ChatComp',

  data: () => {
    return {
      message: null,
      chat: []
    }
  },

  methods: {
    send () {
      this.chat.push(
        {
          from: 'user',
          msg: this.message
        })
      this.message = null
      this.addReply()
    },

    addReply () {
      this.chat.push({
        from: 'sushant',
        msg: 'Hmm'
      })
    }
  }
}
</script>

<style scoped>
.send {
  padding: 0 12px;
  border-color: #272727;
}
.v-btn {
  background: #272727 !important;
}
.v-btn:hover {
  color: #5110e7 !important;
}
.v-icon {
  color: #FFFFFFB3 !important;
  margin-left: 3px !important;
}
.v-text-field {
  margin-top: 4px !important;
}
.row {
  margin: 0 !important;
  overflow-y:auto !important;
}
</style>
