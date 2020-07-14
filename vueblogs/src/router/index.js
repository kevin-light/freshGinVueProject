import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '@/store';
import Home from '../views/Home.vue';
import userRoutes from './module/user';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

// https://router.vuejs.org/zh/guide/advanced/navigation-guards.html
router.beforeEach((to, from, next) => {
  if (to.meta.auth) { // 判断用户是否需要登录 = 即将要进入的目标 路由对象
    // 判断用户是否登录
    if (store.state.userModule.token) {
      // 判断token有效性（有没有过期，需要后台发token的时候带上有效时间，过期重新请求
      next();
    } else {
    // 跳转登录
      router.push({ name: 'login' });
    }
  } else {
    next();
  }
});

export default router;
