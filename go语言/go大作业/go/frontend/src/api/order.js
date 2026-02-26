import request from '@/utils/request.js';

const orderApi = {
  createOrder(data) {
    return request.post('/api/order/create', data).then(res => res.data);
  },
  getOrderList() {
    return request.get('/api/orders').then(res => res.data);
  },
  getOrderDetail(id) {
    return request.get(`/api/order/${id}`).then(res => res.data);
  },
  updateOrderStatus(data) {
    return request.post('/api/order/update_status', data).then(res => res.data);
  }
};

export default orderApi; 