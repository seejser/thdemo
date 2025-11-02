<template>
  <nut-config-provider :theme="theme">
    <div :class="['app-container', themeClass]">
      <main class="page-container">
        <router-view />
      </main>

      <!-- 只有一级页面显示 TabBar -->
      <nut-tabbar v-if="showTabBar" v-model="activeTab" bottom safe-area-inset-bottom>
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
import { ref, computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { Home, My } from "@nutui/icons-vue";

const route = useRoute();
const router = useRouter();

// 主题
const theme = ref("light");
const themeClass = computed(() => (theme.value === "dark" ? "theme-dark" : "theme-light"));

// Tab
const activeTab = ref(0);
const tabRoutes = ["/", "/me"];
const showTabBar = ref(true);

// 页面切换
watch(
  () => route.path,
  (path) => {
    showTabBar.value = route.meta.showTabBar !== false; // 默认显示 TabBar，meta=false 不显示
    const index = tabRoutes.indexOf(path);
    if (index >= 0) activeTab.value = index;
  },
  { immediate: true }
);

// Tab 切换跳路由
watch(activeTab, (index) => {
  const path = tabRoutes[index];
  if (path && path !== route.path) router.push(path);
});
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
