<template>
  <nut-config-provider :theme="theme">
    <div :class="themeClass">
      <RouterView />
    </div>
  </nut-config-provider>
</template>
<script setup>
import { ref, computed } from "vue";
const theme = ref("");
const themeClass = computed(() =>
  theme.value === "dark" ? "dark-theme" : "light-theme"
);
// const change = (v) => {
//   theme.value = v ? "dark" : "light";
// };
</script>
<style></style>
