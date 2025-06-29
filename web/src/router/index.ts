import { createRouter, createWebHashHistory } from 'vue-router';

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('../views/home/index.vue'),
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/auth/Login.vue'),
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('../views/auth/Register.vue'),
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('../views/dashboard/index.vue'),
    },
    {
      path: '/profile',
      name: 'Profile',
      component: () => import('../views/profile/index.vue'),
    },
    {
      path: '/settings',
      name: 'Settings',
      component: () => import('../views/settings/index.vue'),
    },
    {
      path: '/channels',
      name: 'Channels',
      component: () => import('../views/channels/index.vue'),
    },
    {
      path: '/topics',
      name: 'Topics',
      component: () => import('../views/topics/index.vue'),
    },
    {
      path: '/topics/:id',
      name: 'TopicDetail',
      component: () => import('../views/topics/detail.vue'),
    },
  ],
})

export default router