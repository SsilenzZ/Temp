<template>
  <v-card
    color="grey lighten-4"
    tile
  >
    <div style="flex-basis: 20%">
      <v-toolbar class="tb2 flex-fill">
        <v-spacer></v-spacer>

        <v-avatar
          color=#5110e7
          class="text-h5 font-weight-light white--text"
          style="min-width: 42px;
          max-width: 42px;
          height: 42px;
          margin-right: 10px;
          padding-left: 2px;"
        >
          {{ name.charAt(0) }}
        </v-avatar>

        <v-toolbar-title>
          {{ name }}
        </v-toolbar-title>

        <v-spacer></v-spacer>
        <v-btn icon @click="closeChatBox">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-toolbar>
    </div>
  </v-card>
</template>

<script>
export default {
  name: 'Toolbar2Comp',

  props: ['name'],

  methods: {
    closeChatBox () {
      this.$emit('closeChatBox')
    }
  }
}
</script>

<style scoped>
.tb2, .v-toolbar__content {
  height: 64px !important;
}
.row {
  margin: 0 !important;
}
</style>
