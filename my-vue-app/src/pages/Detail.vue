<template>
  <nut-navbar title="设备详情" left-show @click-back="onClick"></nut-navbar>
  <h1>{{$route.params.id }}</h1>
  <p><strong>Current route path:</strong> {{ $route.fullPath }}</p>
  
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
