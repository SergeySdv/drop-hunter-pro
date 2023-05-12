<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      permanent
    >
      <v-list-item
        prepend-avatar="https://i1.sndcdn.com/artworks-000642312244-143ip1-t500x500.jpg"
        title="Хуй Жопа Залупа"
        nav
      >
      </v-list-item>

      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-home-city" title="Profiles" value="home"
                     @click="$router.push({name: 'Profiles'})"></v-list-item>
        <v-list-item prepend-icon="mdi-home-city" title="Exchange withdraw" value="Withdraw"
                     @click="$router.push({name: 'Withdraw'})"></v-list-item>
        <v-list-item prepend-icon="mdi-home-city" title="Constructor" value="Constructor"
                     @click="$router.push({name: 'Constructor'})"></v-list-item>
        <v-list-item prepend-icon="mdi-home-city" title="Processes" value="Processes"
                     @click="$router.push({name: 'Processes'})"></v-list-item>
        <v-list-item prepend-icon="mdi-home-city" title="Settings" value="Settings"
                     @click="$router.push({name: 'Settings'})"></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <notifications/>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import SideBar from "@/components/SideBar.vue";

export default defineComponent({
  components: {
    SideBar,
  },
  data() {
    return {
      drawer: true,
      rail: false,
    }
  },
  mounted() {

  }
})


</script>

<style>
app {
  font-family: 'UniNeueHeavy-Italic', serif;
  background-color: bisque;
}

.app-c {
  font-family: 'UniNeueHeavy-Italic', serif;
  background-color: bisque;
}
</style>
