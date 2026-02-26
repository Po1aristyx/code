<template>
  <div>
    <h1>欢迎来到电商平台</h1>
    <div class="nav-links">
      <el-button @click="goTo('/login')">去登录</el-button>
      <el-button @click="goTo('/register')">去注册</el-button>
      <el-button type="primary" @click="goTo('/products')">查看商品</el-button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goTo = (path) => {
  router.push(path)
}
</script>

<style scoped>
div {
  padding-top: 50px;
  text-align: center;
}

.nav-links {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  gap: 10px;
}
</style> 