<template>
  <v-card>
    <v-card-title primary-title>
      <v-container>
        <v-row
          class="d-flex flex-row align-center"
          v-for="(item, index) in places"
          :key="index"
        >
          <v-col cols="12"  class="d-flex flex-row align-center pa-0">
          
          <v-list-item>
          <v-list-item-content>
              <v-list-item-title>{{item.name}} --- Rating: {{ item.rating }}</v-list-item-title>
          </v-list-item-content>
          </v-list-item>

          <v-list-item>
            <v-list-item-content>
              <v-list-item-title>Address: {{ item.vicinity }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          </v-col>

          <v-col>
            <v-divider></v-divider>
          </v-col>
        </v-row>
      </v-container>
    </v-card-title>
  </v-card>
</template>

<script>
export default {
  name: "ListResult",
  props: ["places"]
};
</script>

<style></style>
