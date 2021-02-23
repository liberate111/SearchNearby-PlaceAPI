<template>
  <v-row>
    <v-col cols="5">
      <v-text-field
        name="name"
        label="Enter the location..."
        id="id"
        v-model="textField"
        @keyup.enter="search"
      ></v-text-field>
    </v-col>

    <v-col cols="2">
      <v-btn block color="success" @click="search">Search</v-btn>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "Search",
  data() {
    return {
      textField: ""
    };
  },
  methods: {
    search() {
      let name = this.textField
      this.$emit("onSubmit", name)
      this.textField = ""
    }
  }
};
</script>

<style></style>
