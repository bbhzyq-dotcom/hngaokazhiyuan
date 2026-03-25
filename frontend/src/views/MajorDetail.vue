<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { majorAPI } from '@/api'

const route = useRoute()
const major = ref(null)
const loading = ref(true)

const loadMajor = async () => {
  try {
    const res = await majorAPI.getDetail(route.params.id)
    major.value = res.data
  } catch (error) {
    console.error('Load major failed:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadMajor()
})
</script>

<template>
  <div class="major-detail" v-loading="loading">
    <el-card v-if="major">
      <template #header>
        <h2>{{ major.name }}</h2>
      </template>
      
      <el-descriptions :column="2" border>
        <el-descriptions-item label="学科门类">{{ major.category }}</el-descriptions-item>
        <el-descriptions-item label="专业代码">{{ major.code }}</el-descriptions-item>
        <el-descriptions-item label="学位">{{ major.degree }}</el-descriptions-item>
        <el-descriptions-item label="学制">{{ major.duration }}年</el-descriptions-item>
      </el-descriptions>
      
      <el-divider />
      
      <h3>专业介绍</h3>
      <p class="description">{{ major.description }}</p>
      
      <el-divider />
      
      <h3>核心课程</h3>
      <div class="courses">
        <el-tag v-for="c in (major.core_courses || [])" :key="c" type="info">{{ c }}</el-tag>
      </div>
      
      <el-divider />
      
      <h3>就业情况</h3>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="就业率">{{ major.employment?.rate }}%</el-descriptions-item>
        <el-descriptions-item label="平均月薪">{{ major.employment?.avg_salary }}元</el-descriptions-item>
      </el-descriptions>
      <div class="industries">
        <p>主要就业行业：</p>
        <el-tag v-for="ind in (major.employment?.top_industries || [])" :key="ind" type="success">{{ ind }}</el-tag>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.major-detail {
  max-width: 800px;
  margin: 0 auto;
}

.description {
  line-height: 1.8;
  color: #666;
}

.courses, .industries {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 10px;
}
</style>
