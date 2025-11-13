<template>
  <div class="container">
    <div class="top-container">
      <nut-space direction="vertical" fill align="center">
        <nut-avatar size="large"> 欢迎 </nut-avatar>
        <p>请登录</p>
      </nut-space>
    </div>
    <nut-row>
      <!-- 用户名 -->
      <nut-col :span="24">
        <nut-input v-model="username" placeholder="请输入您的用户名" />
      </nut-col>

      <!-- 密码 -->
      <nut-col :span="24">
        <nut-input
          v-model="password"
          placeholder="请输入您的密码"
          type="password"
          clearable
        />
      </nut-col>

      <!-- 操作按钮 -->
      <nut-col :span="24">
        <nut-space direction="vertical" fill>
          <nut-button type="primary" @click="doLogin" class="btn">
            登录
          </nut-button>
          <nut-button type="default" @click="goRegister" class="btn">
            注册
          </nut-button>
        </nut-space>
      </nut-col>
    </nut-row>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { showNotify } from "@nutui/nutui";
import { login as apiLogin } from "@/api/auth";

const router = useRouter();
const username = ref("");
const password = ref("");

// 登录逻辑
const doLogin = async () => {
  if (!username.value || !password.value) {
    return showNotify.warn("用户名或密码不能为空");
  }

  try {
    const res = await apiLogin({
      username: username.value,
      password: password.value,
    });
    console.log("res:", res);
    if (res.data.code === 0 && res.data.data) {
      // 登录成功，保存 token
      const { access_token, refresh_token } = res.data.data;
      localStorage.setItem("token", access_token);
      localStorage.setItem("refreshToken", refresh_token);
      showNotify.success("登录成功");
      router.push("/"); // 跳转首页
    } else {
      showNotify.error("登录失败");
    }
  } catch (err) {
    console.error(err);
    showNotify.error("登录异常，请稍后重试");
  }
};

// 跳转注册页
const goRegister = () => router.push("/register");
</script>

<style scoped>
.container {
  padding: 16px;
}

.top-container {
  padding: 12px 16px;
  margin-top: 144px;
  display: flex;
  align-items: center;
  flex-direction: column;
  border-bottom: 1px solid #eee;
}

.top-container p {
  font-size: 14px;
  color: #666;
  margin: 0;
}

.btn {
  width: 100% !important;
}
</style>
