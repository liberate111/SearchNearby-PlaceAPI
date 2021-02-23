<template>
    <v-container grid-list-xs>
        <v-card elevation="10">
            <v-card-text>
                <div></div>
                <p class="display-1 text--primary">
                    X: {{resultXYZ.x}}
                    <v-divider></v-divider>
                    Y: {{resultXYZ.y}}
                    <v-divider></v-divider>
                    Z: {{resultXYZ.z}}
                </p>
                <p></p>
                <div class="text--primary">
                </div>
            </v-card-text>
        </v-card>
    </v-container>
</template>

<script>
import axios from "axios";
export default {
    name: "XYZ",
    data() {
        return {
            resultXYZ: {}
        }
    },
    async mounted() {
    let result = await axios.get("http://localhost:8090/xyz")
    this.resultXYZ = result.data
    },
}
</script>

<style>

</style>