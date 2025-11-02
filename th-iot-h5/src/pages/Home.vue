<template>
  <div id="Home">
    <nut-navbar title="设备列表"></nut-navbar>
    <nut-infinite-loading
      v-model="infinityValue"
      :has-more="hasMore"
      @load-more="loadMore"
    >
      <nut-pull-refresh
        v-model="refresh"
        @refresh="refreshFun"
        loosing-txt="松开吧"
        loading-txt="玩命刷新中..."
        :complete-duration="1000"
      >
        <div class="test" v-for="(item, index) in sum" :key="index">
          {{ index }}
        </div>
      </nut-pull-refresh>
    </nut-infinite-loading>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";

const theme = ref("");
const checked = ref(false);
const themeClass = computed(() =>
  theme.value === "dark" ? "dark-theme" : "light-theme"
);

const cycle = ref(0);
const tabsValue = ref(0);
const sum = ref(24);
const infinityValue = ref(false);
const hasMore = ref(true);
const loadMore = () => {
  setTimeout(() => {
    sum.value += 24;
    cycle.value++;
    if (cycle.value > 2) hasMore.value = false;
    infinityValue.value = false;
  }, 1000);
};
const refresh = ref(false);
const refreshFun = () => {
  setTimeout(() => {
    refresh.value = false;
    sum.value = 24;
  }, 3000);
};
const change = (v) => {
  theme.value = v ? "dark" : "light";
};
</script>

<style>
.test {
  padding: 12px 0 12px 20px;
  border-top: 1px solid #eee;
}

.dark-theme .test {
  background-color: #333;
  color: #fff;
}

.light-theme .test {
  background-color: #fff;
  color: #333;
}
</style>
