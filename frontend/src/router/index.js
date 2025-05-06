import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '@/composables/useAuth'

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
  {
    path: '/my-posts',
    name: 'MyPosts',
    component: () => import('@/views/MyPostsView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  const { user } = useAuth()
  const isLoggedIn = !!user.value
  if (to.path !== '/login' && !isLoggedIn) {
    next('/login')
  } else if (to.path === '/login' && isLoggedIn) {
    next('/')
  } else {
    next()
  }
})

export default router
