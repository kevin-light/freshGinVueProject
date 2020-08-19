import axios from 'axios';
import storageService from '../service/storageService';

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL, // process是vue维护的全局变量
  timeout: 1000 * 5,
});

// Add a request interceptor
service.interceptors.request.use((config) => {
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => Promise.reject(error));

export default service; // 导出地址才能用
