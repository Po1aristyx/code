import request from '@/utils/request.js';

const userApi = {
  login(data) {
    return request.post('/api/user/login', data).then(res => res.data);
  },
  register(data) {
    return request.post('/api/user/register', data).then(res => res.data);
  }
};

export default userApi; 