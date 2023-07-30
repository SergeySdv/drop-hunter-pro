<template>
  <v-row justify="center">
    <v-dialog
      v-model="show"
      persistent
      width="1024"
    >

      <template v-slot:activator="{ props }">
        <v-icon
          icon="mdi-account-edit-outline"
          v-bind="props"
          @click="loadProfile"
        />
      </template>
      <v-card>
        <v-card-title>
          <span class="text-h5">Edit profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-form validate-on="submit" ref="form">
              <v-row>
                <v-col cols="12" sm="6" md="4">
                  <ProfileLabelInput v-model="item.label" :profile-id="profileId"/>
                </v-col>
                <v-col cols="12">
                  <ProxyInput :required="false" v-model="item.proxy"/>
                </v-col>
                <v-col cols="12">
                  <v-text-field label="user agent" v-model="item.userAgent" :disabled="true" variant="outlined"
                                hide-details/>
                </v-col>
                <v-col cols="12">
                  <v-textarea
                    hide-details
                    label="Meta"
                    type="Meta"
                    v-model="item.meta"
                    density="comfortable"
                    variant="outlined"
                  ></v-textarea>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="Close"
          >
            Close
          </v-btn>
          <v-btn
            type="submit"
            color="blue-darken-1"
            variant="text"
            :loading="updating"
            @click="Update">
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, profileService} from "@/generated/services"
import {CreateProfileRequest, OkexAccount, Profile} from "@/generated/profile";
import {tr} from "vuetify/locale";
import ProxyInput from "@/components/ProxyInput.vue";
import ProfileLabelInput from "@/components/profile/ProfileLabelInput.vue";

export default defineComponent({
  name: "EditProfile",
  components: {
    ProfileLabelInput,
    ProxyInput
  },
  emits: ['updated'],
  props: {
    profileId: {
      required: true,
      type: String
    },
  },
  data() {
    return {
      show: false,
      item: {
        id: "",
        proxy: "",
        meta: "",
        mmskId: "",
        label: "",
      } as Profile,
      mmskWalletId: "",
      saveLoading: false,
      profileLoading: false,
      updating: false,
    }
  },
  methods: {
    async proxyChanged(e: string) {
      this.item.proxy = e
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async Update() {
      this.updating = true
      if (!await this.validateForm()) {
        this.updating = false
        return
      }
      try {

        await profileService.profileServiceUpdateProfile({
          body: {
            label: this.item.label,
            profileId: this.profileId,
            meta: this.item.meta,
            proxy: this.item.proxy,
            userAgent: this.item.userAgent,
          }
        })
        this.$emit('updated')
        this.show = false
      } finally {
        this.updating = false
      }
    },
    async loadProfile() {
      try {
        this.profileLoading = true
        const res = await profileService.profileServiceGetProfile({body: {profileId: this.profileId}})
        this.item = res.profile
      } catch (e) {
        this.profileLoading = false
      }
    },
    Close() {
      this.show = false
    }
  }
})


</script>


<style>


</style>

