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
            v-for="item in devices"
            :key="item.id"
            class="device-item"
            :class="{
              'status-online': item.status === 0,
              'status-offline': item.status === 1,
              'status-fault': item.status === 2
            }"
            @click="goToDetail(item.id)"
          >
            <div class="device-info">
              <p>{{ item.name }}</p>
              <p>产品 {{ item.product }}</p>
            </div>

            <div class="device-actions">
              <div class="device-status">
                <span class="status-dot"></span>
                <span class="status-text">{{ statusText(item.status) }}</span>
              </div>

              <label class="toggle-switch">
                <input type="checkbox" v-model="item.switch" @change="toggleSwitch(item)" />
                <span class="slider"></span>
              </label>
            </div>
          </div>

          <!-- 空列表提示 -->
          <div v-if="devices.length === 0" class="empty-list-message">
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
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import HyTabBar from "./../components/Hytabbar.vue";
import { getDeviceList } from "@/api/device";

const router = useRouter();
const devices = ref([]);
const page = ref(1);
const limit = ref(10);
const total = ref(0);
const hasMore = ref(true);
const infinityValue = ref(false);
const refresh = ref(false);

const fetchDevices = async (reset = false) => {
  try {
    const res = await getDeviceList(page.value, limit.value);

    if (res.data.code !== 0) return;

    const list = res.data.data.list || [];
    total.value = res.data.data.total;

    if (reset) {
      devices.value = list;
    } else {
      devices.value = [...devices.value, ...list];
    }

    hasMore.value = devices.value.length < total.value;
  } catch (err) {
    console.error("获取设备列表失败", err);
    hasMore.value = false;
  } finally {
    infinityValue.value = false;
    refresh.value = false;
  }
};


const refreshFun = () => {
  page.value = 1;
  fetchDevices(true);
};

const loadMore = () => {
  page.value++;
  fetchDevices();
};

const goToDetail = (id) => {
  router.push("/detail/" + id);
};

const statusText = (status) => {
  switch (status) {
    case 0: return "在线";
    case 1: return "离线";
    case 2: return "故障";
    default: return "未知";
  }
};

const toggleSwitch = (item) => {
  item.switch = !item.switch;
  console.log("设备开关状态:", item.name, item.switch);
  // TODO: 调用后端接口更新开关状态
};

onMounted(() => {
  fetchDevices(true);
});

</script>

<style scoped>
/* 保留你原有的样式，不做改动 */
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

.device-list {
  padding: 60px 10px 10px 10px;
}

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
