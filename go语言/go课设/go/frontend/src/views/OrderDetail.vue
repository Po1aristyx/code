<template>
  <div class="order-detail-container" v-if="order">
    <h1>订单详情</h1>
    <el-card>
      <div>订单号：{{ order.id }}</div>
      <div>收货地址：{{ order.address }}</div>
      <div>状态：{{ order.status }}</div>
      <div>总金额：¥{{ order.totalAmount }}</div>
      <div>
        <h3>商品列表</h3>
        <el-table :data="order.items" style="width: 100%">
          <el-table-column prop="product.name" label="商品名称"></el-table-column>
          <el-table-column prop="product.price" label="单价"></el-table-column>
          <el-table-column prop="quantity" label="数量"></el-table-column>
          <el-table-column prop="price" label="小计">
            <template #default="scope">¥{{ scope.row.price * scope.row.quantity }}</template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
  <el-empty v-else description="未找到订单"></el-empty>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import orderApi from '@/api/order.js';

const route = useRoute();
const order = ref(null);

onMounted(async () => {
  const id = route.params.id;
  const res = await orderApi.getOrderDetail(id);
  if (res.code === 0) {
    order.value = res.data;
  }
});
</script>

<style scoped>
.order-detail-container {
  max-width: 800px;
  margin: 40px auto;
  padding: 20px;
}
</style> 