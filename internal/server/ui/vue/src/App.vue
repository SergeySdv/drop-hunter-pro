<template>
  <v-app class="font">
    <v-navigation-drawer
      v-if="userLoggedIn"
      :location="navBarLocation"
      expand-on-hover
      permanent
      :rail="rail"
      @update:rail="railUpdated"
      @click="rail = true"
    >
      <v-list-item
        prepend-avatar="/favicon-32x32.png"
        :title="userEmail"
        nav
        @click="rail = false"
      >
        <div style="height: 20px" v-if="ass">Balance: {{ ass }}</div>
        <template v-slot:append>
          <v-btn
            v-if="rail"
            variant="text"
            icon="mdi-logout"
            @click="logout"
          ></v-btn>
        </template>
      </v-list-item>

      <v-divider/>

      <v-list density="compact" nav height="vh 90">
        <v-list-item prepend-icon="mdi-home-floor-g" color="green" title="General" value="general"
                     @click="$router.push({name: 'LandingPage'})"></v-list-item>
        <v-list-item prepend-icon="mdi-star-circle-outline" color="green" title="Modules" value="airdrops"
                     @click="$router.push({name: 'Modules'})"></v-list-item>
        <v-list-item prepend-icon="mdi-star-circle-outline" color="green" title="Get started" value="getstarted"
                     @click="$router.push({name: 'GetStarted'})"></v-list-item>
        <v-list-item prepend-icon="mdi-account-multiple-outline" color="green" title="Profiles" value="home"
                     @click="$router.push({name: 'Profiles'})"></v-list-item>
        <v-list-item prepend-icon="mdi-account-cash-outline" color="green" title="Exchanges" value="Withdraw"
                     @click="$router.push({name: 'Withdraw'})"></v-list-item>
        <v-list-item prepend-icon="mdi-account-hard-hat-outline" color="green" title="Constructor" value="Constructor"
                     @click="$router.push({name: 'Constructor'})"></v-list-item>
        <v-list-item prepend-icon="mdi-blender-outline" color="green" title="Processes" value="Processes"
                     @click="$router.push({name: 'Processes'})"></v-list-item>
        <v-list-item prepend-icon="mdi-cog" title="Settings" color="green" value="Настройки"
                     @click="$router.push({name: 'Settings'})"></v-list-item>
        <v-list-item prepend-icon="mdi-cash-multiple" title="Billing" color="green" value="Мани"
                     @click="$router.push({name: 'Billing'})"></v-list-item>
        <v-list-item prepend-icon="mdi-information-outline" title="About" color="green" value="Обо мне"
                     @click="$router.push({name: 'About'})"></v-list-item>


        <div v-if="coc" class="h-auto py-4 pl-3">
          <div class="flex-column  justify-start">

            <a href="https://t.me/drop_hunter_alert_bot" target="_blank">
              <div style="height: 20px" class="d-inline-flex">
                <v-img src="/icons/telegram.png"/>
                <span class="ml-2">Bot</span>
              </div>
            </a>
            <br/>
            <a href="https://t.me/+WX8iCLaJBelhNTVi" target="_blank">
              <div style="height: 20px" class="d-inline-flex">
                <v-img src="/icons/telegram.png"/>
                <span class="ml-2">Support</span>
              </div>
            </a>
          </div>


          <v-list-item v-if="coc" class="d-flex pt-8" style="font-size: 14px">
            <div v-if="impact.top">
              <i>{{ `24 Hour statistics` }}
                <br/>
                <div>{{ ` total - ${impact.total} tx` }}</div>
                <div>{{ ` top - ${impact.top} tx` }}</div>
                <div>{{ ` me - ${impact.my} tx` }}</div>
              </i>
            </div>
          </v-list-item>
        </div>

      </v-list>

    </v-navigation-drawer>
    <v-main>
      <Login v-if="!userLoggedIn" class="mt-4 mr-4" style="position: absolute; top: 0; right: 0; "/>
      <notifications/>
      <router-view/>
      <Snackbar/>
    </v-main>

  </v-app>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import SideBar from "@/components/SideBar.vue";
import Login from "@/components/Login.vue";
import {helperService} from "@/generated/services";
import {mapActions, mapStores} from "pinia";
import {useUserStore, useSysStore} from "@/plugins/pinia";
import Snackbar from "@/components/Snackbar.vue";

export default defineComponent({
  components: {
    Snackbar,
    Login,
    SideBar,
  },
  data() {
    return {
      drawer: true,
      rail: true,
      coc: false,
    }
  },
  computed: {
    navBarLocation(): "end" | "start" | "left" | "top" | "bottom" | "right" {
      return 'left'
    },
    ...mapStores(useUserStore),
    impact() {
      return this.userStore.impact
    },
    userLoggedIn(): boolean {
      return this.userStore.login
    },
    userEmail(): string {
      return this.userStore.email
    },
    ass(): string {
      return this.userStore.ass
    },
    ...mapActions(useUserStore, ['syncUser', "syncDailyImpact"])


  },
  methods: {
    railUpdated() {
      this.coc = !this.coc
    },
    logout() {
      document.cookie.split(";").forEach(function (c) {
        document.cookie = c.replace(/^ +/, "").replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
      });
      localStorage.clear();
      window.location.reload()
    }
  },
  async mounted() {


    const period = import.meta.env.DEV ? 60_000 : 5000

    await this.userStore.syncUser()
    await this.userStore.syncDailyImpact()

    setInterval(async () => {
      await this.userStore.syncUser()
      await this.userStore.syncDailyImpact()
    }, period)
  }
})


</script>

<style>

.font {
  /*font-family: "Roboto";*/
}
</style>
