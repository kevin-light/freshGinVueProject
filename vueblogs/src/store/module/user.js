import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModules = {
  namespace: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO),
  },

  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      // 更新state
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // 更新state
      state.userInfo = userInfo;
    },
  },
  actions: {
    registers(context, { Name, Telephone, Password }) {
      return new Promise((resolve, reject) => {
        userService.register({ Name, Telephone, Password }).then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    logins(context, { Telephone, Password }) {
      return new Promise((resolve, reject) => {
        userService.login({ Telephone, Password }).then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
    logouts({ commit }) {
      // 清除token
      commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');
      // 清除用户信息
      commit('SET_USERINFO', '');
      storageService.set(storageService.USER_INFO, '');

      window.location.reload();
    },
  },
};
export default userModules;
