<template>
  <v-row justify="center">
    <v-dialog
      v-model="show"
      persistent
      width="1024"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">Create profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-form validate-on="submit" ref="form">
              <v-row>
                <v-col cols="12" sm="6" md="4">
                  <ProfileLabelInput v-model="item.label"/>
                </v-col>
                <v-col cols="12" sm="12" md="8">
                  <v-text-field
                    ref="pk-input"
                    :rules="mmskRules()"
                    @input="mmskPkChanged"
                    v-model="item.mmskPk"
                    :messages="mmskWalletId"
                    density="comfortable"
                    variant="outlined"
                    label="metamask wallet private key*"
                    hint="metamask wallet private key"/>
                </v-col>
                <v-col cols="12">
                  <ProxyInput :required="false" v-model="item.proxy"/>
                </v-col>
                <v-col cols="12">
                  <v-textarea
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
            :loading="saveLoading"
            @click="CreateProfile">
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
import {CreateProfileRequest, Profile} from "@/generated/profile";
import {tr} from "vuetify/locale";
import ProxyInput from "@/components/ProxyInput.vue";
import ProfileLabelInput from "@/components/ProfileLabelInput.vue";

export default defineComponent({
  name: "CreateProfile",
  components: {
    ProfileLabelInput,
    ProxyInput
  },
  props: {
    showProp: Boolean,
  },
  emits: ['updated', 'change'],
  watch: {
    showProp: {
      immediate: true,
      handler(newValue, oldValue) {
        this.show = newValue
      },
    }
  },
  data() {
    return {
      show: this.showProp,
      item: {} as CreateProfileRequest,
      mmskWalletId: "",
      saveLoading: false,
    }
  },
  methods: {
    resetForm() {
      this.item.proxy = ""
      this.item.label = ""
      this.item.meta = ""
      this.item.mmskPk = ""
      this.saveLoading = false
    },
    mmskRules() {
      return [
        (v: string) => {
          if (this.item.mmskPk) {
            if (this.item.mmskPk.length < 64) {
              return 'PK must is less than 64 characters'
            }
            if (this.item.mmskPk.length > 64) {
              return 'PK must is more than 64 characters'
            }
            return true
          }
          return 'PK is required'
        },
        async (v: string) => {
          if (this.item.mmskPk && this.item.mmskPk.length == 64) {
            const res = await helperService.helperServiceValidatePk({body: {pk: this.item.mmskPk}})
            if (res.walletId) {
              return true
            } else {
              return "invalid PK"
            }
          }
          return "API error, call an ambulance!"
        }
      ]
    },
    async proxyChanged(e: string) {
      this.item.proxy = e
    },
    async mmskPkChanged(e: InputEvent) {
      await this.validatePk()
      if (this.item.mmskPk.length == 64) {
        const res = await helperService.helperServiceValidatePk({body: {pk: this.item.mmskPk}})
        if (res.walletId) {
          this.mmskWalletId = "wallet id: " + res.walletId
        }
      } else {
        this.mmskWalletId = ""
      }
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
    async validatePk(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['pk-input'].validate()

      return valid
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"].form.validate()

      return valid
    },
    async CreateProfile() {
      this.saveLoading = true
      if (!(await this.validateForm())) {
        this.saveLoading = false
        return
      }

      try {
        await profileService.profileServiceCreateProfile({body: this.item})
        this.$emit("change", false)
        this.$emit("updated")
        this.resetForm()
      } catch (e) {
        this.saveLoading = false
      }

    },
    Close() {
      this.$emit("change", false)
    }
  }
})


</script>


<style>


</style>

