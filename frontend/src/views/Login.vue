<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const loginForm = ref({
  login_type: 'password',
  phone: '',
  password: '',
  code: ''
})

const loading = ref(false)

const handleLogin = async () => {
  loading.value = true
  try {
    await userStore.login(loginForm.value)
    router.push('/home')
  } catch (error) {
    console.error('Login failed:', error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <h2>用户登录</h2>
      </template>
      <el-form :model="loginForm" label-width="80px">
        <el-form-item label="登录方式">
          <el-radio-group v-model="loginForm.login_type">
            <el-radio label="password">密码登录</el-radio>
            <el-radio label="phone">验证码登录</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item label="手机号">
          <el-input v-model="loginForm.phone" placeholder="请输入手机号" />
        </el-form-item>
        
        <el-form-item v-if="loginForm.login_type === 'password'" label="密码">
          <el-input 
            v-model="loginForm.password" 
            type="password" 
            placeholder="请输入密码"
            show-password
          />
        </el-form-item>
        
        <el-form-item v-if="loginForm.login_type === 'phone'" label="验证码">
          <el-input v-model="loginForm.code" placeholder="请输入验证码">
            <template #append>
              <el-button>获取验证码</el-button>
            </template>
          </el-input>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleLogin" :loading="loading" style="width: 100%">
            登录
          </el-button>
        </el-form-item>
      </el-form>
      
      <div class="register-link">
        还没有账号？ <router-link to="/register">立即注册</router-link>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 120px);
}

.login-card {
  width: 400px;
}

.register-link {
  text-align: center;
  margin-top: 10px;
}
</style>
