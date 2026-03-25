<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { collegeAPI } from '@/api'

const route = useRoute()
const college = ref(null)
const majors = ref([])
const loading = ref(true)

const loadCollege = async () => {
  try {
    const res = await collegeAPI.getDetail(route.params.id)
    college.value = res.data
  } catch (error) {
    console.error('Load college failed:', error)
  } finally {
    loading.value = false
  }
}

const loadMajors = async () => {
  try {
    const res = await collegeAPI.getMajors(route.params.id)
    majors.value = res.data.list || []
  } catch (error) {
    console.error('Load majors failed:', error)
  }
}

onMounted(() => {
  loadCollege()
  loadMajors()
})
</script>

<template>
  <div class="college-detail" v-loading="loading">
    <el-card v-if="college">
      <template #header>
        <div class="header">
          <h2>{{ college.name }}</h2>
          <el-tag>{{ college.type }}</el-tag>
          <el-tag type="success">{{ college.level }}</el-tag>
        </div>
      </template>
      
      <el-descriptions :column="2" border>
        <el-descriptions-item label="所在地">{{ college.province }} {{ college.city }}</el-descriptions-item>
        <el-descriptions-item label="创办年份">{{ college.established_year }}年</el-descriptions-item>
        <el-descriptions-item label="院校类型">{{ college.type }}</el-descriptions-item>
        <el-descriptions-item label="学历层次">{{ college.level }}</el-descriptions-item>
        <el-descriptions-item label="院系数量">{{ college.departments }}个</el-descriptions-item>
        <el-descriptions-item label="专任教师">{{ college.faculties }}人</el-descriptions-item>
      </el-descriptions>
      
      <el-divider />
      
      <h3>学校简介</h3>
      <p class="description">{{ college.description }}</p>
      
      <el-divider />
      
      <h3>学科优势</h3>
      <div class="disciplines">
        <el-tag v-for="d in (college.disciplines || [])" :key="d" type="info">{{ d }}</el-tag>
      </div>
      
      <el-divider />
      
      <h3>近年排名</h3>
      <el-descriptions :column="3" border>
        <el-descriptions-item label="综合排名">第{{ college.rankings?.comprehensive }}名</el-descriptions-item>
        <el-descriptions-item label="QS世界排名">第{{ college.rankings?.qs_world }}名</el-descriptions-item>
        <el-descriptions-item label="Nature Index">{{ college.rankings?.nature_index }}</el-descriptions-item>
      </el-descriptions>
    </el-card>
    
    <el-card style="margin-top: 20px">
      <template #header>
        <h3>开设专业</h3>
      </template>
      <el-table :data="majors" stripe>
        <el-table-column prop="name" label="专业名称" />
        <el-table-column prop="category" label="学科门类" />
        <el-table-column prop="degree" label="学位" />
        <el-table-column prop="duration" label="学制" width="80" />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped>
.college-detail {
  max-width: 1000px;
  margin: 0 auto;
}

.header {
  display: flex;
  align-items: center;
  gap: 10px;
}

.header h2 {
  margin: 0;
}

.description {
  line-height: 1.8;
  color: #666;
}

.disciplines {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
</style>
