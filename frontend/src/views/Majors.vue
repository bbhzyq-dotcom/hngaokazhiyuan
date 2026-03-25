<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { majorAPI } from '@/api'

const router = useRouter()

const majors = ref([])
const total = ref(0)
const loading = ref(false)
const searchForm = ref({
  keyword: '',
  category: '',
  page: 1,
  page_size: 20
})

const categories = ['工学', '理学', '管理学', '经济学', '法学', '文学', '医学', '教育学']

const loadMajors = async () => {
  loading.value = true
  try {
    const res = await majorAPI.getList(searchForm.value)
    majors.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('Load majors failed:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  searchForm.value.page = 1
  loadMajors()
}

const handlePageChange = (page) => {
  searchForm.value.page = page
  loadMajors()
}

const viewDetail = (id) => {
  router.push(`/majors/${id}`)
}

onMounted(() => {
  loadMajors()
})
</script>

<template>
  <div class="majors-container">
    <el-card>
      <template #header>
        <h2>专业查询</h2>
      </template>
      
      <el-form :model="searchForm" inline>
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="搜索专业名称" clearable />
        </el-form-item>
        <el-form-item label="学科门类">
          <el-select v-model="searchForm.category" placeholder="选择学科" clearable>
            <el-option v-for="c in categories" :key="c" :label="c" :value="c" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
        </el-form-item>
      </el-form>
      
      <el-divider />
      
      <div v-loading="loading">
        <el-table :data="majors" stripe @row-click="viewDetail">
          <el-table-column prop="name" label="专业名称" />
          <el-table-column prop="category" label="学科门类" />
          <el-table-column prop="code" label="专业代码" width="120" />
          <el-table-column prop="degree" label="学位" />
          <el-table-column prop="duration" label="学制" width="80" />
        </el-table>
        
        <el-empty v-if="!loading && majors.length === 0" description="暂无数据" />
        
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
.majors-container {
  max-width: 1000px;
  margin: 0 auto;
}
</style>
