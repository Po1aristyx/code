<template>
  <div class="product-list-container">
    <el-row :gutter="20">
      <el-col
        v-for="product in products"
        :key="product.id"
        :xs="24" :sm="12" :md="8" :lg="6"
      >
        <ProductCard :product="product" class="product-item" @add-to-cart="handleAddToCart" @buy-now="handleBuyNow"/>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import ProductCard from '@/components/ProductCard.vue';
import productApi from '@/api/product.js';
import { useCartStore } from '@/store/cart.js';
import orderApi from '@/api/order.js';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';

const products = ref([]);
const cartStore = useCartStore();
const router = useRouter();

onMounted(async () => {
  const res = await productApi.list();
  if (res.code === 0) {
    products.value = res.data;
  }
});

const handleAddToCart = (product) => {
  cartStore.addItem(product, 1);
};

const handleBuyNow = async (product) => {
  const res = await orderApi.createOrder({ address: '默认地址', productId: product.id, quantity: 1 });
  if (res.code === 0) {
    ElMessage.success('下单成功');
    router.push('/orders');
  } else {
    ElMessage.error(res.msg || '下单失败');
  }
};
</script>

<style scoped>
.product-list-container {
  padding: 20px;
}
.product-item {
  margin-bottom: 20px;
}
</style>
