<template>
  <div class="container">
    <div class="top-container">
      <nut-space direction="vertical" fill align="center">
        <nut-avatar size="large"> 欢迎 </nut-avatar>
        <p>请注册</p>
      </nut-space>
    </div>

    <nut-row>
      <!-- 用户名 -->
      <nut-col :span="24">
        <nut-input v-model="username" placeholder="请输入您的用户名" />
      </nut-col>

      <!-- 邮箱 -->
      <nut-col :span="24">
        <nut-input v-model="email" placeholder="请输入您的邮箱" type="email" />
      </nut-col>

      <!-- 图形验证码 -->
      <nut-col :span="24">
        <nut-row>
          <nut-col :span="12">
            <nut-input
              v-model="calcCode"
              placeholder="请输入图片验证码"
              clearable
            />
          </nut-col>
          <nut-col :span="8" class="captcha-box">
            <nut-image
              :src="captchaImage"
              width="100"
              height="34"
              @click="loadCaptcha"
            />
          </nut-col>
        </nut-row>
      </nut-col>

      <!-- 邮箱验证码 -->
      <nut-col :span="24">
        <nut-row>
          <nut-col :span="12">
            <nut-input
              v-model="verifyCode"
              placeholder="请输入邮箱验证码"
              clearable
            />
          </nut-col>
          <nut-col :span="8">
            <nut-button
              plain
              type="info"
              :disabled="countdown > 0"
              @click="sendEmailVerifyCode"
            >
              {{ countdown > 0 ? countdown + 's' : '发送验证码' }}
            </nut-button>
          </nut-col>
        </nut-row>
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

      <!-- 确认密码 -->
      <nut-col :span="24">
        <nut-input
          v-model="varpassword"
          placeholder="请确认您的密码"
          type="password"
          clearable
        />
      </nut-col>

      <!-- 操作按钮 -->
      <nut-col :span="24">
        <nut-space direction="vertical" fill>
          <nut-button type="primary" @click="doRegister" class="btn">
            注册
          </nut-button>
          <nut-button type="default" @click="goLogin" class="btn">
            登录
          </nut-button>
        </nut-space>
      </nut-col>
    </nut-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { showNotify } from "@nutui/nutui";
import { getCaptcha, sendEmailCode, register as apiRegister } from "@/api/auth";

const router = useRouter();

// 表单字段
const username = ref("");
const email = ref("");
const calcCode = ref("");
const verifyCode = ref("");
const password = ref("");
const varpassword = ref("");

// 验证码图片 & id
const captchaImage = ref("");
const captchaId = ref("");

// 邮箱验证码倒计时
const countdown = ref(0);
let timer = null;

// 获取验证码
const loadCaptcha = async () => {
  try {
    const res = await getCaptcha();
    if (res.data.code === 0) {
      captchaImage.value = res.data.data.image;
      captchaId.value = res.data.data.captcha_id;
    } else {
      showNotify.error(res.data.msg || "获取验证码失败");
    }
  } catch (err) {
    console.error(err);
    showNotify.error("获取验证码异常");
  }
};

// 发送邮箱验证码
const sendEmailVerifyCode = async () => {
  if (!email.value) return showNotify.warn("请先输入邮箱");
  if (!calcCode.value) return showNotify.warn("请输入图片验证码");
  if (!captchaId.value) return showNotify.warn("请刷新验证码");

  try {
    const res = await sendEmailCode({
      email,
      captcha: calcCode.value,
      captcha_id: captchaId.value,
    });
    if (res.data.code === 0) {
      showNotify.success("验证码已发送");
      startCountdown();
    } else {
      showNotify.error(res.data.msg || "发送验证码失败");
      loadCaptcha();
    }
  } catch (err) {
    console.error(err);
    showNotify.error("发送验证码异常");
    loadCaptcha();
  }
};

// 倒计时逻辑
const startCountdown = () => {
  countdown.value = 60;
  timer = setInterval(() => {
    countdown.value--;
    if (countdown.value <= 0) {
      clearInterval(timer);
    }
  }, 1000);
};

// 注册
const doRegister = async () => {
  if (!username.value || !password.value) return showNotify.warn("用户名或密码不能为空");
  if (password.value !== varpassword.value) return showNotify.warn("两次密码输入不一致");
  if (!email.value) return showNotify.warn("请输入邮箱");
  if (!verifyCode.value) return showNotify.warn("请输入邮箱验证码");
  if (!captchaId.value) return showNotify.warn("请刷新验证码");

  try {
    const res = await apiRegister({
      username: username.value,
      password: password.value,
      email: email.value,
      verifyCode: verifyCode.value,
      captcha_id: captchaId.value,
    });
    if (res.data.code === 0) {
      showNotify.success("注册成功");
      router.push("/login");
    } else {
      showNotify.error(res.data.msg || "注册失败");
      loadCaptcha();
    }
  } catch (err) {
    console.error(err);
    showNotify.error("注册异常");
    loadCaptcha();
  }
};

// 跳转登录页
const goLogin = () => router.push("/login");

// 页面挂载时获取验证码
onMounted(() => {
  loadCaptcha();
});
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

.captcha-box {
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn {
  width: 100% !important;
}
</style>
