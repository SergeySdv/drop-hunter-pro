<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
    location="end"
  >

    <template v-slot:activator="{ props }">
      <v-btn variant="outlined" density="comfortable" v-bind="props">{{ profile.label }}
      </v-btn>
    </template>

    <template v-slot:default="{ props }">
      <v-card class="px-4 py-4">
        <b> {{ `ID: ${profile.mmskId}` }}</b>
        <br/>
        <a :href="`https://debank.com/profile/${profile.mmskId}`" target="_blank">Debank</a>
        <v-card-item>
          <Balance :profile-id="profileId" :network="Network.ARBITRUM"/>
        </v-card-item>
        <v-card-item>
          <Balance :profile-id="profileId" :network="Network.Etherium"/>
        </v-card-item>
        <v-card-item>
          <Balance :profile-id="profileId" :network="Network.OPTIMISM"/>
        </v-card-item>
        <v-card-item>
          <Balance :profile-id="profileId" :network="Network.BinanaceBNB"/>
        </v-card-item>
        <v-card-item>
          <Balance :profile-id="profileId" :network="Network.POLIGON"/>
        </v-card-item>
        <v-card-item>
          <Balance :profile-id="profileId" :network="Network.AVALANCHE"/>
        </v-card-item>
        <v-card-actions>
        </v-card-actions>
      </v-card>
    </template>
  </v-menu>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {profileService} from "@/generated/services"
import Balance from "@/components/Balance.vue";
import {Network, Profile} from "@/generated/profile";

export default defineComponent({
  name: "ProfileCard",
  components: {Balance},
  props: {
    profileId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      profile: {} as Profile,
      menu: false,
    }
  },
  computed: {
    Network() {
      return Network
    }
  },
  methods: {
    async loadProfile() {
      await profileService.profileServiceGetBalance({body: {profileId: this.profileId, network: Network.ARBITRUM}})
    }
  },
  async created() {
    const res = await profileService.profileServiceGetProfile({body: {profileId: this.profileId}})
    this.profile = res.profile
  }
})


</script>


<style>


</style>

