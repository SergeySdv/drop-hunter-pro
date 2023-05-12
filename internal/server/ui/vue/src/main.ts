/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import {createApp} from 'vue'

// Plugins
import Notifications from '@kyvg/vue3-notification'
import vuetify from "@/plugins/vuetify";
import router from "@/router";
import pinia from "@/store";

const app = createApp(App)
  .use(vuetify)
  .use(router)
  .use(pinia)
  .use(Notifications)


app.mount('#app')
