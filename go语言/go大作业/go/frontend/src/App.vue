<template>
  <el-container class="app-container">
    <AppHeader />
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script setup>
import AppHeader from '@/components/AppHeader.vue';
</script>

<style>
/* 全局样式重置 */
body {
  margin: 0;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
}

.app-container {
  min-height: 100vh;
}

.el-main {
  padding: 0;
  margin-top: 60px; /* <-- 关键：为固定的头部留出空间 */
}
</style>
 