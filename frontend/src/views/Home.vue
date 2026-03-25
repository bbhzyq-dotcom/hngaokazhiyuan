<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const hasProfile = ref(false)

onMounted(async () => {
  if (userStore.token) {
    await userStore.getProfile()
    hasProfile.value = !!(userStore.profile?.score)
  }
})
</script>

<template>
  <div class="home-container">
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card class="hero-card">
          <template #header>
            <h2>欢迎使用高考志愿填报指导系统</h2>
          </template>
          <div class="hero-content">
            <p class="hero-desc">
              为河南高考考生提供智能化的志愿填报指导服务，通过AI技术整合全国高校数据、专业排名、历年录取分数线等信息，结合考生自身情况，提供科学的报考建议和未来职业规划指导。
            </p>
            <div class="hero-actions">
              <el-button type="primary" size="large" @click="router.push('/chat')">
                开始智能问答
              </el-button>
              <el-button size="large" @click="router.push('/colleges')">
                查找高校
              </el-button>
            </div>
          </div>
        </el-card>
        
        <el-row :gutter="20" class="feature-row">
          <el-col :span="8">
            <el-card shadow="hover">
              <el-icon size="48" color="#409eff"><Reading /></el-icon>
              <h3>高校查询</h3>
              <p>查询全国高校基本信息、排名、专业设置</p>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover">
              <el-icon size="48" color="#67c23a"><Document /></el-icon>
              <h3>专业查询</h3>
              <p>了解专业详情、培养方案、就业前景</p>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover">
              <el-icon size="48" color="#e6a23c"><ChatDotRound /></el-icon>
              <h3>智能推荐</h3>
              <p>根据分数和偏好智能推荐院校和专业</p>
            </el-card>
          </el-col>
        </el-row>
      </el-col>
      
      <el-col :span="8">
        <el-card v-if="!userStore.isLoggedIn" class="login-card">
          <template #header>
            <h3>登录后享受更多服务</h3>
          </template>
          <div class="login-content">
            <p>登录后可享受：</p>
            <ul>
              <li>保存考生信息</li>
              <li>获取个性化推荐</li>
              <li>智能问答服务</li>
              <li>录取概率分析</li>
            </ul>
            <el-button type="primary" @click="router.push('/login')">立即登录</el-button>
          </div>
        </el-card>
        
        <el-card v-else class="profile-card">
          <template #header>
            <h3>考生信息</h3>
          </template>
          <div v-if="userStore.profile" class="profile-content">
            <el-descriptions :column="1" border>
              <el-descriptions-item label="高考年份">{{ userStore.profile.gaokao_year }}</el-descriptions-item>
              <el-descriptions-item label="分数">{{ userStore.profile.score }}分</el-descriptions-item>
              <el-descriptions-item label="位次">{{ userStore.profile.rank }}名</el-descriptions-item>
              <el-descriptions-item label="偏好地区">{{ userStore.profile.preferred_regions || '未设置' }}</el-descriptions-item>
            </el-descriptions>
            <el-button type="primary" @click="router.push('/profile')">编辑信息</el-button>
          </div>
          <div v-else class="no-profile">
            <p>请先完善您的考生信息</p>
            <el-button type="primary" @click="router.push('/profile')">填写信息</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.home-container {
  max-width: 1200px;
  margin: 0 auto;
}

.hero-card {
  margin-bottom: 20px;
}

.hero-desc {
  font-size: 16px;
  line-height: 1.8;
  color: #666;
  margin-bottom: 20px;
}

.hero-actions {
  display: flex;
  gap: 10px;
}

.feature-row {
  margin-top: 20px;
}

.feature-row .el-card {
  text-align: center;
  padding: 20px 0;
}

.feature-row h3 {
  margin: 15px 0 10px;
}

.feature-row p {
  color: #999;
  font-size: 14px;
}

.login-card, .profile-card {
  margin-bottom: 20px;
}

.login-content ul, .profile-content ul {
  list-style: none;
  padding: 0;
  margin: 15px 0;
}

.login-content li, .profile-content li {
  padding: 5px 0;
}

.login-content li::before, .profile-content li::before {
  content: "✓ ";
  color: #67c23a;
}

.no-profile {
  text-align: center;
  color: #999;
}
</style>
