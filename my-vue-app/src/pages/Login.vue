<template>
  <div class="container">
    <div class="top-container">
      <nut-space direction="vertical" fill align="center">
        <nut-avatar size="large"> 欢迎 </nut-avatar>
        <p>请登录</p>
      </nut-space>
    </div>
    <nut-row>
      <nut-col :span="24">
        <nut-input v-model="username" placeholder="请输入您的用户名" />
      </nut-col>
      <nut-col :span="24">
        <nut-input
          v-model="password"
          placeholder="请输入您的密码"
          type="password"
        />
      </nut-col>
      <nut-col :span="24">
        <nut-space direction="vertical" fill>
          <nut-button type="primary" @click="login" class="btn"
            >登录</nut-button
          >
          <nut-button type="default" @click="register" class="btn"
            >注册</nut-button
          >
        </nut-space>
      </nut-col>
    </nut-row>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { showNotify } from '@nutui/nutui'

const router = useRouter();
const username = ref("");
const password = ref("");
const login = () => {
  if (!username.value || !password.value) {
    showNotify.warn('用户名或密码不能为空');
    return;
  }
  console.log("username:", username.value, "password:", password.value);
  router.push("/");
};
const register = () => {
  router.push("/register");
};

onMounted(() => {
  console.log("mounted");
});
</script>
<style>
.container {
  padding: 16px;
  .top-container {
    /* background-color: #f5f5f5; */
    padding: 12px 16px;
    border-bottom: 1px solid #eee;
    margin-top: 144px;
    display: flex;
    align-items: center;
    gap: 12px;
    p {
      font-size: 14px;
      color: #666;
    }
  }
}
.btn {
  width: 100% !important;
}
</style>
