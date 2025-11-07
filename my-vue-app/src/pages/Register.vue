<template>
  <div class="container">
    <div class="top-container">
      <nut-space direction="vertical" fill align="center">
        <nut-avatar size="large"> 欢迎 </nut-avatar>
        <p>请注册</p>
      </nut-space>
    </div>
    <nut-row>
      <nut-col :span="24">
        <nut-input v-model="username" placeholder="请输入您的用户名" />
      </nut-col>
      <nut-col :span="24">
        <nut-input v-model="email" placeholder="请输入您的邮箱" />
      </nut-col>
      <nut-col :span="24">
        <nut-row>
          <nut-col :span="12">
            <nut-input v-model="calcCode" placeholder="请输入图片中的验证码" />
          </nut-col>
          <nut-col :span="8">
            <nut-image :src="url" width="100" height="34" />
          </nut-col>
        </nut-row>
      </nut-col>
      <nut-col :span="24">
        <nut-row>
          <nut-col :span="12">
            <nut-input
              v-model="verifyCode"
              placeholder="请输入您的邮箱验证码"
            />
          </nut-col>
          <nut-col :span="8">
            <nut-button plain type="info" @click="sendCode"
              >发送验证码</nut-button
            >
          </nut-col>
        </nut-row>
      </nut-col>
      <nut-col :span="24">
        <nut-input
          v-model="password"
          placeholder="请输入您的密码"
          type="password"
        />
      </nut-col>
      <nut-col :span="24">
        <nut-input
          v-model="varpassword"
          placeholder="请确认您的密码"
          type="password"
        />
      </nut-col>
      <nut-col :span="24">
        <nut-space direction="vertical" fill>
          <nut-button type="primary" @click="register" class="btn"
            >注册</nut-button
          >
          <nut-button type="default" @click="login" class="btn"
            >登录</nut-button
          >
        </nut-space>
      </nut-col>
    </nut-row>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { showNotify } from "@nutui/nutui";

const router = useRouter();
const username = ref("");
const password = ref("");
const varpassword = ref("");
const email = ref("");
const calcCode = ref("");
const verifyCode = ref("");

const url =
  "https://img10.360buyimg.com/ling/jfs/t1/181258/24/10385/53029/60d04978Ef21f2d42/92baeb21f907cd24.jpg";

const login = () => {
  router.push("/login");
};
const register = () => {
  if (!username.value || !password.value) {
    showNotify.warn("用户名或密码不能为空");
    return;
  }
  if (!calcCode.value) {
    showNotify.warn("验证码不能为空");
    return;
  }
  if (!verifyCode.value) {
    showNotify.warn("验证码不能为空");
    return;
  }
  if (password.value !== varpassword.value) {
    showNotify.warn("两次输入的密码不一致");
    return;
  }
  router.push("/login");
};
//发送验证码
const sendCode = () => {
  if (!calcCode.value) {
    showNotify.warn("验证码不能为空");
    return;
  }
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
