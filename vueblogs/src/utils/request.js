import axios from 'axios';
import storageService from '../service/storageService';

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL, // process全局环境变量及信息
  timeout: 1000 * 5,
});

// 添加请求拦截器
service.interceptors.request.use((config) => {
  // 在发送请求之前做些什么
  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => {
  // 对请求错误做些什么
  return Promise.reject(error);
});

export default service;
