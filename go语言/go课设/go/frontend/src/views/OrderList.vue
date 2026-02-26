<template>
  <div class="order-list-container">
    <h1>我的订单</h1>
    <el-tabs v-model="activeTab" @tab-click="handleTabClick">
      <el-tab-pane label="全部订单" name="all"></el-tab-pane>
      <el-tab-pane label="待付款" name="pending_payment"></el-tab-pane>
      <el-tab-pane label="待发货" name="to_be_shipped"></el-tab-pane>
      <el-tab-pane label="待收货" name="in_transit"></el-tab-pane>
      <el-tab-pane label="已完成" name="completed"></el-tab-pane>
    </el-tabs>

    <div class="order-list">
      <div v-if="filteredOrders.length > 0">
        <el-card v-for="order in filteredOrders" :key="order.id" class="order-card" @click="handleViewDetail(order)">
          <template #header>
            <div class="card-header">
              <span>订单号：{{ order.id }}</span>
              <span :class="['order-status', 'status-' + order.status]">{{ formatStatus(order.status) }}</span>
            </div>
          </template>
          <div v-for="item in order.items" :key="item.id" class="order-item">
            <img :src="item.product.image" :alt="item.product.name" class="item-image" @error="e => e.target.src = '/default.png'">
            <div class="item-info">
              <span>{{ item.product.name }}</span>
              <span>x{{ item.quantity }}</span>
            </div>
            <div class="item-price">¥{{ item.price }}</div>
          </div>
          <div class="order-summary">
            <span>总计：<strong>¥{{ order.totalAmount }}</strong></span>
            <div class="order-actions">
              <el-button type="primary" plain size="small" @click="handleViewDetail(order)">查看详情</el-button>
              <el-button v-if="order.status === 'pending_payment' || order.status === '待支付'" type="success" size="small" @click.stop="handlePay(order)">去支付</el-button>
              <el-button v-if="order.status === 'in_transit' || order.status === '待收货'" type="warning" size="small" @click.stop="handleReceive(order)">确认收货</el-button>
              <el-button v-if="order.status === 'to_be_shipped' || order.status === '待发货'" type="info" size="small" @click.stop="handleShip(order)">发货</el-button>
            </div>
          </div>
        </el-card>
      </div>
      <el-empty v-else :description="'暂无' + formatStatus(activeTab, true) + '订单'"></el-empty>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import orderApi from '@/api/order.js';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';

const activeTab = ref('all');
const orders = ref([]);
const router = useRouter();

const fetchOrders = async () => {
  const res = await orderApi.getOrderList();
  if (res.code === 0) {
    orders.value = res.data;
  }
};

onMounted(fetchOrders);

const filteredOrders = computed(() => {
  if (activeTab.value === 'all') {
    return orders.value;
  }
  // 兼容中英文状态
  const tabMap = {
    pending_payment: ['pending_payment', '待支付'],
    to_be_shipped: ['to_be_shipped', '待发货'],
    in_transit: ['in_transit', '待收货'],
    completed: ['completed', '已完成'],
  };
  return orders.value.filter(order => tabMap[activeTab.value]?.includes(order.status));
});

const formatStatus = (status, forEmpty = false) => {
  const statusMap = {
    pending_payment: '待付款',
    to_be_shipped: '待发货',
    in_transit: '待收货',
    completed: '已完成',
    all: forEmpty ? '全部' : '',
    // 兼容后端老数据
    '待支付': '待付款',
    '待发货': '待发货',
    '待收货': '待收货',
    '已完成': '已完成',
  };
  return statusMap[status] || '';
};

const handleTabClick = (tab) => {
  // console.log(tab.props.name);
};

const handleViewDetail = (order) => {
  router.push(`/order/${order.id}`);
};

const handlePay = async (order) => {
  const res = await orderApi.updateOrderStatus(order.id, { status: 'to_be_shipped' });
  if (res.code === 0) {
    ElMessage.success('支付成功，等待发货');
    fetchOrders();
  } else {
    ElMessage.error(res.msg || '操作失败');
  }
};

const handleReceive = async (order) => {
  const res = await orderApi.updateOrderStatus(order.id, { status: 'completed' });
  if (res.code === 0) {
    ElMessage.success('收货成功，订单已完成');
    fetchOrders();
  } else {
    ElMessage.error(res.msg || '操作失败');
  }
};

const handleShip = async (order) => {
  const res = await orderApi.updateOrderStatus(order.id, { status: 'in_transit' });
  if (res.code === 0) {
    ElMessage.success('发货成功，等待收货');
    fetchOrders();
  } else {
    ElMessage.error(res.msg || '操作失败');
  }
};

</script>

<style scoped>
.order-list-container {
  max-width: 1200px;
  margin: 40px auto;
  padding: 20px;
}
.order-card {
  margin-bottom: 20px;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.order-status {
  font-weight: bold;
}
.status-completed { color: #67c23a; }
.status-in_transit { color: #e6a23c; }
.status-to_be_shipped { color: #409eff; }
.status-pending_payment { color: #f56c6c; }

.order-item {
  display: flex;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid #f2f2f2;
}
.order-item:last-child {
  border-bottom: none;
}
.item-image {
  width: 80px;
  height: 80px;
  object-fit: cover;
  margin-right: 20px;
}
.item-info {
  flex-grow: 1;
}
.item-price {
  width: 100px;
  text-align: right;
}
.order-summary {
  text-align: right;
  margin-top: 10px;
  padding-top: 10px;
  border-top: 1px solid #f2f2f2;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 10px;
}
.order-summary strong {
  font-size: 18px;
  color: #f56c6c;
}
</style> 