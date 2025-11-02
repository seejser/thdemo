<template>
  <div :class="['me-container', themeClass]">
    <!-- 顶部用户信息 -->
    <div class="user-info">
      <nut-avatar size="64">
        <img :src="avatarUrl" alt="用户头像" />
      </nut-avatar>
      <div class="user-details">
        <div class="username">{{ username }}</div>
      </div>
      <nut-button
        class="set-btn"
        shape="circle"
        size="small"
        type="default"
        @click="openSettings"
      >
        <SettingIcon />
      </nut-button>
    </div>

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
    <!-- 底部固定 TabBar -->
    <nut-tabbar v-model="activeTab" bottom safe-area-inset-bottom>
      <nut-tabbar-item tab-title="首页">
        <template #icon>
          <HomeIcon />
        </template>
      </nut-tabbar-item>
      <nut-tabbar-item tab-title="我的">
        <template #icon>
          <MyIcon />
        </template>
      </nut-tabbar-item>
    </nut-tabbar>
  </div>
</template>

<script setup>
import { ref, computed, createVNode } from "vue";
import { useRouter } from "vue-router"; // 导入 router
import {
  Home as HomeIcon,
  My as MyIcon,
  Setting as SettingIcon,
} from "@nutui/icons-vue";
import { showDialog } from '@nutui/nutui'


// 主题切换
const theme = ref("light");
const themeClass = computed(() =>
  theme.value === "dark" ? "dark-theme" : "light-theme"
);
const router = useRouter(); // 获取 router 实例
// 底部 Tab
const activeTab = ref(1); // 默认选中“我的”

// 用户信息
const avatarUrl = ref("https://i.pravatar.cc/150?img=3");
const username = ref("张三");

// 编辑图标按钮
const openSettings = () => {
  console.log("点击设置按钮");
  // 可以跳转到设置页面或打开设置弹窗
};

// 菜单列表
const menuItems = ref([
  { text: "我的设备", action: () => console.log("点击我的设备") },
  { text: "我的消息", action: () => console.log("点击我的消息") },
  { text: "关于", action: () => console.log("点击关于") },
]);

// 软件版本
const version = ref("v1.0.0");

// 退出登录
const logout = () => {
  console.log("退出登录");
  showDialog({
    title: "退出登录",
    content: createVNode(
      "span",
      { style: { color: "red" } },
      "确定要退出登录吗？"
    ),
    onCancel,
    onOk,
  });
};
const onCancel = () => {
  console.log('event cancel')
}
const onOk = () => {
  console.log('event ok')
  router.push({ path: `/login` });
}
</script>

<style scoped>
.me-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  /* background-color: var(--bg-color); */
  color: var(--text-color);
  padding: 12px 16px;
}

/* 顶部用户信息 */
.user-info {
  display: flex;
  align-items: center;
  padding: 20px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  margin-bottom: 15px;
  position: relative;
}

.user-details {
  margin-left: 16px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.username {
  font-size: 18px;
  font-weight: bold;
}

/* 编辑按钮图标 */
.set-btn {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  width: 40px;
  height: 40px;
  padding: 0;
  line-height: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;
  border: none;
  box-shadow: none;
  background-color: rgba(0, 0, 0, 0.001); /* 背景弱化 */
  color: #666; /* 图标颜色弱化 */
  font-size: 18px;
}

/* 可选：调整图标大小 */
.set-btn svg {
  width: 20px;
  height: 20px;
  display: block; /* 防止 svg 继承行高 */
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

/* 软件版本和退出按钮 */
.footer-card {
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  padding: 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  margin-bottom: 15px;
}

.version {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: #888;
  font-size: 14px;
}

/* TabBar 固定样式 */
.nut-tabbar {
  position: sticky;
  bottom: 0;
  z-index: 10;
}

/* TabBar 高亮 */
.nut-tabbar-item--active {
  color: #1989fa;
}

/* 主题 */
.light-theme {
  --bg-color: #f0f2f5;
  --text-color: #333;
}

.dark-theme {
  --bg-color: #1e1e1e;
  --text-color: #fff;
}
</style>
