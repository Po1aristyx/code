import request from '@/utils/request.js';

const cartApi = {
  getCart() {
    return request.get('/api/cart').then(res => res.data);
  },
  addToCart(data) {
    return request.post('/api/cart/add', data).then(res => res.data);
  },
  updateCart(data) {
    return request.post('/api/cart/update', data).then(res => res.data);
  },
  deleteCartItem(data) {
    return request.post('/api/cart/delete', data).then(res => res.data);
  }
};

export default cartApi; 