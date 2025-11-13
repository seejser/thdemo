<template>
  <div class="container">
    <!-- 顶部固定导航栏 -->
    <div class="header">
      <p class="title">设备列表</p>
      <span class="refresh-icon" @click="refreshFun">&#x21BB;</span>
    </div>

    <nut-infinite-loading
      v-model="infinityValue"
      :has-more="hasMore"
      @load-more="loadMore"
    >
      <nut-pull-refresh
        v-model="refresh"
        @refresh="refreshFun"
        loosing-txt="松开刷新"
        loading-txt="刷新中..."
        :complete-duration="1000"
      >
        <!-- 设备列表 -->
        <div class="device-list">
          <div
            v-for="(item, index) in sum"
            :key="index"
            class="device-item"
            :class="{
              'status-online': index % 3 === 0,
              'status-offline': index % 3 === 1,
              'status-fault': index % 3 === 2
            }"
            @click="goToDetail(index)"
          >
            <div class="device-info">
              <p>设备 {{ index }}</p>
              <p>区域 {{ index + 1 }}号</p>
            </div>

            <div class="device-actions">
              <div class="device-status">
                <span class="status-dot"></span>
                <span class="status-text">
                  {{ statusText(index % 3) }}
                </span>
              </div>

              <label class="toggle-switch">
                <input type="checkbox" :checked="index % 2 === 0" />
                <span class="slider"></span>
              </label>
            </div>
          </div>

          <!-- 空列表提示 -->
          <div v-if="sum === 0" class="empty-list-message">
            <p>暂无设备数据</p>
            <p>请点击刷新按钮或联系管理员添加设备</p>
          </div>
        </div>
      </nut-pull-refresh>
    </nut-infinite-loading>

    <HyTabBar />
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import HyTabBar from "./../components/Hytabbar.vue";

const router = useRouter();
const sum = ref(24);
const cycle = ref(0);
const infinityValue = ref(false);
const hasMore = ref(true);
const refresh = ref(false);

const loadMore = () => {
  setTimeout(() => {
    sum.value += 24;
    cycle.value++;
    if (cycle.value > 2) hasMore.value = false;
    infinityValue.value = false;
  }, 1000);
};

const refreshFun = () => {
  setTimeout(() => {
    refresh.value = false;
    sum.value = 24;
  }, 1500);
};

const goToDetail = (index) => {
  router.push("/detail/" + index);
};

const statusText = (statusIndex) => {
  if (statusIndex === 0) return "在线";
  if (statusIndex === 1) return "离线";
  if (statusIndex === 2) return "故障";
};
</script>

<style scoped>
/* 顶部固定导航栏 */
.header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 50px;
  background-color: #fff;
  border-bottom: 1px solid #eee;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10;
  font-weight: bold;
  font-size: 18px;
}
.header .title {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}
.header .refresh-icon {
  position: absolute;
  right: 15px;
  font-size: 22px;
  cursor: pointer;
}

/* 列表容器 */
.device-list {
  padding: 60px 10px 10px 10px; /* 留出顶部导航高度 */
}

/* 设备卡片 */
.device-item {
  background-color: #fff;
  border-radius: 10px;
  padding: 12px 15px;
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

/* 设备信息 */
.device-info {
  flex-grow: 1;
  margin-right: 10px;
}
.device-info p:first-child {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 4px;
}
.device-info p:last-child {
  font-size: 14px;
  color: #999;
}

/* 状态 + 开关 */
.device-actions {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  min-width: 90px;
}

.device-status {
  display: flex;
  align-items: center;
  margin-bottom: 6px;
  height: 24px;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 5px;
}
.status-online .status-dot {
  background-color: #2ecc71;
}
.status-offline .status-dot,
.status-fault .status-dot {
  background-color: #f44336;
}

.status-text {
  font-size: 14px;
  min-width: 35px;
  text-align: right;
}
.status-online .status-text {
  color: #2ecc71;
}
.status-offline .status-text,
.status-fault .status-text {
  color: #f44336;
}

/* 开关样式 */
.toggle-switch {
  position: relative;
  display: inline-block;
  width: 44px;
  height: 24px;
}
.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: 0.4s;
  border-radius: 34px;
}
.slider:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}
input:checked + .slider {
  background-color: #2ecc71;
}
input:checked + .slider:before {
  transform: translateX(20px);
}

/* 空列表提示 */
.empty-list-message {
  text-align: center;
  padding: 50px 20px;
  color: #999;
  font-size: 16px;
}
.empty-list-message p:first-child {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 10px;
}
</style>
