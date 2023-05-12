<template>
  <div>
    <div class="d-flex justify-end">
      <v-btn v-if="selectedSome" color="black" @click=DeleteProfiles class="ma-3">Delete profiles</v-btn>
      <v-btn @click=CreateProfile class="ma-3">Create profile</v-btn>
    </div>

    <v-table>
      <thead>
      <tr>
        <th></th>
        <th class="text-left">Profile</th>
        <th class="text-left">Metamask wallet key</th>
        <th class="text-left">Proxy</th>
        <th class="text-left">Meta</th>
        <th class="text-left">Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in getList" :key="item.id">
        <td>
          <CheckBox style="height: 40px" density="compact" :id="item.id" focused
                    @CheckboxChanged="CheckboxChanged"></CheckBox>
        </td>
        <td>
          <ProfileCard :profile-id="item.id"/>
        </td>
        <td>{{ item.mmskId }}</td>
        <td>
          <BtnCheckProxy v-if="item.proxy" :proxy="item.proxy"/>
          {{ item.proxy || 'disabled' }}
        </td>
        <td>{{ item.meta }}</td>
        <td>
          <EditProfile :profile-id="item.id" @updated="UpdateList"/>
        </td>
      </tr>
      </tbody>
    </v-table>
    <CreateProfile :showProp="showCreateProfileDialog" @change="CreateProfileChange" @updated="UpdateList"/>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {profileService} from "@/generated/services"
import {Profile} from "@/generated/profile";
import CreateProfile from "@/components/CreateProfile.vue";
import CheckBox from "@/components/CheckBox.vue";
import ProfileCard from "@/components/ProfileCard.vue";
import BtnCheckProxy from "@/components/BtnCheckProxy.vue";
import EditProfile from "@/components/EditProfile.vue";

export default defineComponent({
  name: "Profiles",
  components: {
    EditProfile,
    BtnCheckProxy,
    CheckBox,
    CreateProfile,
    ProfileCard
  },
  props: {},
  data() {
    return {
      list: [] as Profile[],
      showCreateProfileDialog: false,
      selected: new Set<string>(),
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
    CheckboxChanged(id: string, value: boolean) {
      if (value) {
        this.selected.add(id)
      } else {
        this.selected.delete(id)
      }
    },
    async UpdateList() {
      const res = await profileService.profileServiceListProfile()
      this.list = res.profiles
      this.selected = new Set<string>()
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

