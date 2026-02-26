import axios from 'axios'

// 创建 axios 实例
const service = axios.create({
  baseURL: '', // 先留空，后续配置
  timeout: 5000 // 请求超时时间
})

// 请求拦截器，自动携带token
service.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers['Authorization'] = 'Bearer ' + token;
  }
  return config;
}, error => Promise.reject(error));

export default service 