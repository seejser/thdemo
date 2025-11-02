<template>
  <div :class="['home-container', themeClass]">
    <!-- 顶部固定导航栏 -->
    <nut-navbar title="设备列表" class="navbar-fixed" fixed />

    <!-- 中间可滚动区域 -->
    <nut-pull-refresh
      v-model="refresh"
      @refresh="refreshFun"
      loosing-txt="松开吧"
      loading-txt="玩命刷新中..."
      :complete-duration="1000"
    >
      <div class="content-container" ref="scrollContainer">
        <nut-infinite-loading
          v-model="infinityValue"
          :has-more="hasMore"
          :container="scrollContainer"
          @load-more="loadMore"
        >
          <div class="list" v-for="(item, index) in sum" :key="index">
            {{ index }}
          </div>
        </nut-infinite-loading>
      </div>
    </nut-pull-refresh>

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
import { ref, computed, onMounted } from "vue";
import { Home as HomeIcon, My as MyIcon } from "@nutui/icons-vue";

// 主题切换
const theme = ref("light");
const themeClass = computed(() =>
  theme.value === "dark" ? "dark-theme" : "light-theme"
);

// 底部 Tab
const activeTab = ref(0);

// 无限加载相关
const cycle = ref(0);
const sum = ref(24);
const infinityValue = ref(false);
const hasMore = ref(true);
const scrollContainer = ref(null);

const loadMore = () => {
  setTimeout(() => {
    sum.value += 24;
    cycle.value++;
    if (cycle.value > 2) hasMore.value = false;
    infinityValue.value = false;
  }, 1000);
};

// 下拉刷新
const refresh = ref(false);
const refreshFun = () => {
  setTimeout(() => {
    sum.value = 24;
    cycle.value = 0;
    hasMore.value = true; // 下拉刷新时重置无限加载
    refresh.value = false;
  }, 1500);
};

// 获取滚动容器 DOM
onMounted(() => {
  scrollContainer.value = document.querySelector(".content-container");
});
</script>

<style scoped>
.home-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  color: var(--text-color);
}

/* 顶部固定 Navbar */
.navbar-fixed {
  position: sticky;
  top: 0;
  z-index: 10;
}

/* 中间内容可滚动 */
.content-container {
  height: calc(100vh - 110px); /* 顶部 Navbar + 底部 TabBar 高度 */
  overflow-y: auto;
}

/* 列表项 */
.list {
  padding: 12px 20px;
  border-top: 1px solid #eee;
}

/* 主题 */
.light-theme {
  --bg-color: #fff;
  --text-color: #333;
}

.dark-theme {
  --bg-color: #1e1e1e;
  --text-color: #fff;
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
</style>
