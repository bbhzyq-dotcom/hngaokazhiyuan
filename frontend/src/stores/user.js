import { defineStore } from 'pinia'
import { userAPI } from '@/api'

export const useUserStore = defineStore('user', {
  state: () => ({
    userInfo: null,
    token: localStorage.getItem('token') || '',
    isLoggedIn: false
  }),
  
  getters: {
    profile: state => state.userInfo?.profile || null
  },
  
  actions: {
    async login(loginData) {
      const res = await userAPI.login(loginData)
      this.token = res.data.token
      this.isLoggedIn = true
      localStorage.setItem('token', this.token)
      return res
    },
    
    async register(registerData) {
      const res = await userAPI.register(registerData)
      this.token = res.data.token
      this.isLoggedIn = true
      localStorage.setItem('token', this.token)
      return res
    },
    
    async getProfile() {
      if (!this.token) return null
      try {
        const res = await userAPI.getProfile()
        this.userInfo = res.data
        return res.data
      } catch (error) {
        this.logout()
        throw error
      }
    },
    
    async updateProfile(data) {
      const res = await userAPI.updateProfile(data)
      await this.getProfile()
      return res
    },
    
    logout() {
      this.userInfo = null
      this.token = ''
      this.isLoggedIn = false
      localStorage.removeItem('token')
    },
    
    async sendCode(phone) {
      return await userAPI.sendCode(phone)
    }
  }
})
