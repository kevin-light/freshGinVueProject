import request from '@/utils/request';

// 用户注册
const register = ({ Name, Telephone, Password }) => request.post('auth/register', { Name, Telephone, Password });

// 用户登录
const login = ({ Telephone, Password }) => request.post('auth/login', { Telephone, Password });

// 获取用户信息
const info = () => request.get('auth/info');

export default {
  register,
  login,
  info,
};
