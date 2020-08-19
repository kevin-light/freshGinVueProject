// 本地缓存服务

const PREFIX = 'ginblogs_';

// user 模块常量

// const USER_PREFIX = 'PREFIX + 'user'';
const USER_PREFIX = `${PREFIX}user`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 使用localStorage存储
const set = (key, data) => {
  localStorage.setItem(key, data);
};

// 读取localstorage
const get = (key) => localStorage.getItem(key);
export default {
  set,
  get,
  USER_TOKEN,
  USER_INFO,
};
