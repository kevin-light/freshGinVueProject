// 本地缓存服务

const PREFIX = 'ginessential_';

// 定义 user 模块的常量
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// localStorage 存储方式
const set = (key, data) => {
  localStorage.setItem(key, data);
};

// 读取
const get = (key) => localStorage.getItem(key);

export default {
  set,
  get,
  USER_TOKEN,
  USER_INFO,
};
