<script setup>
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const router = useRouter()

const handleLogout = () => {
  userStore.logout()
  router.push('/login')
}
</script>

<template>
  <div id="app">
    <el-config-provider>
      <el-container>
        <el-header>
          <div class="header-content">
            <h1 class="logo">高考志愿填报指导系统</h1>
            <el-menu
              mode="horizontal"
              :router="true"
              :default-active="$route.path"
            >
              <el-menu-item index="/home">首页</el-menu-item>
              <el-menu-item index="/colleges">高校查询</el-menu-item>
              <el-menu-item index="/majors">专业查询</el-menu-item>
              <el-menu-item index="/chat">智能问答</el-menu-item>
              <el-menu-item index="/rankings">排名榜单</el-menu-item>
            </el-menu>
            <div class="user-actions">
              <template v-if="userStore.isLoggedIn">
                <el-dropdown>
                  <span class="user-name">
                    {{ userStore.userInfo?.username || '用户' }}
                    <el-icon><arrow-down /></el-icon>
                  </span>
                  <template #dropdown>
                    <el-dropdown-menu>
                      <el-dropdown-item @click="router.push('/profile')">个人中心</el-dropdown-item>
                      <el-dropdown-item divided @click="handleLogout">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </template>
              <template v-else>
                <el-button @click="router.push('/login')">登录</el-button>
                <el-button type="primary" @click="router.push('/register')">注册</el-button>
              </template>
            </div>
          </div>
        </el-header>
        <el-main>
          <router-view />
        </el-main>
        <el-footer class="footer">
          <p>高考志愿填报指导系统 - 为河南考生提供智能报考服务</p>
        </el-footer>
      </el-container>
    </el-config-provider>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.el-header {
  background-color: #409eff;
  color: white;
  padding: 0;
}

.header-content {
  display: flex;
  align-items: center;
  padding: 0 20px;
  height: 60px;
}

.logo {
  font-size: 20px;
  margin-right: 40px;
  color: white;
}

.el-menu {
  flex: 1;
  border-bottom: none;
}

.el-menu-item {
  color: white !important;
}

.el-menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1) !important;
}

.user-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-name {
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.el-main {
  flex: 1;
  padding: 20px;
  background-color: #f5f5f5;
}

.footer {
  text-align: center;
  padding: 20px;
  background-color: #333;
  color: #999;
}
</style>
