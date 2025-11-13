<template>
  <div class="container">
    <nut-navbar title="设备详情" left-show @click-back="onClick"></nut-navbar>

    <div v-if="device" class="device-detail">
      <h2>设备号：{{ device.name }}</h2>
      <p>产品：{{ device.product }}</p>
      <p>状态：{{ statusText(device.status) }}</p>
      <p>开关状态：{{ device.switch ? "开" : "关" }}</p>
      <p>继电器：{{ device.relay ? "开" : "关" }}</p>
      <p>输出 J9：{{ device.out_j9 ? "开" : "关" }}</p>
      <p>创建时间：{{ device.created_at }}</p>
      <p>更新时间：{{ device.updated_at }}</p>
    </div>

    <div v-else class="loading">
      加载中...
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import { getDeviceById } from "@/api/device";

const router = useRouter();
const route = useRoute();
const device = ref(null);

const onClick = () => {
  router.push("/");
};

const statusText = (status) => {
  switch (status) {
    case 0: return "在线";
    case 1: return "离线";
    case 2: return "故障";
    default: return "未知";
  }
};

const fetchDevice = async () => {
  const id = route.params.id;
  try {
    const res = await getDeviceById(id);
    if (res.data.code === 0) {
      device.value = res.data.data;
    } else {
      console.error("获取设备详情失败:", res.data.msg);
    }
  } catch (err) {
    console.error("请求设备详情出错:", err);
  }
};

onMounted(() => {
  fetchDevice();
});
</script>

<style scoped>
.container {
  padding: 60px 15px 15px 15px; /* 顶部导航高度留白 */
}

.device-detail h2 {
  margin-bottom: 10px;
  font-size: 20px;
  font-weight: bold;
}

.device-detail p {
  margin-bottom: 6px;
  font-size: 16px;
}

.loading {
  text-align: center;
  margin-top: 50px;
  color: #999;
}
</style>
