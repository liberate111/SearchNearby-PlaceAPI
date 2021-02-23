<template>
  <div class="home">
    <v-container>
      <Search @onSubmit="search" />
      <h1>{{testname}}</h1>
      <ListResult :places="places" />
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import Search from "@/components/Search"
import ListResult from "@/components/ListResult"
export default {
  name: "Home",
  components: {
    Search,
    ListResult,
  },
  data() {
    return {
      textField: "",
      places: [],
      testname: "",
    };
  },
  methods: {
    async search(value) {
      let nPlace = value;
      let result = await axios.get(`http://localhost:8090/nearby?name=${nPlace}`);
      this.places = result.data.Results;
      // console.log(result.data);
    }
  },
};
</script>