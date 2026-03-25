<script setup>
import { ref, onMounted } from 'vue'
import { rankingAPI } from '@/api'

const rankings = ref([])
const total = ref(0)
const loading = ref(false)
const activeTab = ref('colleges')
const searchForm = ref({
  year: 2024,
  ranking_type: 'comprehensive',
  category: '',
  page: 1,
  page_size: 50
})

const loadCollegeRankings = async () => {
  loading.value = true
  try {
    const res = await rankingAPI.getCollegeRankings(searchForm.value)
    rankings.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('Load rankings failed:', error)
  } finally {
    loading.value = false
  }
}

const loadMajorRankings = async () => {
  loading.value = true
  try {
    const res = await rankingAPI.getMajorRankings(searchForm.value)
    rankings.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('Load rankings failed:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  searchForm.value.page = 1
  if (activeTab.value === 'colleges') {
    loadCollegeRankings()
  } else {
    loadMajorRankings()
  }
}

const handleTabChange = (tab) => {
  activeTab.value = tab
  handleSearch()
}

const handlePageChange = (page) => {
  searchForm.value.page = page
  if (activeTab.value === 'colleges') {
    loadCollegeRankings()
  } else {
    loadMajorRankings()
  }
}

onMounted(() => {
  loadCollegeRankings()
})
</script>

<template>
  <div class="rankings-container">
    <el-card>
      <template #header>
        <h2>排名榜单</h2>
      </template>
      
      <el-tabs v-model="activeTab" @tab-change="handleTabChange">
        <el-tab-pane label="高校排名" name="colleges">
          <el-form :model="searchForm" inline>
            <el-form-item label="年份">
              <el-select v-model="searchForm.year">
                <el-option label="2024" :value="2024" />
                <el-option label="2023" :value="2023" />
                <el-option label="2022" :value="2022" />
              </el-select>
            </el-form-item>
            <el-form-item label="类型">
              <el-select v-model="searchForm.ranking_type">
                <el-option label="综合排名" value="comprehensive" />
                <el-option label="专业排名" value="subject" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSearch">查询</el-button>
            </el-form-item>
          </el-form>
          
          <el-table :data="rankings" stripe v-loading="loading">
            <el-table-column type="index" label="排名" width="80" />
            <el-table-column prop="name" label="高校名称" />
            <el-table-column prop="province" label="所在地" />
            <el-table-column prop="type" label="类型" />
            <el-table-column prop="score" label="综合得分" />
          </el-table>
        </el-tab-pane>
        
        <el-tab-pane label="专业排名" name="majors">
          <el-form :model="searchForm" inline>
            <el-form-item label="年份">
              <el-select v-model="searchForm.year">
                <el-option label="2024" :value="2024" />
                <el-option label="2023" :value="2023" />
              </el-select>
            </el-form-item>
            <el-form-item label="学科">
              <el-select v-model="searchForm.category" placeholder="选择学科">
                <el-option label="工学" value="工学" />
                <el-option label="理学" value="理学" />
                <el-option label="管理学" value="管理学" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleSearch">查询</el-button>
            </el-form-item>
          </el-form>
          
          <el-table :data="rankings" stripe v-loading="loading">
            <el-table-column type="index" label="排名" width="80" />
            <el-table-column prop="name" label="专业名称" />
            <el-table-column prop="category" label="学科门类" />
            <el-table-column prop="schools_count" label="开设学校数" />
            <el-table-column prop="score" label="综合得分" />
          </el-table>
        </el-tab-pane>
      </el-tabs>
      
      <el-pagination
        v-if="total > 0"
        v-model:current-page="searchForm.page"
        :page-size="searchForm.page_size"
        :total="total"
        layout="prev, pager, next"
        @current-change="handlePageChange"
        style="margin-top: 20px; justify-content: center"
      />
    </el-card>
  </div>
</template>

<style scoped>
.rankings-container {
  max-width: 1000px;
  margin: 0 auto;
}
</style>
