<template>
  <nut-config-provider :theme="theme">
    <div :class="['app-container', themeClass]">
      <!-- 页面内容 -->
      <main class="page-container">
        <router-view />
      </main>

      <!-- 底部 TabBar -->
      <nut-tabbar
        v-model="activeTab"
        bottom
        safe-area-inset-bottom
      >
        <nut-tabbar-item tab-title="首页">
          <template #icon>
            <Home />
          </template>
        </nut-tabbar-item>
        <nut-tabbar-item tab-title="我的">
          <template #icon>
            <My />
          </template>
        </nut-tabbar-item>
      </nut-tabbar>
    </div>
  </nut-config-provider>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { Home, My } from '@nutui/icons-vue'

const router = useRouter()
const route = useRoute()

// 主题
const theme = ref('light')
const themeClass = computed(() => (theme.value === 'dark' ? 'theme-dark' : 'theme-light'))

// Tab 选中索引
const activeTab = ref(0)

// 路由对应关系
const tabRoutes = ['/', '/me']

// 页面切换时同步 Tab 高亮
watch(
  () => route.path,
  (path) => {
    const index = tabRoutes.indexOf(path)
    if (index >= 0) activeTab.value = index
  },
  { immediate: true }
)

// Tab 索引变化时跳转路由
watch(
  activeTab,
  async (index) => {
    const path = tabRoutes[index]
    if (path && path !== route.path) {
      try {
        await router.push(path)
      } catch (err) {
        console.error('路由跳转失败:', err)
      }
    }
  }
)
</script>

<style scoped>
.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  color: var(--text-color);
}

.page-container {
  flex: 1;
  /* padding: 16px; */
  overflow-y: auto;
}

.theme-light {
  --bg-color: #fff;
  --text-color: #333;
  --active-color: #1989fa;
}

.theme-dark {
  --bg-color: #1e1e1e;
  --text-color: #eee;
  --active-color: #64b5f6;
}

.nut-tabbar-item--active {
  color: var(--active-color);
}
</style>
