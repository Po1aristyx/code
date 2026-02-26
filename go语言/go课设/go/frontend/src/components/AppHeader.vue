<template>
  <el-header class="app-header">
    <div class="logo" @click="goHome">
      电商平台
    </div>

    <!-- 使用 flex-grow: 1 来占据所有剩余空间 -->
    <el-menu mode="horizontal" :ellipsis="false" router class="header-menu">
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/products">商品列表</el-menu-item>
      <el-menu-item index="/orders">我的订单</el-menu-item>
      <el-menu-item v-if="isLogin" index="/user">个人中心</el-menu-item>
    </el-menu>

    <!-- 这个 div 会被 margin-left: auto 推到最右边 -->
    <div class="actions">
      <el-badge :value="cartStore.totalItemsCount" :hidden="cartStore.totalItemsCount === 0">
        <el-button :icon="ShoppingCart" circle @click="goToCart"></el-button>
      </el-badge>
      <template v-if="isLogin">
        <el-button @click="handleLogout">退出登录</el-button>
      </template>
      <template v-else>
        <el-button @click="goToLogin">登录/注册</el-button>
      </template>
    </div>
  </el-header>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useCartStore } from '@/store/cart.js';
import { ShoppingCart } from '@element-plus/icons-vue';
import { ref, onMounted } from 'vue';

const router = useRouter();
const cartStore = useCartStore();

const isLogin = ref(!!localStorage.getItem('token'));

window.addEventListener('storage', () => {
  isLogin.value = !!localStorage.getItem('token');
});

onMounted(() => {
  isLogin.value = !!localStorage.getItem('token');
});

const goHome = () => {
  router.push('/');
}

const goToCart = () => {
  router.push('/cart');
}

const goToLogin = () => {
  router.push('/login');
}

const handleLogout = () => {
  localStorage.removeItem('token');
  window.location.reload();
}
</script>

<style scoped>
.app-header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000;
  
  display: flex;
  align-items: center;
  padding: 0 20px;
  box-sizing: border-box;
  border-bottom: 1px solid #e6e6e6;
  background-color: #fff;
  height: 60px;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  cursor: pointer;
  margin-right: 40px;
}

.header-menu {
  border-bottom: none;
}

.actions {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-left: auto;
}
</style> 