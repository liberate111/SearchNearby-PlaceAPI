<template>
  <v-container >
    <ListResult :places="places" />
  </v-container>
</template>

<script>
import axios from "axios";
import ListResult from "@/components/ListResult.vue";
export default {
  name: "MTT",
  data() {
    return {
      places: [],
    };
  },
  async mounted() {
    let nPlace = "MuangThongThani";
    let result = await axios.get(`http://localhost:8090/nearby?name=${nPlace}`);
    this.places = result.data.Results;
    // console.log(result.data);
  },
  components: { ListResult },
};
</script>

<style></style>
