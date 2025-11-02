<template>
  <div class="register-wrapper">
    <!-- Logo -->
    <div class="login-logo">欢迎</div>
    <p>请注册一个新账户</p>

    <!-- 注册表单 -->
    <div class="login-form">
      <!-- 邮箱 -->
      <div class="form-item">
        <nut-input
          v-model="formData.email"
          placeholder="请输入邮箱"
          clearable
          type="text"
          class="login-input"
          @blur="validateField('email')"
        />
        <div class="error-msg" v-if="errors.email">
          {{ errors.email }}
        </div>
      </div>

      <!-- 验证码 -->
      <div class="form-item code-item">
        <nut-input
          v-model="formData.code"
          placeholder="输入验证码"
          clearable
          type="text"
          class="login-input"
        />
        <nut-button
          size="small"
          type="primary"
          class="send-code-btn"
          :disabled="countdown > 0"
          @click="sendCode"
        >
          {{ countdown > 0 ? `${countdown}s` : '发送验证码' }}
        </nut-button>
      </div>
      <div class="error-msg" v-if="errors.code">
        {{ errors.code }}
      </div>

      <!-- 密码 -->
      <div class="form-item">
        <nut-input
          v-model="formData.password"
          placeholder="设置密码（大于6位）"
          clearable
          type="password"
          class="login-input"
          @blur="validateField('password')"
        />
        <div class="error-msg" v-if="errors.password">
          {{ errors.password }}
        </div>
      </div>

      <!-- 注册按钮 -->
      <div class="login-actions">
        <nut-button type="primary" block @click="register" :loading="isLoading">
          注册
        </nut-button>
        <nut-button type="default" block @click="goLogin">
          去登录
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
const formData = reactive({
  email: "",
  code: "",
  password: ""
});
const errors = reactive({ email: "", code: "", password: "" });
const isLoading = ref(false);
const countdown = ref(0);
let timer = null;

// 校验规则
const rules = {
  email: (val) => {
    if (!val) return "请输入邮箱";
    const emailReg = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailReg.test(val)) return "邮箱格式不正确";
    return "";
  },
  code: (val) => {
    if (!val) return "请输入验证码";
    return "";
  },
  password: (val) => {
    if (!val) return "请输入密码";
    if (val.length <= 6) return "密码必须大于6位";
    return "";
  }
};

// 校验单项
const validateField = (key) => {
  errors[key] = rules[key](formData[key]);
  if (errors[key]) showNotify.warn(errors[key]);
};

// 发送验证码
const sendCode = async () => {
  errors.email = rules.email(formData.email);
  if (errors.email) return;

  try {
    const res = await axios.post("/auth/send-email-code", { email: formData.email });
    if (res.data.code === 0) {
      showNotify.success("验证码已发送");
      startCountdown();
    } else {
      showNotify.warn(res.data.message || "发送失败");
    }
  } catch (err) {
    console.error(err);
    showNotify.warn("网络错误，请稍后重试");
  }
};

// 倒计时
const startCountdown = () => {
  countdown.value = 60;
  timer = setInterval(() => {
    countdown.value--;
    if (countdown.value <= 0) clearInterval(timer);
  }, 1000);
};

// 注册提交
const register = async () => {
  errors.email = rules.email(formData.email);
  errors.code = rules.code(formData.code);
  errors.password = rules.password(formData.password);

  if (errors.email || errors.code || errors.password) return;

  isLoading.value = true;
  try {
    const res = await axios.post("/auth/register-email", formData);
    if (res.data.code === 0) {
      showNotify.success("注册成功");
      router.push("/login");
    } else {
      showNotify.warn(res.data.message || "注册失败");
    }
  } catch (err) {
    console.error(err);
    showNotify.warn("网络错误，请稍后重试");
  } finally {
    isLoading.value = false;
  }
};

// 去登录
const goLogin = () => router.push("/login");
</script>

<style scoped>
.register-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  padding-top: 4rem;
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
  margin-bottom: 0.8rem;
}

/* 副标题 */
.register-wrapper p {
  margin: 0;
  margin-bottom: 2rem;
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

/* 输入框去背景线 */
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

/* 验证码行 */
.code-item {
  display: flex;
  gap: 10px;
  align-items: center;
  max-width: 92%;
}

.send-code-btn {
  border-radius: 24px;
  font-size: 14px;
  padding: 0 16px;
  height: 42px;
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

/* 按钮样式 */
.nut-button {
  border-radius: 24px !important;
  font-size: 16px !important;
}
</style>
