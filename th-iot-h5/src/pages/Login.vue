<template>
  <div class="login-wrapper">
    <!-- Logo -->
    <div class="login-logo">欢迎</div>
    <p>请登录您的账户</p>

    <!-- 表单 -->
    <div class="login-form">
      <!-- 用户名 -->
      <div class="form-item">
        <nut-input
          v-model="formData.username"
          placeholder="用户名/邮箱/手机号"
          clearable
          type="text"
          class="login-input"
          @blur="validateField('username')"
        />
        <div class="error-msg" v-if="errors.username">
          {{ errors.username }}
        </div>
      </div>

      <!-- 密码 -->
      <div class="form-item">
        <nut-input
          v-model="formData.password"
          placeholder="密码"
          clearable
          type="password"
          class="login-input"
          @blur="validateField('password')"
        />
        <div class="error-msg" v-if="errors.password">
          {{ errors.password }}
        </div>
      </div>

      <!-- 按钮 -->
      <div class="login-actions">
        <nut-button type="primary" block @click="login" :loading="isLoading">
          登 录
        </nut-button>
        <nut-button type="default" block @click="goRegister">
          去注册
        </nut-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import { showNotify } from "@nutui/nutui";

const router = useRouter();
const formData = reactive({ username: "", password: "" });
const isLoading = ref(false);
const errors = reactive({ username: "", password: "" });

// 校验规则
const rules = {
  username: (val) => {
    if (!val) return "请输入用户名";
    const phoneReg = /^1\d{10}$/; // 简单手机号校验
    const emailReg = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; // 邮箱校验
    if (!phoneReg.test(val) && !emailReg.test(val)) return "请输入有效的手机号或邮箱";
    return "";
  },
  password: (val) => {
    if (!val) return "请输入密码";
    if (val.length <= 6) return "密码必须大于6位";
    return "";
  },
};


// 单字段校验
const validateField = (prop) => {
  errors[prop] = rules[prop](formData[prop]);
  if (errors[prop]) showNotify.warn(errors[prop]);
};

// 登录
const login = async () => {
  // 全字段校验
  errors.username = rules.username(formData.username);
  errors.password = rules.password(formData.password);

  if (errors.username || errors.password) return;

  isLoading.value = true;
  try {
    const res = await axios.post("/auth/login", formData);
    if (res.data.code === 0) {
      const data = res.data.data;
      localStorage.setItem("token", data.accessToken);
      localStorage.setItem("refreshToken", data.refreshToken);
      localStorage.setItem("expiresAt", data.expiresAt);
      localStorage.setItem("uid", data.id);
      localStorage.setItem("username", data.username);
      showNotify.success("登录成功");
      router.push("/");
    } else {
      showNotify.warn(res.data.message || "登录失败");
    }
  } catch (err) {
    console.error(err);
    showNotify.warn("网络错误，请稍后重试");
  } finally {
    isLoading.value = false;
  }
};

// 去注册
const goRegister = () => router.push("/register");
</script>

<style scoped>
.login-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  padding-top: 4rem; /* 稍微靠上 */
}

/* Logo */
.login-logo {
  width: 8rem;
  height: 8rem;
  background-color: #10a78e;
  color: #fff;
  font-size: 2.5rem;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  margin-bottom: 0.8rem; /* 让下方文字更贴近 */
}

/* 副标题 */
.login-wrapper p {
  margin: 0;
  margin-bottom: 2rem; /* 与输入框保持合适间距 */
  font-size: 1rem;
  color: #666;
  text-align: center;
}

/* 表单整体 */
.login-form {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 移除 NutUI 的多余样式 */
.login-form >>> .nut-form-item {
  border: none !important;
  padding: 0 !important;
  background: none !important;
}

/* 输入框 */
.login-input >>> input {
  width: 100%;
  border-radius: 24px !important;
  padding: 12px 16px !important;
  font-size: 15px !important;
  border: 1px solid #ccc !important;
  background: #fff !important;
  box-sizing: border-box;
}

/* 聚焦效果 */
.login-input >>> input:focus {
  border-color: #10a78e !important;
}

/* 错误提示 */
.error-msg {
  font-size: 12px;
  color: #ff4d4f;
  margin-top: 2px;
}

/* 按钮 */
.login-actions {
  display: flex;
  flex-direction: column;
  gap: 14px;
  margin-left: 20px;
  margin-right: 20px;
  margin-top: 10px;
}

/* 按钮圆角 */
.nut-button {
  border-radius: 24px !important;
  font-size: 16px !important;
}
</style>
