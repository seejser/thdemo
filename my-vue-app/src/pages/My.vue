<template>
  <div class="container">
    <nut-avatar> <img :src="url" /> </nut-avatar>
    <h1>春天的羊</h1>
    <HyTabBar />
  </div>
</template>
<script setup>
import { My } from "@nutui/icons-vue";

import HyTabBar from "./../components/Hytabbar.vue";

const url =
  "https://img12.360buyimg.com/imagetools/jfs/t1/196430/38/8105/14329/60c806a4Ed506298a/e6de9fb7b8490f38.png";
</script>
