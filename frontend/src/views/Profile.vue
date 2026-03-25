<script setup>
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()

const profileForm = ref({
  province: '河南省',
  gaokao_year: 2025,
  score: 0,
  rank: 0,
  preferred_subjects: [],
  preferred_regions: [],
  interest_tags: []
})

const loading = ref(false)

const subjects = ['计算机', '电子信息', '机械', '医学', '经济管理', '法学', '文学', '理学', '工学']
const regions = ['北京', '上海', '浙江', '江苏', '广东', '湖北', '四川', '河南', '陕西']
const interests = ['人工智能', '软件开发', '数据分析', '金融', '机械设计', '医学研究', '学术研究']

const loadProfile = () => {
  if (userStore.profile) {
    profileForm.value = {
      province: userStore.profile.province || '河南省',
      gaokao_year: userStore.profile.gaokao_year || 2025,
      score: userStore.profile.score || 0,
      rank: userStore.profile.rank || 0,
      preferred_subjects: userStore.profile.preferred_subjects || [],
      preferred_regions: userStore.profile.preferred_regions || [],
      interest_tags: userStore.profile.interest_tags || []
    }
  }
}

const handleSubmit = async () => {
  loading.value = true
  try {
    await userStore.updateProfile(profileForm.value)
    ElMessage.success('保存成功')
  } catch (error) {
    ElMessage.error('保存失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  if (userStore.token) {
    userStore.getProfile().then(() => loadProfile())
  }
})
</script>

<template>
  <div class="profile-container">
    <el-card>
      <template #header>
        <h2>个人中心 - 考生信息</h2>
      </template>
      
      <el-form :model="profileForm" label-width="120px">
        <el-form-item label="所在省份">
          <el-input v-model="profileForm.province" disabled />
        </el-form-item>
        
        <el-form-item label="高考年份">
          <el-select v-model="profileForm.gaokao_year">
            <el-option label="2025" :value="2025" />
            <el-option label="2024" :value="2024" />
            <el-option label="2023" :value="2023" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="高考分数">
          <el-input-number v-model="profileForm.score" :min="0" :max="750" />
        </el-form-item>
        
        <el-form-item label="省内位次">
          <el-input-number v-model="profileForm.rank" :min="0" :max="1000000" />
        </el-form-item>
        
        <el-form-item label="偏好专业">
          <el-select v-model="profileForm.preferred_subjects" multiple placeholder="选择偏好专业">
            <el-option v-for="s in subjects" :key="s" :label="s" :value="s" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="偏好地区">
          <el-select v-model="profileForm.preferred_regions" multiple placeholder="选择偏好地区">
            <el-option v-for="r in regions" :key="r" :label="r" :value="r" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="兴趣标签">
          <el-select v-model="profileForm.interest_tags" multiple placeholder="选择兴趣标签">
            <el-option v-for="i in interests" :key="i" :label="i" :value="i" />
          </el-select>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleSubmit" :loading="loading">
            保存信息
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.profile-container {
  max-width: 600px;
  margin: 0 auto;
}
</style>
