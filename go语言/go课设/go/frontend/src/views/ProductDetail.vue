<template>
  <div class="product-detail-overlay" @click.self="goBack">
    <el-card class="product-detail-card">
      <div v-if="product">
        <el-row :gutter="20">
          <el-col :span="12">
            <img :src="product.image" class="product-image" alt="商品大图" />
          </el-col>
          <el-col :span="12">
            <div class="product-info">
              <h1 class="product-name">{{ product.name }}</h1>
              <p class="product-description">{{ product.description }}</p>
              <div class="product-price">
                <span>价格：</span>
                <span class="price-value">¥{{ product.price }}</span>
              </div>
              <div class="product-quantity">
                <span>数量：</span>
                <el-input-number v-model="quantity" :min="1" :max="99" />
              </div>
              <div class="actions">
                <el-button type="primary" size="large" @click="handleBuyNow">立即购买</el-button>
                <el-button type="danger" size="large" @click="handleAddToCart">加入购物车</el-button>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
      <div v-else>
        <p>加载商品信息中...</p>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useCartStore } from '@/store/cart.js';
import productApi from '@/api/product.js';
import orderApi from '@/api/order.js';
import { ElMessage } from 'element-plus';

const route = useRoute();
const router = useRouter(); // 引入 useRouter
const cartStore = useCartStore();
const product = ref(null);
const quantity = ref(1);

// 处理加入购物车的逻辑
const handleAddToCart = () => {
  if (product.value) {
    cartStore.addItem(product.value, quantity.value);
  }
};

// 返回商品列表页
const goBack = () => {
  router.push('/products');
};

const handleBuyNow = async () => {
  if (product.value) {
    const res = await orderApi.createOrder({ address: '默认地址' });
    if (res.code === 0) {
      ElMessage.success('下单成功');
      router.push('/orders');
    } else {
      ElMessage.error(res.msg || '下单失败');
    }
  }
};

onMounted(async () => {
  const productId = Number(route.params.id);
  const res = await productApi.detail(productId);
  if (res.code === 0) {
    product.value = res.data;
  }
});
</script>

<style scoped>
.product-detail-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2000;
}

.product-detail-card {
  width: 80%;
  max-width: 900px;
}

.product-image {
  width: 100%;
  border-radius: 4px;
}
.product-info {
  padding-left: 20px;
}
.product-name {
  font-size: 28px;
  margin-bottom: 20px;
}
.product-description {
  color: #606266;
  margin-bottom: 20px;
}
.product-price {
  font-size: 18px;
  color: #f56c6c;
  margin-bottom: 20px;
}
.price-value {
  font-size: 24px;
  font-weight: bold;
}
.product-quantity {
  margin-bottom: 30px;
  display: flex;
  align-items: center;
  gap: 10px;
}
.actions {
  display: flex;
  gap: 15px;
}
</style>