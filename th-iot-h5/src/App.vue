<template>
  <nut-config-provider :theme="theme">
    <div :class="['app-container', themeClass]">
      <!-- 顶部导航 -->
      <nav class="nav-bar">
        <router-link to="/" class="nav-item" exact-active-class="active">首页</router-link>
        <router-link to="/detail" class="nav-item" exact-active-class="active">详情</router-link>
      </nav>

      <!-- 页面内容区域 -->
      <main class="page-container">
        <router-view />
      </main>
    </div>
  </nut-config-provider>
</template>

<script setup>
import { ref, computed } from 'vue'

// 当前主题，可切换为 'light' 或 'dark'
const theme = ref('light')

// 根据主题切换对应样式类
const themeClass = computed(() => (theme.value === 'dark' ? 'theme-dark' : 'theme-light'))
</script>

<style scoped>
/* ======= 全局布局 ======= */
.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  font-family: system-ui, Avenir, Helvetica, Arial, sans-serif;
  background-color: var(--bg-color);
  color: var(--text-color);
}

/* ======= 导航栏 ======= */
.nav-bar {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 12px 0;
  border-bottom: 1px solid #eee;
  background-color: var(--nav-bg);
}

.nav-item {
  color: var(--link-color);
  text-decoration: none;
  margin: 0 12px;
  font-size: 16px;
  transition: color 0.2s;
}

.nav-item.active {
  color: var(--active-color);
  font-weight: bold;
}

/* ======= 页面内容 ======= */
.page-container {
  flex: 1;
  padding: 16px;
}

/* ======= 主题样式 ======= */
.theme-light {
  --bg-color: #fff;
  --text-color: #333;
  --link-color: #666;
  --active-color: #1989fa;
  --nav-bg: #fafafa;
}

.theme-dark {
  --bg-color: #1e1e1e;
  --text-color: #eee;
  --link-color: #aaa;
  --active-color: #64b5f6;
  --nav-bg: #2c2c2c;
}
</style>
