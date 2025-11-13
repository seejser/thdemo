<template>
  <div class="container">
    <!-- 用户头像 -->
    <nut-space direction="vertical" fill align="center">
      <nut-avatar size="large">
        <img :src="avatarUrl" alt="用户头像" />
      </nut-avatar>
      <h1>{{ username }}</h1>
      <p>{{ welcomeMessage }}</p>
    </nut-space>

    <!-- 主内容，可扩展 -->
    <div class="main-content">
      <p>这里是用户的主页内容，可以放一些信息或卡片。</p>
    </div>

    <!-- 底部导航 -->
    <HyTabBar />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import HyTabBar from "@/components/Hytabbar.vue";
import { showNotify } from "@nutui/nutui";
import { getUserInfo } from "@/api/user"; 

// 用户信息
const username = ref("春天的羊");
const avatarUrl = ref(
  "https://img12.360buyimg.com/imagetools/jfs/t1/196430/38/8105/14329/60c806a4Ed506298a/e6de9fb7b8490f38.png"
);
const welcomeMessage = ref("欢迎回来！");

// 页面加载时可以请求用户信息
onMounted(async () => {
  try {
    const res = await api.get("/user/info");
    if (res.data && res.data.code === 0 && res.data.data) {
      username.value = res.data.data.username || username.value;
      avatarUrl.value = res.data.data.avatar || avatarUrl.value;
      welcomeMessage.value = `欢迎回来，${username.value}！`;
    } else {
      showNotify.warn(res.data.msg || "获取用户信息失败");
    }
  } catch (err) {
    console.error(err);
    showNotify.error("获取用户信息异常");
  }
});
</script>

<style scoped>
.container {
  padding: 16px;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  justify-content: space-between;
}

.main-content {
  flex: 1;
  padding: 16px;
  text-align: center;
}

h1 {
  font-size: 20px;
  margin: 8px 0;
}

p {
  font-size: 14px;
  color: #666;
  margin: 0;
}
</style>
