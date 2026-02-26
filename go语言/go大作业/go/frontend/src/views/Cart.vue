<template>
  <div class="cart-container">
    <h1>我的购物车</h1>
    <el-card v-if="cartStore.items.length > 0">
      <el-table :data="cartStore.items" style="width: 100%">
        <el-table-column label="商品图片" width="120">
          <template #default="scope">
            <img :src="scope.row.product.image" alt="商品图片" class="product-image"/>
          </template>
        </el-table-column>
        <el-table-column prop="product.name" label="商品名称"></el-table-column>
        <el-table-column prop="product.price" label="单价" width="120">
          <template #default="scope">¥{{ scope.row.product.price }}</template>
        </el-table-column>
        <el-table-column label="数量" width="180">
          <template #default="scope">
            <el-input-number 
              :model-value="scope.row.quantity"
              @change="(val) => cartStore.updateItemQuantity(scope.row.product.id, val)"
              :min="1" 
              :max="99" 
              size="small"
            />
          </template>
        </el-table-column>
        <el-table-column label="小计" width="120">
          <template #default="scope">¥{{ scope.row.product.price * scope.row.quantity }}</template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="danger" size="small" @click="cartStore.removeItem(scope.row.product.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="cart-summary">
        <span class="total-price">总计：¥{{ cartStore.totalPrice }}</span>
        <el-button type="success" size="large" @click="handleCheckout">去结算</el-button>
      </div>
    </el-card>
    <el-empty v-else description="购物车是空的，快去逛逛吧！">
      <el-button type="primary" @click="$router.push('/products')">去逛逛</el-button>
    </el-empty>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useCartStore } from '@/store/cart.js';
import cartApi from '@/api/cart.js';
import orderApi from '@/api/order.js';
import { ElMessage } from 'element-plus';

const cartStore = useCartStore();
const cartItems = ref([]);
const totalPrice = ref(0);

const fetchCart = async () => {
  const res = await cartApi.getCart();
  if (res.code === 0) {
    cartItems.value = res.data;
    totalPrice.value = cartItems.value.reduce((sum, item) => sum + item.product.price * item.quantity, 0);
  }
};

const handleCheckout = async () => {
  // 简化：直接创建订单，实际可弹窗填写地址
  const res = await orderApi.createOrder({ address: '默认地址' });
  if (res.code === 0) {
    ElMessage.success('下单成功');
    fetchCart();
  } else {
    ElMessage.error(res.msg || '下单失败');
  }
};

onMounted(fetchCart);
</script>

<style scoped>
.cart-container {
  max-width: 1200px;
  margin: 40px auto;
  padding: 20px;
}
h1 {
  margin-bottom: 20px;
}
.product-image {
  width: 80px;
  height: 80px;
  object-fit: cover;
}
.cart-summary {
  margin-top: 20px;
  text-align: right;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 20px;
}
.total-price {
  font-size: 22px;
  color: #f56c6c;
  font-weight: bold;
}
</style>
