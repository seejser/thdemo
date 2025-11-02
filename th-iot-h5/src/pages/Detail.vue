<template>
  <div class="detail-container">
    <!-- 顶部导航栏 -->
    <nut-navbar
      title="设备详情"
      fixed
      class="navbar-fixed"
      left-show
      @click-back="onClickBack"
    ></nut-navbar>

    <!-- 下拉刷新区域 -->
    <nut-pull-refresh
      v-model="refreshing"
      @refresh="onRefresh"
      loosing-txt="松开刷新"
      loading-txt="刷新中..."
      class="content"
    >
      <!-- 自定义下拉提示 -->
      <template #pulling-txt>
        <div class="pulling-text">用力拉，刷新数据吧！</div>
      </template>

      <!-- 内容区域 -->
      <div class="content-inner">
        <p>向下拉试试吧！</p>
        <p v-for="i in 20" :key="i">更多内容行 {{ i }}</p>
      </div>
    </nut-pull-refresh>
  </div>
</template>

<script>
import { showToast } from '@nutui/nutui'

export default {
  name: "Detail",
  data() {
    return {
      refreshing: false,
    };
  },
  methods: {
    onClickBack() {
      console.log("[Navbar]: 返回按钮点击");
      this.$router.back();
    },
    async onRefresh() {
      console.log("[PullRefresh]: 开始刷新");

      // 模拟请求数据
      await new Promise((resolve) => setTimeout(resolve, 1500));

      this.refreshing = false; // 结束刷新
      showToast.text("刷新完成", 1500);
      console.log("[PullRefresh]: 刷新完成");
    },
  },
};
</script>

<style scoped>
.detail-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.navbar-fixed {
  z-index: 10;
}

.content {
  flex: 1;
  margin-top: 2px; /* 保留导航栏空间 */
  overflow-y: auto;
  background-color: #f5f5f5;
}

.pulling-text {
  text-align: center;
  color: #888;
  padding: 10px 0;
}

.content-inner {
  padding: 16px;
}
</style>
