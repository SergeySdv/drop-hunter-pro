// Components
import App from './App.vue'

// Composables
import {createApp} from 'vue'

// Plugins
import Notifications from '@kyvg/vue3-notification'
import vuetify from "@/plugins/vuetify";
import router from "@/router";

import '@mdi/font/css/materialdesignicons.css'
import {pinia} from "@/plugins/pinia";


const app = createApp(App)
  .use(vuetify)
  .use(router)
  .use(Notifications)
  .use(pinia)


app.mount('#app')
