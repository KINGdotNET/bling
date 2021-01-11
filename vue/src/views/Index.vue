<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="kingdotnet.bling.bling" type="add" :fields="[ ['creator', 1, 'string'] , ['post', 2, 'string'] , ['title', 3, 'string'] , ['body', 4, 'string'] , ['ipfs', 5, 'string'] , ['parent', 6, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
