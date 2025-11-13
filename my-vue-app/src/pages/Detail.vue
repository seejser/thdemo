<template>
  <div id="app">
    <nut-navbar title="设备详情" left-show @click-back="goBack"></nut-navbar>

    <div class="detail-content">
      <div class="card">
        <h2>基础信息</h2>
        <div class="info-item">
          <span class="label">设备号</span>
          <span class="value">{{ device.name }}</span>
        </div>
        <div class="info-item">
          <span class="label">所属产品</span>
          <span class="value">{{ device.product }}</span>
        </div>
        <div class="info-item">
          <span class="label">小区基站信息</span>
          <span class="value">{{ device.cell_info || '暂无' }}</span>
        </div>
        <div class="info-item">
          <span class="label">位置</span>
          <span class="value">{{ device.region || '暂无' }}</span>
        </div>
        <div class="info-item">
          <span class="label">SIM卡号</span>
          <span class="value">{{ device.sim_number }}</span>
        </div>
        <div class="info-item">
          <span class="label">MAC地址</span>
          <span class="value">{{ device.mac }}</span>
        </div>
        <div class="info-item">
          <span class="label">上报周期</span>
          <span class="value">{{ device.report_cycle }} 分钟/次</span>
        </div>
        <div class="info-item">
          <span class="label">设备状态</span>
          <span
            class="value"
            :style="{ color: device.status === 'online' ? '#4caf50' : '#f44336' }"
          >
            {{ device.status === 'online' ? '在线' : '离线' }}
          </span>
        </div>
      </div>

      <div class="card">
        <h2>实时状态</h2>
        <div class="status-grid">
          <div class="status-box">
            <div class="label">信号质量</div>
            <div class="phone-signal-bars">
              <div
                v-for="n in 5"
                :key="n"
                class="signal-bar-segment"
                :class="{'active': n <= activeSignalBars, [signalColorClass]: n <= activeSignalBars}"
              ></div>
            </div>
            <div class="value">{{ device.signal }}/32</div>
          </div>
          <div class="status-box">
            <div class="label">温度</div>
            <div class="value">{{ device.temp }}°C</div>
            <div class="label">-55~125°C</div>
          </div>
        </div>
      </div>

      <div class="card">
        <h2>设备控制</h2>
        <div class="control-item">
          <span class="label">控制继电器开关</span>
          <label class="toggle-switch">
            <input type="checkbox" :checked="device.relay === 1" @change="relayToggle" />
            <span class="slider"></span>
          </label>
        </div>
        <div class="control-item">
          <span class="label">控制OUT_J9开关</span>
          <label class="toggle-switch">
            <input type="checkbox" :checked="device.out_j9 === 1" @change="J9Toggle" />
            <span class="slider"></span>
          </label>
        </div>
      </div>

      <div class="card">
        <h2>事件</h2>
        <div
          v-for="(warning, index) in warnings"
          :key="index"
          class="warning-item"
          @click="viewWarningDetail(warning)"
        >
          <span class="text">{{ warning.name }}</span>
          <span class="text">{{ warning.alarm_summary }}</span>
          <span class="time">{{ warning.time }}</span>
          <span class="arrow">&gt;</span>
        </div>
        <div v-if="warnings.length === 0" class="info-item">
          <span class="label">暂无报警</span>
        </div>
      </div>
    </div>

    <div v-if="messageBox.visible" class="message-box">{{ messageBox.text }}</div>
    <div v-if="isLoading" class="loader-overlay">
      <div class="loader"><span>loading...</span></div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();

const device = ref({
  id: 10415,
  name: '866560088910415',
  status: 'online',
  product: 'Ay3w0oGD25',
  cell_info: 'XX花园基站-信号塔3',
  sim_number: '8986011912345678901',
  mac: 'AA:BB:CC:DD:EE:FF',
  report_cycle: 5,
  signal: 28,
  temp: 26,
  relay: 1,
  out_j9: 0,
});

const warnings = ref([
  { name: '预警事件', alarm_summary: '断电', time: '2023-10-20 14:30' },
  { name: '预警事件', alarm_summary: '断电', time: '2023-10-20 14:30' },
]);

const messageBox = ref({ visible: false, text: '', timer: null });
const isLoading = ref(false);

const activeSignalBars = computed(() => Math.ceil(device.value.signal / 32 * 5));
const signalColorClass = computed(() => {
  if (device.value.signal < 11) return 'signal-low-color';
  if (device.value.signal < 22) return 'signal-medium-color';
  return 'signal-high-color';
});

const goBack = () => {
  console.log('执行返回上一页操作...');
  router.back();
};

const relayToggle = () => {
  device.value.relay = device.value.relay === 1 ? 0 : 1;
};
const J9Toggle = () => {
  device.value.out_j9 = device.value.out_j9 === 1 ? 0 : 1;
};

const viewWarningDetail = (warning) => {
  console.log('查看预警事件详情:', warning);
};

onMounted(() => {
  console.log('Detail.vue mounted');
});
</script>

<style scoped>
/* 你提供的样式可以直接放在这里或在外部 css 文件中引用 */
</style>
