<template>
  <div class="d-flex justify-space-between">
    <v-btn variant="tonal" color="blue" width="230px" height="35px" :href="loginGoogleHref()">
      <v-img width="25" src="/google_oauth_icon.png"/>
      login with google
    </v-btn>
  </div>
</template>


<script lang="ts">

import {defineComponent} from 'vue';
import {loginGoogleHref} from "@/generated/services"
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";


export default defineComponent({
  name: "Login",
  components: {},
  computed: {
    ...mapStores(useUserStore),
    userLoggedIn(): boolean {
      return this.userStore.$state.login
    }
  },
  methods: {
    loginGoogleHref,
  },
  created() {

  }
})


</script>


<style>


</style>

