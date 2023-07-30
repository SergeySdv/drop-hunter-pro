<template>
  <v-container>
    <ol>
      <li class="mt-5 text-h6"> Add existing or generate new profiles <a href="/profiles/batch/create" target="_blank">here</a>
      </li>
      <v-img class="pb-5" src="/onboarding/getstarted/profiles-creation.png"/>
      <li class="mt-5 text-h6">When you complete filling profiles press validate</li>
      <v-img class="pb-5" src="/onboarding/getstarted/profiles-validation.png"/>
      <li class="mt-5 text-h6">If all profiles are okey press 'create', then go to <a href="/profiles" target="_blank">profiles
        page</a></li>
      <li class="mt-5 text-h6">You will see profiles you added</li>
      <v-img class="pb-5" src="/onboarding/getstarted/profiles.png"/>
      <li class="mt-5 text-h6"> For easy TopUp you can add your exchange accounts <a href="/exchange_accounts"
                                                                                     target="_blank">here</a>.
        Watch
        instruction by pressing button 'DOCS'
      </li>
      <v-img class="pb-5" src="/onboarding/getstarted/exacc-creation.png"/>
      <li class="mt-5 text-h6">After exchange account successfully added you will see it on a page</li>
      <v-img class="pb-5" src="/onboarding/getstarted/exacc-created.png"/>

      <li class="mt-5 text-h6">Press 'View' button and you will see exchange account page</li>
      <v-img class="pb-5" src="/onboarding/getstarted/exacc-opt.png"/>
      <li class="mt-5 text-h6">If you added OKEX account it is possible to add okex sub-accounts for safe withdrawal
      </li>
      <v-img class="pb-5" src="/onboarding/getstarted/okex-subb.png"/>

      <li class="mt-5 text-h6"> Go to <a href="/constructor" target="_blank">constructor section</a> and create flow you
        want
      </li>
      <v-img class="pb-5" src="/onboarding/getstarted/flow.png"/>
      <li class="mt-5 text-h6">Select profile you want to run with flow and press 'start process'</li>
      <v-img class="pb-5" src="/onboarding/getstarted/process_creation.png"/>
      <li class="mt-5 text-h6">Next page shows you process you created. Press 'start'</li>
      <v-img class="pb-5" src="/onboarding/getstarted/process_execution.png"/>
      <li class="mt-5 text-h6">If you want to keep in tune of you process - subscribe for telegram bot</li>
      <v-img class="pb-5" width="200px" src="/onboarding/getstarted/alert.png"/>

      <li class="mt-5 text-h6">When process is created you can find it on this <a href="/processes"
                                                                                  target="_blank">page</a>
      </li>
      <li class="mt-5 text-h6">Checkout <a href="/settings" target="_blank">settings page</a> to see options</li>
      <li class="mt-5 text-h6">By default you have 10$ on balance.
        When you reach the limit you can visit <a href="/billing" target="_blank">billing page</a></li>
      <li class="mt-5 text-h6">If you want to learn more about project - go <a href="/about" target="_blank">here</a>
      </li>
    </ol>

  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Zksync from "@/components/modules/Zksync.vue";
import Snapshot from "@/components/modules/Snapshot.vue";
import Orbiter from "@/components/modules/Orbiter.vue";
import LayerZero from "@/components/modules/LayerZero.vue";


export default defineComponent({
  name: "GetStarted",
  created() {
  },
})


</script>


<style>
</style>

