<template>
  <div>
    <Loader v-if="loading"/>
    <div v-else>

      <div class="d-flex justify-space-between">
        <div class="mx-2 my-2">
          <v-checkbox
            class="mr-3"
            density="compact"
            label="Address"
            v-model="showWalletAddrs"
            hide-details
          />
          <v-checkbox
            class="mr-3"
            density="compact"
            label="Labels"
            v-model="showLabels"
            hide-details
          />
        </div>
        <div class="mr-3" style="width: auto">
          <!--              <ProfilesExport v-if="list.length > 0" class="pr-2"/>-->
          <ProfilesDocs class="mx-2 my-2"/>
          <v-btn class="mx-2 my-2" v-if="selectedSome" color="black" @click=DeleteProfiles>Delete</v-btn>
          <v-btn class="mx-2 my-2" @click="$router.push({name: 'CreateProfileBatch'})">Create</v-btn>
        </div>
      </div>


      <v-table fixed-header height="92vh">
        <thead>
        <tr>
          <th></th>
          <th class="text-left" v-if="showWalletAddrs">Metamask</th>
          <th class="text-left" v-if="showLabels">Label</th>
          <th class="text-left">Profile</th>
          <th class="text-left">Proxy</th>
          <th class="text-left">Meta</th>
          <th class="text-left">Accounts</th>
          <th class="text-left">CreatedAt</th>
          <th class="text-left">Actions</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in getList" :key="item.id">
          <td>
            <CheckBox style="height: 40px" density="compact" :id="item.id" focused
                      @CheckboxChanged="CheckboxChanged"></CheckBox>
          </td>
          <td v-if="showWalletAddrs"><span>{{ item.mmskId }}</span></td>
          <td v-if="showLabels"><span>{{ item.label }}</span></td>

          <td>
            <ProfileCard :label="item.num" :profile-id="item.id"/>
          </td>
          <td>
            <BtnCheckProxy v-if="item.proxy" :proxy="item.proxy"/>
            <span v-else>{{ 'disabled' }}</span>
          </td>
          <td>
            {{ item.meta }}
          </td>
          <td>
            <v-chip rounded variant="outlined" v-if="item.okexAccount">Okex</v-chip>
          </td>
          <td>{{ formatTime(item.createdAt) }}</td>
          <td>
            <EditProfile :profile-id="item.id" @updated="UpdateList"/>
          </td>
        </tr>
        </tbody>
      </v-table>
      <!--      <CreateProfile :showProp="showCreateProfileDialog" @change="CreateProfileChange" @updated="UpdateList"/>-->
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {profileService} from "@/generated/services"
import {Profile} from "@/generated/profile";
import CreateProfile from "@/components/profile/CreateProfile.vue";
import CheckBox from "@/components/CheckBox.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import BtnCheckProxy from "@/components/BtnCheckProxy.vue";
import EditProfile from "@/components/profile/EditProfile.vue";
import ProfilesDocs from "@/components/profile/ProfilesDocs.vue";
import ProfilesExport from "@/components/profile/ProfilesExport.vue";
import Loader from "@/components/Loader.vue";
import {formatTime} from "../helper";

export default defineComponent({
  name: "Profiles",
  components: {
    Loader,
    ProfilesExport,
    ProfilesDocs,
    EditProfile,
    BtnCheckProxy,
    CheckBox,
    CreateProfile,
    ProfileCard
  },
  props: {},
  data() {
    return {
      showWalletAddrs: false,
      list: [] as Profile[],
      showCreateProfileDialog: false,
      selected: new Set<string>(),
      loading: true,
      loadingError: false,
      showLabels: false,
    }
  },
  computed: {
    getList(): Profile[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    }
  },
  methods: {
    formatTime,
    CheckboxChanged(id: string, value: boolean) {
      if (value) {
        this.selected.add(id)
      } else {
        this.selected.delete(id)
      }
    },
    async UpdateList() {
      try {
        this.loadingError = false
        this.loading = true
        const res = await profileService.profileServiceListProfile()
        this.list = res.profiles
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }

    },
    CreateProfile() {
      this.showCreateProfileDialog = true
    },
    CreateProfileChange(b: boolean) {
      this.showCreateProfileDialog = b
    },
    async DeleteProfiles() {
      for (let id of this.selected.keys()) {
        await profileService.profileServiceDeleteProfile({body: {id: id}})
      }
      await this.UpdateList()
    }
  },
  created() {
    this.UpdateList()
  }
})
</script>

<style>
</style>

