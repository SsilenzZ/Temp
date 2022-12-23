<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.png')"
          contain
          height="200"
        />
      </v-col>

      <v-col class="dialog-center">
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to Web-chat
        </h1><br><br><br><br>
        <v-dialog
          transition="dialog-top-transition"
          max-width="400"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              bottom
              color=#6200EE
              v-bind="attrs"
              v-on="on"
              class="start"
            >
              <b>Get started&nbsp;</b>
              <v-icon>mdi-arrow-right-circle</v-icon>
            </v-btn>
          </template>
          <template>
            <v-card
              height=500px>
              <v-banner
                single-line
                color=#1E1E1E
              >
                <h2>Registration</h2>
              </v-banner>
              <v-icon
                color=#6200EE
                size=150px
                class="icon"
              >mdi-account-box</v-icon>
              <template>
                <v-form v-model="isValid">
                  <v-card-text>&nbsp;&nbsp;&nbsp;Email:</v-card-text>
                  <v-text-field :rules="[rules.required, rules.max, rules.email]"
                                outlined
                                v-model="login"
                                color=#6200EE
                                label="Enter your email here"
                                class="text-field"
                  ></v-text-field>
                  <v-card-text>&nbsp;&nbsp;&nbsp;Password:</v-card-text>
                  <v-text-field :rules="[rules.required, rules.min, rules.password]"
                                outlined
                                v-model="password"
                                :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
                                :type="show1 ? 'text' : 'password'"
                                color=#6200EE
                                label="Enter your password here"
                                class="text-field"
                                @click:append="show1 = !show1"
                  ></v-text-field>
                  <v-card-text>&nbsp;&nbsp;&nbsp;Repeat password:</v-card-text>
                  <v-text-field :rules="rules.repeat"
                                outlined
                                v-model="confirm"
                                :append-icon="show2 ? 'mdi-eye' : 'mdi-eye-off'"
                                :type="show2 ? 'text' : 'password'"
                                color=#6200EE
                                label="Repeat your password here"
                                class="text-field"
                                @click:append="show2 = !show2"
                  ></v-text-field>
                  <v-btn
                    :disabled="!isValid"
                    height=60px;
                    width=120px;
                    color=#6200EE
                    class="signup"
                    @click = "SignUp"
                  >
                    <b>Sign Up&nbsp;</b>
                    <v-icon>mdi-check-circle</v-icon>
                  </v-btn>
                </v-form>
              </template>
            </v-card>
          </template>
        </v-dialog>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HomeView',

  data () {
    return {
      isValid: false,
      show1: false,
      show2: false,
      login: '',
      password: '',
      confirm: '',
      rules: {
        required: value => !!value || 'Required.',
        max: v => (v || '').length <= 30 || 'Max 30 characters',
        min: v => v.length >= 8 || 'Min 8 characters',
        email: value => {
          const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Invalid e-mail.'
        },
        password: value => {
          const pattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/
          return pattern.test(value) || '1 uppercase, 1 lowercase symbol and 1 number'
        },
        repeat: [
          v => !!v || 'Confirm your password',
          v => v === this.password || 'The password does not match.'
        ]
      }
    }
  },

  methods: {
    async SignUp () {
      const response = await axios({
        method: 'post',
        url: process.env.VUE_APP_API_URL + '/sign-up',
        headers: {},
        data: {
          email: this.login,
          password: this.password
        }
      })
        .catch(function (error) {
          if (error.response) {
            return alert('This email has been already used')
          }
        })
      if (response !== undefined) {
        alert('The account has been created')
      }
    }
  }
}
</script>

<style>
h1, .start, .signup {
  color: #6200EE;
}
h2 {
  text-align: center;
  margin: 20px;
  color: #FFFFFFB3;
}
.start {
  font-size: 18px !important;
  height: 70px !important;
  width: 190px !important;
}
.signup {
  font-size: 14px !important;
  height: 50px !important;
  width: 110px !important;
  margin-left: 35% !important;
}
.dialog-center {
  height: 180px;
}
.icon {
  margin: 20px 125px 0;
}
.text-field {
  max-width: 400px;
  padding: 0 8% !important;
}
.v-dialog {
  height: 745px !important;
}
div.v-card.v-sheet.theme--dark {
  height: inherit !important;
}
.mdi-eye-off::before, .mdi-eye::before {
  color: #FFFFFFB3;
}
.v-banner {
  border-bottom: thin solid #6200EE !important;
}
</style>
