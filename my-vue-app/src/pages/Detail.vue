<template>
  <nut-navbar title="设备详情" left-show @click-back="onClick"></nut-navbar>
  <h1>设备号：</h1>
 
  
</template>
<script setup>
import { onMounted } from 'vue'
import { useRouter } from "vue-router";

const router = useRouter();

const onClick = () => {
  console.log("[Navbar]: on click title");
  router.push("/");
};
console.log("Detail");
onMounted(() => {
  console.log("Detail.vue mounted");
})
</script>
