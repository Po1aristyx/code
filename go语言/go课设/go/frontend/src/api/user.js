import request from '@/utils/request.js';

const userApi = {
  login(data) {
    return request.post('/api/login', data).then(res => res.data);
  },
  register(data) {
    return request.post('/api/users', data).then(res => res.data);
  }
};

export default userApi; 