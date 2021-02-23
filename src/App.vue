<template>
  <v-app>
    <v-main>
      <v-card
    class="pa-12"
    color="indigo darken-2"
    flat
  >
    <v-card
      elevation="12"
      width="300"
    >
      <v-navigation-drawer
        floating
        permanent
      >
        <v-list
          dense
          rounded
        >
          <v-list-item
            v-for="item in items"
            :key="item.title"
            link
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <div style="width:220px" @click="$router.push({ path: `${item.path}` })">
              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </div>


          </v-list-item>
        </v-list>
      </v-navigation-drawer>
    </v-card>
      </v-card>
      

      <router-view></router-view>

    </v-main>

    <v-row class="d-flex flex-row justify-content-center align-center">
      <v-img 
      class="mt-10"
      contain  
      maxHeight="200"
      src="data:image/png;base64,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"
      />
    </v-row>

  </v-app>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      items: [
          { title: 'Home', icon: 'mdi-view-dashboard', path: "/" },
          { title: 'XYZ', icon: 'mdi-calculator-variant-outline', path: "/XYZ" },
          { title: 'MuangThongThani', icon: 'mdi-map-search', path: "/MTT" },
        ],
    }
  },
  components: {
  },
};
</script>
