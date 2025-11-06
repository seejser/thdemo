<template>
  <div class="container">
    <!-- <nut-cell title="Dark">
        <template #link>
          <nut-switch v-model="checked" @change="change" />
        </template>
      </nut-cell> -->

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
        <div
          class="test"
          v-for="(item, index) in sum"
          :key="index"
          @click="goToDetail(index)"
        >
          {{ index }}
        </div>
      </nut-pull-refresh>
    </nut-infinite-loading>
    <HyTabBar />
    <!-- <nut-tabbar v-model="active" bottom safe-area-inset-bottom placeholder>
        <nut-tabbar-item
          v-for="(item, index) in List"
          :key="index"
          :tab-title="item.title"
          :icon="item.icon"
        >
        </nut-tabbar-item>
      </nut-tabbar> -->
  </div>
</template>

<script setup>
import { h, ref } from "vue";
import { useRouter } from "vue-router";

import HyTabBar from "./../components/Hytabbar.vue";

const cycle = ref(0);
const router = useRouter();
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
const goToDetail = (index) => {
  console.log("index:", index);
   router.push("/detail");
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
