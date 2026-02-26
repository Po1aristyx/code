import request from '@/utils/request.js';

const orderApi = {
  createOrder(data) {
    return request.post('/api/orders', data).then(res => res.data);
  },
  getOrderList() {
    return request.get('/api/orders').then(res => res.data);
  },
  getOrderDetail(id) {
    return request.get(`/api/orders/${id}`).then(res => res.data);
  },
  updateOrderStatus(id, data) {
    return request.put(`/api/orders/${id}/status`, data).then(res => res.data);
  }
};

export default orderApi; 