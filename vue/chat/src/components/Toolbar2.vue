<template>
  <v-card
    color="grey lighten-4"
    height="200px"
    tile
  >
    <div style="flex-basis: 20%">
      <v-toolbar class="tb2">
        <v-spacer></v-spacer>
        <v-btn icon>
          <v-icon>mdi-dots-vertical</v-icon>
        </v-btn>
      </v-toolbar>
    </div>
  </v-card>
</template>

<script>
export default {
  name: 'Toolbar2Comp'
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
