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
      <!-- 编辑图标按钮，放在右上角 -->
      <nut-button
        class="edit-btn"
        shape="circle"
        size="small"
        type="primary"
        @click="editProfile"
      >
        <EditIcon />
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
      <div class="version">{{ version }}</div>
    </div>

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
import { ref, computed } from "vue";
import {
  Home as HomeIcon,
  My as MyIcon,
  Edit as EditIcon,
} from "@nutui/icons-vue";

// 主题切换
const theme = ref("light");
const themeClass = computed(() =>
  theme.value === "dark" ? "dark-theme" : "light-theme"
);

// 底部 Tab
const activeTab = ref(1); // 默认选中“我的”

// 用户信息
const avatarUrl = ref("https://i.pravatar.cc/150?img=3");
const username = ref("张三");

// 编辑图标按钮
const editProfile = () => {
  console.log("点击编辑用户信息");
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
  // 可在这里清理 token 并跳转登录页
};
</script>

<style scoped>
.me-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  /* background-color: var(--bg-color); */
  color: var(--text-color);
  width: 100%;
  height: 100%;
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
.edit-btn {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  padding: 4px;
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
  padding: 16px 20px;
  border-bottom: 1px solid #eee;
  cursor: pointer;
  transition: background 0.2s;
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
