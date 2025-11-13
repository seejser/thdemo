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

    <!-- 功能菜单列表 -->
    <div class="menu-list">
      <div
        class="menu-item"
        v-for="item in menuItems"
        :key="item.text"
        @click="item.action"
      >
        <span>{{ item.text }}</span>
        <span class="arrow">></span>
      </div>
    </div>

    <!-- 软件版本和退出按钮 -->
    <div class="footer-card">
      <nut-button type="danger" block @click="logout">退出登录</nut-button>
    </div>
    <div class="version">{{ version }}</div>

    <!-- 底部导航 -->
    <HyTabBar />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import HyTabBar from "@/components/Hytabbar.vue";
import { showNotify } from "@nutui/nutui";
import { getUserInfo, goLogout } from "@/api/user";

// 用户信息
const username = ref("春天的羊");
const version = ref("V0.1.1");
// 菜单列表
const menuItems = ref([
  { text: "我的设备", action: () => console.log("点击我的设备") },
  { text: "我的消息", action: () => console.log("点击我的消息") },
  { text: "关于", action: () => console.log("点击关于") },
]);
const avatarUrl = ref(
  "https://img12.360buyimg.com/imagetools/jfs/t1/196430/38/8105/14329/60c806a4Ed506298a/e6de9fb7b8490f38.png"
);
const welcomeMessage = ref("欢迎回来！");
// 退出登录
function logout() {
  goLogout();
}
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
.version {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: #888;
  font-size: 14px;
}
p {
  font-size: 14px;
  color: #666;
  margin: 0;
}
/* 功能菜单列表 */
.menu-list {
  background-color: #fff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  margin-bottom: 15px;
}

.menu-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 20px; /* 略微紧凑 */
  border-bottom: 1px solid #f0f0f0; /* 更柔和的边线 */
  /* margin-bottom: 12px; 项目间距 */
  border-radius: 8px; /* 圆润角 */
  cursor: pointer;
  transition: background 0.3s, transform 0.2s;
  background-color: #fff; /* 白色背景更干净 */
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05); /* 轻微阴影 */
}

.menu-item:last-child {
  border-bottom: none;
}

.menu-item:hover {
  background-color: #f5f5f5;
}

.menu-item .arrow {
  color: #ccc;
  font-weight: bold;
}
</style>
