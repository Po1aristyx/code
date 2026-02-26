import request from '@/utils/request.js';

const productApi = {
  list() {
    return request.get('/api/products').then(res => res.data);
  },
  detail(id) {
    return request.get(`/api/product/${id}`).then(res => res.data);
  }
};

export default productApi; 