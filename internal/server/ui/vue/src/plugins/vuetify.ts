/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import {createVuetify, ThemeDefinition} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {md3} from 'vuetify/blueprints'
import {VSkeletonLoader} from "vuetify/labs/VSkeletonLoader";
// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides

const myCustomLightTheme: ThemeDefinition = {
  dark: false,
  colors: {
    background: '#E6FFFD',
    surface: '#AEE2FF',
    primary: '#ACBCFF',
    'primary-darken-1': '#3700B3',
    secondary: '#B799FF',
    'secondary-darken-1': '#018786',
    error: '#B00020',
    info: '#2196F3',
    success: '#4CAF50',
    warning: '#FB8C00',
  },
}
export default createVuetify({
  components: {...components, VSkeletonLoader},
  directives,
  blueprint: md3,
  // icons: {
  //   defaultSet: 'mdi'
  // }
})

