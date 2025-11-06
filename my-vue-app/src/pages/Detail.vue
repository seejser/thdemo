<template>
  <nut-navbar title="设备详情" left-show @click-back="onClick"></nut-navbar>
  <h1>Hello App!</h1>
  <p><strong>Current route path:</strong> {{ $route.fullPath }}</p>
  <nav>
    <RouterLink to="/">Go to Home</RouterLink>
    <RouterLink to="/detail">Go to Detail</RouterLink>
  </nav>
  <main>
    <RouterView />
  </main>
</template>
<script setup>
import { useRouter } from "vue-router";

const router = useRouter();

const onClick = () => {
  console.log("[Navbar]: on click title");
  router.push("/");
};
</script>
