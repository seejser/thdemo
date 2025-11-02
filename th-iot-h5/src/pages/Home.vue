<template>
  <div :class="['home-container', themeClass]">
    <!-- 顶部固定 Navbar -->
    <nut-navbar title="设备列表" fixed class="navbar-fixed" />

    <!-- 中间滚动区域 -->
    <nut-pull-refresh
      v-model="refresh"
      @refresh="onPullRefresh"
      loosing-txt="松开吧"
      loading-txt="刷新中..."
      :complete-duration="1000"
    >
      <div class="content-container" ref="scrollContainer">
        <!-- 空列表提示 -->
        <div v-if="devices.length === 0" class="empty-list-message">
          <p>暂无设备数据</p>
          <p>请下拉刷新重试</p>
        </div>

        <!-- 设备列表 -->
        <div
          v-for="device in devices"
          :key="device.id"
          class="device-item"
          :class="{
            'status-online': device.status === 'online',
            'status-offline': device.status === 'offline',
            'status-fault': device.status === 'fault'
          }"
        >
          <div class="device-info" @click="getDeviceDetail(device.id)">
            <p>{{ device.name }}</p>
            <p>{{ device.region }}</p>
          </div>

          <div class="device-actions">
            <div class="device-status">
              <span class="status-dot"></span>
              <span class="status-text">{{ statusText(device.status) }}</span>
            </div>
            <label class="toggle-switch">
              <input
                type="checkbox"
                :checked="device.relay === 1"
                @change="relayToggle(device.id, device.relay)"
                :disabled="device.status !== 'online'"
              />
              <span class="slider"></span>
            </label>
          </div>
        </div>

        <!-- 无限加载 -->
        <nut-infinite-loading
          v-model="infinityValue"
          :has-more="hasMore"
          :container="scrollContainer"
          @load-more="loadMore"
        >
        </nut-infinite-loading>
      </div>
    </nut-pull-refresh>

    <!-- 底部 TabBar -->
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

// 主题
const theme = ref("light");
const themeClass = computed(() =>
  theme.value === "dark" ? "dark-theme" : "light-theme"
);

// 底部 Tab
const activeTab = ref(0);

// 模拟设备数据
const devices = ref([]);
const refresh = ref(false);

// 无限加载
const cycle = ref(0);
const infinityValue = ref(false);
const hasMore = ref(true);
const scrollContainer = ref(null);

// 模拟请求设备数据
const fetchDevices = (start = 0, count = 10) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      const data = [];
      for (let i = start; i < start + count; i++) {
        data.push({
          id: 10000 + i,
          name: `86656008891041${i}`,
          status: ["online", "offline", "fault"][i % 3],
          relay: i % 2,
          region: "上海市嘉定区胜辛北路1661号向东200米",
        });
      }
      resolve(data);
    }, 1000);
  });
};

// 下拉刷新
const onPullRefresh = async () => {
  refresh.value = true;
  cycle.value = 0;
  const data = await fetchDevices(0, 10);
  devices.value = data;
  hasMore.value = true;
  refresh.value = false;
};

// 无限加载
const loadMore = async () => {
  infinityValue.value = true;
  const start = devices.value.length;
  const newData = await fetchDevices(start, 10);
  devices.value.push(...newData);
  cycle.value++;
  if (cycle.value > 2) hasMore.value = false;
  infinityValue.value = false;
};

// 状态文字
const statusText = (status) => {
  if (status === "online") return "在线";
  if (status === "offline") return "离线";
  if (status === "fault") return "故障";
  return "";
};

// 继电器开关
const relayToggle = (id, relay) => {
  const device = devices.value.find((d) => d.id === id);
  if (!device || device.status !== "online") return;
  device.relay = relay === 1 ? 0 : 1;
};

// 设备详情跳转
const getDeviceDetail = (id) => {
  alert(`跳转到设备详情页，ID: ${id}`);
};

// 生命周期：组件挂载时加载初始设备
onMounted(async () => {
  scrollContainer.value = document.querySelector(".content-container");
  devices.value = await fetchDevices(0, 6);
});
</script>

<style scoped>
.home-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  color: var(--text-color);
}

/* Navbar */
.navbar-fixed {
  position: sticky;
  top: 0;
  z-index: 10;
}

/* 内容 */
.content-container {
  height: calc(100vh - 44px);
  overflow-y: auto;
  padding: 10px;
}

/* 设备列表样式 */
.device-item {
  background-color: #fff;
  border-radius: 10px;
  padding: 15px;
  margin-bottom: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}
.device-info p:first-child {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 5px;
}
.device-info p:last-child {
  font-size: 14px;
  color: #999;
}
.device-actions {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  min-width: 90px;
}
.device-status {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
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

/* 开关 */
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
.status-offline input:checked + .slider,
.status-fault input:checked + .slider {
  background-color: #ccc;
}

/* 空列表 */
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

/* 主题 */
.light-theme {
  --bg-color: #f7f7f7;
  --text-color: #333;
}
.dark-theme {
  --bg-color: #1e1e1e;
  --text-color: #fff;
}

/* TabBar */
.nut-tabbar {
  position: sticky;
  bottom: 0;
  z-index: 10;
}
.nut-tabbar-item--active {
  color: #1989fa;
}
</style>
