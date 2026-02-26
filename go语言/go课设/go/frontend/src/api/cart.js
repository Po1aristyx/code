import request from '@/utils/request.js';

const cartApi = {
  getCart() {
    return request.get('/api/cart').then(res => res.data);
  },
  addToCart(data) {
    return request.post('/api/cart', data).then(res => res.data);
  },
  updateCart(data) {
    return request.put('/api/cart', data).then(res => res.data);
  },
  deleteCartItem(data) {
    return request.delete('/api/cart', { data }).then(res => res.data);
  }
};

export default cartApi; 