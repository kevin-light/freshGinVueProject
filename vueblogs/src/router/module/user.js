const userRoutes = [
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/register/Register.vue'), // 组件的 匿名函数异步惰性加载
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/Login.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    component: () => import('@/views/profile/Profile.vue'),
  },
  {
    path: '/demo',
    name: 'demo',
    component: () => import('@/views/Demotesting.vue'),
  },
];

export default userRoutes;
