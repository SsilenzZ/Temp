<template>
  <v-col>
    <v-row
      style="margin-top: 10px;
      margin-right: 0;
      overflow: auto;
      height: 732px;"
    >
      <v-col
        style="padding: 0;
        margin: 0 12px;"
      >
        <v-list
          v-for="message in chat" ref="chat"
          :key="message.id"
        >
          <v-list-item v-if="message.senderID === chatWith.id">
            <v-list-item-content class="received-message justify-start">
              <v-card color=#5110e7 class="flex-none message">
                <v-card-text class="white--text pa-2 d-flex flex-column">
                  <span class="align-self-start text-subtitle-1">
                    {{ message.text }}
                  </span>

                  <span class="text-caption font-italic align-self-end">
                    {{ message.createdAt }}
                  </span>
                </v-card-text>
              </v-card>
            </v-list-item-content>
          </v-list-item>
          <v-list-item v-else>
            <v-list-item-content class="sent-message justify-end">
              <v-card color=#5110e7 class="flex-none message">
                <v-card-text class="white--text pa-2 d-flex flex-column">
                  <span class="text-subtitle-1 chat-message">
                    {{ message.text }}
                  </span>

                  <span class="text-caption font-italic align-self-start">
                    {{ message.createdAt }}
                  </span>
                </v-card-text>
              </v-card>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>

    <v-row class="send" align="bottom">
      <v-textarea
        v-model="text"
        v-on:keydown.enter.prevent="send"
        no-resize
        solo
        rows="1"
        background-color=#272727
        color=#5110e7
        elevation="4"
        clear-icon="mdi-close-circle"
        clearable
        label="Message"
        type="text"
      ></v-textarea>
      &nbsp;
      <v-btn
        :value="send"
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
      text: ''
    }
  },

  props: ['chat', 'chatWith'],

  watch: {
    chat: {
      deep: true,
      handler () {
        this.$nextTick(() => {
          this.scrollDown()
        })
      }
    }
  },

  methods: {
    send (text) {
      this.$emit('send', text)
      this.text = ''
    }
  },

  scrollDown () {
    var messages = this.$refs.chat
    messages.scrollTop = messages.scrollHeight - messages.clientHeight
  }
}
</script>

<style scoped>
.send {
  padding: 5px 26px;
  border-color: #272727;
}
.fill-height {
  max-height: 100% !important;
}
.chat-message {
  display: unset !important;
  white-space: break-spaces;
}
.sent-message .received-message {
  width: fit-content;
  height: fit-content;
}
.message {
  max-width: 45%;
  max-height: fit-content;
}
.flex-none {
  flex: unset;
}
.received-message::after {
  content: ' ';
  position: absolute;
  width: 0;
  height: 0;
  left: 5px;
  right: auto;
  top: 12px;
  bottom: auto;
  border: 12px solid;
  border-color: #5110e7 transparent transparent transparent;
}
.sent-message::after {
  content: ' ';
  position: absolute;
  width: 0;
  height: 0;
  left: auto;
  right: 5px;
  top: 12px;
  bottom: auto;
  border: 12px solid;
  border-color: #5110e7 transparent transparent transparent;
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
</style>
