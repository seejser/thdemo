<template>
  <div class="error-page">
    <nut-empty image="error" description="页面出错了">
      <div class="btn-group">
        <nut-button type="primary" @click="goBack">刷新</nut-button>
        <nut-button plain type="info" @click="goHome">返回首页</nut-button>
      </div>
    </nut-empty>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  // 如果有历史记录则返回上一页，否则回首页
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push('/')
  }
}

const goHome = () => {
  router.push('/')
}
</script>

<style scoped>
.error-page {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.btn-group {
  display: flex;
  gap: 10px;
  margin-top: 12px;
  justify-content: center;
}
</style>
