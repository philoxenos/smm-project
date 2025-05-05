import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/LoginView.vue'),
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/HomeView.vue'),
  },
  {
    path: '/post/:id',
    name: 'PostDetail',
    component: () => import('@/views/PostDetailView.vue'),
    props: true,
  },
  {
    path: '/create',
    name: 'CreatePost',
    component: () => import('@/views/PostFormView.vue'),
  },
  {
    path: '/edit/:id',
    name: 'EditPost',
    component: () => import('@/views/PostFormView.vue'),
    props: true,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
