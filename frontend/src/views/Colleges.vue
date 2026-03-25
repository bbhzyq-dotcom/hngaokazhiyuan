<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { collegeAPI } from '@/api'

const router = useRouter()

const colleges = ref([])
const total = ref(0)
const loading = ref(false)
const searchForm = ref({
  keyword: '',
  province: '',
  type: '',
  level: '本科',
  page: 1,
  page_size: 20
})

const loadColleges = async () => {
  loading.value = true
  try {
    const res = await collegeAPI.getList(searchForm.value)
    colleges.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('Load colleges failed:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  searchForm.value.page = 1
  loadColleges()
}

const handlePageChange = (page) => {
  searchForm.value.page = page
  loadColleges()
}

const viewDetail = (id) => {
  router.push(`/colleges/${id}`)
}

onMounted(() => {
  loadColleges()
})
</script>

<template>
  <div class="colleges-container">
    <el-card>
      <template #header>
        <h2>高校查询</h2>
      </template>
      
      <el-form :model="searchForm" inline>
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="搜索高校名称" clearable />
        </el-form-item>
        <el-form-item label="省份">
          <el-input v-model="searchForm.province" placeholder="如：北京" clearable />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="searchForm.type" placeholder="选择类型" clearable>
            <el-option label="综合" value="综合" />
            <el-option label="理工" value="理工" />
            <el-option label="师范" value="师范" />
            <el-option label="文科" value="文科" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
        </el-form-item>
      </el-form>
      
      <el-divider />
      
      <div v-loading="loading">
        <el-row :gutter="20">
          <el-col :span="8" v-for="college in colleges" :key="college.id">
            <el-card shadow="hover" class="college-card" @click="viewDetail(college.id)">
              <h3>{{ college.name }}</h3>
              <p class="college-info">
                <span>{{ college.province }}</span> · 
                <span>{{ college.type }}</span> · 
                <span>{{ college.level }}</span>
              </p>
              <p v-if="college.ranking" class="college-rank">
                综合排名：第{{ college.ranking.comprehensive }}名
              </p>
              <div v-if="college.admission_score" class="college-score">
                <span>理科：{{ college.admission_score.science }}分</span>
                <span>文科：{{ college.admission_score.arts }}分</span>
              </div>
            </el-card>
          </el-col>
        </el-row>
        
        <el-empty v-if="!loading && colleges.length === 0" description="暂无数据" />
        
        <el-pagination
          v-if="total > 0"
          v-model:current-page="searchForm.page"
          :page-size="searchForm.page_size"
          :total="total"
          layout="prev, pager, next"
          @current-change="handlePageChange"
          style="margin-top: 20px; justify-content: center"
        />
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.colleges-container {
  max-width: 1200px;
  margin: 0 auto;
}

.college-card {
  margin-bottom: 20px;
  cursor: pointer;
}

.college-card h3 {
  margin-bottom: 10px;
  color: #409eff;
}

.college-info {
  color: #666;
  font-size: 14px;
  margin-bottom: 5px;
}

.college-rank {
  color: #e6a23c;
  font-size: 14px;
  margin-bottom: 5px;
}

.college-score {
  display: flex;
  gap: 15px;
  font-size: 14px;
  color: #67c23a;
}
</style>
