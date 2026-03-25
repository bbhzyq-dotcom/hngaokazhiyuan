<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const registerForm = ref({
  username: '',
  password: '',
  confirmPassword: '',
  phone: '',
  code: '',
  user_type: 'student'
})

const loading = ref(false)
const countdown = ref(0)

const sendCode = async () => {
  if (countdown.value > 0) return
  try {
    await userStore.sendCode(registerForm.value.phone)
    countdown.value = 60
    const timer = setInterval(() => {
      countdown.value--
      if (countdown.value <= 0) {
        clearInterval(timer)
      }
    }, 1000)
  } catch (error) {
    console.error('Send code failed:', error)
  }
}

const handleRegister = async () => {
  if (registerForm.value.password !== registerForm.value.confirmPassword) {
    alert('两次输入的密码不一致')
    return
  }
  
  loading.value = true
  try {
    await userStore.register(registerForm.value)
    router.push('/profile')
  } catch (error) {
    console.error('Register failed:', error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="register-container">
    <el-card class="register-card">
      <template #header>
        <h2>用户注册</h2>
      </template>
      <el-form :model="registerForm" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="registerForm.username" placeholder="请输入用户名" />
        </el-form-item>
        
        <el-form-item label="手机号">
          <el-input v-model="registerForm.phone" placeholder="请输入手机号">
            <template #append>
              <el-button @click="sendCode" :disabled="countdown > 0">
                {{ countdown > 0 ? `${countdown}s` : '获取验证码' }}
              </el-button>
            </template>
          </el-input>
        </el-form-item>
        
        <el-form-item label="验证码">
          <el-input v-model="registerForm.code" placeholder="请输入验证码" />
        </el-form-item>
        
        <el-form-item label="密码">
          <el-input 
            v-model="registerForm.password" 
            type="password" 
            placeholder="请输入密码（8-20位）"
            show-password
          />
        </el-form-item>
        
        <el-form-item label="确认密码">
          <el-input 
            v-model="registerForm.confirmPassword" 
            type="password" 
            placeholder="请再次输入密码"
            show-password
          />
        </el-form-item>
        
        <el-form-item label="用户类型">
          <el-radio-group v-model="registerForm.user_type">
            <el-radio label="student">考生</el-radio>
            <el-radio label="parent">家长</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleRegister" :loading="loading" style="width: 100%">
            注册
          </el-button>
        </el-form-item>
      </el-form>
      
      <div class="login-link">
        已有账号？ <router-link to="/login">立即登录</router-link>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 120px);
}

.register-card {
  width: 450px;
}

.login-link {
  text-align: center;
  margin-top: 10px;
}
</style>
