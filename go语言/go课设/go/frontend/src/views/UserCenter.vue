<template>
  <div class="user-center-container">
    <h1>个人中心</h1>
    <el-tabs v-model="activeTab">
      <el-tab-pane label="购物车" name="cart">
        <el-table :data="cartItems" style="width: 100%">
          <el-table-column prop="product.name" label="商品名称"></el-table-column>
          <el-table-column prop="product.price" label="单价"></el-table-column>
          <el-table-column prop="quantity" label="数量"></el-table-column>
          <el-table-column label="小计">
            <template #default="scope">¥{{ scope.row.product.price * scope.row.quantity }}</template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="购买历史" name="orders">
        <el-table :data="orders" style="width: 100%">
          <el-table-column prop="id" label="订单号"></el-table-column>
          <el-table-column prop="totalAmount" label="总金额"></el-table-column>
          <el-table-column prop="status" label="状态"></el-table-column>
          <el-table-column label="下单时间">
            <template #default="scope">{{ scope.row.createdAt }}</template>
          </el-table-column>
          <el-table-column label="操作">
            <template #default="scope">
              <el-button type="primary" size="small" @click="goToOrderDetail(scope.row.id)">查看详情</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import cartApi from '@/api/cart.js';
import orderApi from '@/api/order.js';

const router = useRouter();
const activeTab = ref('cart');
const cartItems = ref([]);
const orders = ref([]);

const fetchCart = async () => {
  const res = await cartApi.getCart();
  if (res.code === 0) {
    cartItems.value = res.data;
  }
};

const fetchOrders = async () => {
  const res = await orderApi.getOrderList();
  if (res.code === 0) {
    orders.value = res.data;
  }
};

const goToOrderDetail = (id) => {
  router.push(`/order/${id}`);
};

onMounted(() => {
  fetchCart();
  fetchOrders();
});
</script>

<style scoped>
.user-center-container {
  max-width: 900px;
  margin: 40px auto;
  padding: 20px;
}
</style> 