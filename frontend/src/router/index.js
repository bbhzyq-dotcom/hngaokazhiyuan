import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register.vue')
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/Home.vue')
  },
  {
    path: '/colleges',
    name: 'Colleges',
    component: () => import('@/views/Colleges.vue')
  },
  {
    path: '/colleges/:id',
    name: 'CollegeDetail',
    component: () => import('@/views/CollegeDetail.vue')
  },
  {
    path: '/majors',
    name: 'Majors',
    component: () => import('@/views/Majors.vue')
  },
  {
    path: '/majors/:id',
    name: 'MajorDetail',
    component: () => import('@/views/MajorDetail.vue')
  },
  {
    path: '/chat',
    name: 'Chat',
    component: () => import('@/views/Chat.vue')
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/Profile.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/rankings',
    name: 'Rankings',
    component: () => import('@/views/Rankings.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  if (to.meta.requiresAuth && !userStore.token) {
    next('/login')
  } else {
    next()
  }
})

export default router
