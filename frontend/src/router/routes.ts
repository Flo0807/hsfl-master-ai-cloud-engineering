import {RouteRecordRaw} from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/auth',
    meta: {
      authRequired: false
    },
    component: () => import('layouts/EmptyLayout.vue'),
    children: [
      {
        name: 'login', path: 'login', meta: {
          authRequired: false
        }, component: () => import('pages/auth/LoginPage.vue')
      },
      {
        name: 'register', path: 'register', meta: {
          authRequired: false
        }, component: () => import('pages/auth/RegisterPage.vue')
      },
    ],
  },
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{
      name: 'home', path: '', meta: {
        authRequired: true
      }, component: () => import('pages/IndexPage.vue')
    }],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    meta: {
      authRequired: false
    },
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
