<template>
  <nut-empty description="404">
    <div style="margin-top: 10px">
      <nut-button type="primary" @click="gotoBack">刷新</nut-button>
      <nut-button type="default" @click="gotoHome">返回首页</nut-button>
    </div>
  </nut-empty>
</template>
<script setup>
import { onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const gotoHome = () => {
  router.push("/");
};
const gotoBack = () => {
  router.back();
};

onMounted(() => {
  console.log("404.vue mounted");
});
</script>
