<template>
  <nut-tabbar
    v-model="active"
    @tab-switch="tabSwitch"
    bottom
    safe-area-inset-bottom
    placeholder
  >
    <nut-tabbar-item
      v-for="(item, index) in List"
      :key="index"
      :tab-title="item.title"
    >
      <template #icon>
        <component :is="item.icon" />
      </template>
    </nut-tabbar-item>
  </nut-tabbar>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { Home, My } from "@nutui/icons-vue";

const router = useRouter();
const active = ref(0);

// console.log("this.$router:", this.$router);

const List = [
  { title: "首页", icon: Home },
  { title: "我的", icon: My },
];

const tabSwitch = (item, index) => {
  console.log("item", item);
  console.log("切换到：", item.title, "索引：", index);
  if (index === 0) {
    router.push("/");
  } else if (index === 1) {
    router.push("/my");
  }
};
</script>
