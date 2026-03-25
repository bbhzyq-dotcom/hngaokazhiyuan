# 高考志愿填报指导系统 - 开发指南

## 1. 开发环境搭建

### 1.1 系统要求

| 软件 | 版本 | 说明 |
|-----|------|------|
| Go | 1.21+ | 后端运行环境 |
| Node.js | 18+ | 前端构建工具 |
| Python | 3.10+ | 数据采集脚本 |
| Docker | 24+ | 容器化支持 |
| Git | 2.40+ | 版本控制 |

### 1.2 开发工具推荐

- **后端**: GoLand / VSCode + Go插件
- **前端**: WebStorm / VSCode + Vue插件
- **数据库**: DataGrip / pgAdmin
- **API调试**: Postman / Apifox
- **Redis**: RedisInsight

### 1.3 环境变量配置

**后端环境变量 (.env)**
```bash
# 应用配置
APP_ENV=development
APP_PORT=8080
APP_HOST=0.0.0.0

# 数据库配置
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=gaokao
POSTGRES_MAX_OPEN_CONNS=25
POSTGRES_MAX_IDLE_CONNS=5

# Redis配置
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Elasticsearch配置
ES_HOSTS=localhost:9200
ES_USERNAME=elastic
ES_PASSWORD=changeme

# JWT配置
JWT_SECRET=your-secret-key-here
JWT_EXPIRES_HOURS=720

# 短信服务配置
SMS_PROVIDER=aliyun
SMS_ACCESS_KEY=
SMS_ACCESS_SECRET=
SMS_SIGN_NAME=高考志愿助手

# AI大模型配置
AI_PROVIDER=baidu
AI_API_KEY=
AI_SECRET_KEY=
AI_MODEL=ernie-4.0
```

### 1.4 服务启动顺序

```
1. Docker服务
   ├── PostgreSQL
   ├── Redis
   └── Elasticsearch

2. 后端服务 (Go)
   ├── user-service
   ├── data-service
   └── qa-service

3. 数据采集服务 (Python)
   └── collector

4. 前端服务 (Vue.js)
```

## 2. 项目结构

```
gaokao-advisor/
├── backend/                      # 后端Go服务
│   ├── cmd/                      # 入口目录
│   │   ├── gateway/              # API网关入口
│   │   ├── user-service/         # 用户服务入口
│   │   ├── data-service/         # 数据服务入口
│   │   └── qa-service/          # 问答服务入口
│   ├── internal/                # 内部包
│   │   ├── config/              # 配置管理
│   │   ├── handler/             # HTTP处理层
│   │   ├── service/             # 业务逻辑层
│   │   ├── repository/          # 数据访问层
│   │   ├── model/               # 数据模型
│   │   ├── middleware/          # 中间件
│   │   └── router/              # 路由注册
│   ├── pkg/                     # 公共包
│   │   ├── utils/               # 工具函数
│   │   ├── response/            # 响应封装
│   │   └── errors/              # 错误定义
│   ├── scripts/                 # 脚本
│   │   └── migrations/          # 数据库迁移
│   └── go.mod
│
├── frontend/                     # 前端Vue项目
│   ├── src/
│   │   ├── api/                # API调用
│   │   ├── components/         # 公共组件
│   │   ├── views/              # 页面
│   │   ├── stores/             # Pinia状态
│   │   ├── router/             # 路由配置
│   │   ├── utils/             # 工具函数
│   │   └── App.vue
│   ├── public/
│   └── package.json
│
├── collector/                   # 数据采集服务
│   ├── spiders/                 # 爬虫脚本
│   ├── processors/             # 数据处理
│   ├── loaders/                # 数据加载
│   ├── requirements.txt
│   └── run.py
│
├── data/                        # 数据文件
│   ├── colleges/               # 高校数据
│   ├── majors/                 # 专业数据
│   └── scores/                 # 录取分数
│
├── deploy/                      # 部署配置
│   ├── docker/
│   ├── kubernetes/
│   └── scripts/
│
└── docs/                        # 文档
```

## 3. 后端开发规范

### 3.1 项目初始化

```bash
# 创建项目目录
mkdir gaokao-advisor && cd gaokao-advisor

# 初始化Go模块
go mod init github.com/xxx/gaokao-advisor

# 安装依赖
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/redis/go-redis/v9
go get github.com/elastic/go-elasticsearch/v8
```

### 3.2 代码分层结构

```
internal/
├── handler/     # HTTP处理层，参数校验，调用service
├── service/     # 业务逻辑，事务管理
├── repository/  # 数据访问，数据库和ES操作
├── model/       # 数据模型，DTO定义
└── middleware/  # 中间件，认证、日志、限流
```

### 3.3 API响应封装

```go
// pkg/response/response.go
package response

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(200, Response{
        Code:    0,
        Message: "success",
        Data:    data,
    })
}

func Error(c *gin.Context, code int, message string) {
    c.JSON(200, Response{
        Code:    code,
        Message: message,
    })
}
```

### 3.4 Service编写规范

```go
// internal/service/user_service.go
package service

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetProfile(userID uint64) (*model.UserProfile, error) {
    profile, err := s.repo.FindProfile(userID)
    if err != nil {
        return nil, err
    }
    // 业务逻辑处理
    return profile, nil
}
```

### 3.5 Repository编写规范

```go
// internal/repository/user_repository.go
package repository

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindProfile(userID uint64) (*model.UserProfile, error) {
    var profile model.UserProfile
    err := r.db.Where("user_id = ?", userID).First(&profile).Error
    if err != nil {
        return nil, err
    }
    return &profile, nil
}
```

## 4. 前端开发规范

### 4.1 项目初始化

```bash
# 创建Vue项目
npm create vite@latest frontend -- --template vue

# 进入目录
cd frontend

# 安装依赖
npm install
npm install vue-router@4 pinia element-plus @element-plus/icons-vue
npm install axios

# 安装开发依赖
npm install -D @vitejs/plugin-vue
```

### 4.2 目录结构

```
src/
├── api/              # API接口定义
│   ├── user.js
│   ├── college.js
│   └── qa.js
├── components/       # 公共组件
│   ├── CollegeCard.vue
│   ├── MajorCard.vue
│   └── ScoreChart.vue
├── views/            # 页面
│   ├── Home.vue
│   ├── Search.vue
│   ├── Chat.vue
│   └── Profile.vue
├── stores/           # Pinia状态
│   ├── user.js
│   └── chat.js
├── router/           # 路由
│   └── index.js
└── App.vue
```

### 4.3 API调用封装

```javascript
// api/request.js
import axios from 'axios'
import { ElMessage } from 'element-plus'

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 30000
})

request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

request.interceptors.response.use(
  response => response.data,
  error => {
    ElMessage.error(error.response?.data?.message || '请求失败')
    return Promise.reject(error)
  }
)

export default request
```

### 4.4 页面开发示例

```vue
<template>
  <div class="college-list">
    <el-input v-model="keyword" placeholder="搜索高校" @search="handleSearch" />
    <div class="college-grid">
      <CollegeCard v-for="college in colleges" :key="college.id" :college="college" />
    </div>
    <el-pagination
      v-model:current-page="page"
      :page-size="pageSize"
      :total="total"
      @current-change="handlePageChange"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getCollegeList } from '@/api/college'
import CollegeCard from '@/components/CollegeCard.vue'

const keyword = ref('')
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)
const colleges = ref([])

const handleSearch = async () => {
  const res = await getCollegeList({ keyword: keyword.value, page: page.value })
  colleges.value = res.data.list
  total.value = res.data.total
}

const handlePageChange = (newPage) => {
  page.value = newPage
  handleSearch()
}

onMounted(handleSearch)
</script>
```

## 5. RAG知识库构建

### 5.1 数据采集流程

```python
# collector/run.py
import asyncio
from collector.spiders.college_spider import CollegeSpider
from collector.spiders.major_spider import MajorSpider
from collector.spiders.score_spider import ScoreSpider
from collector.loaders.es_loader import ESLoader

async def main():
    # 初始化ES加载器
    loader = ESLoader()
    
    # 采集高校数据
    college_spider = CollegeSpider()
    colleges = await college_spider crawl()
    await loader.load_colleges(colleges)
    
    # 采集专业数据
    major_spider = MajorSpider()
    majors = await major_spider crawl()
    await loader.load_majors(majors)
    
    # 采集录取分数
    score_spider = ScoreSpider()
    scores = await score_spider crawl()
    await loader.load_scores(scores)

if __name__ == '__main__':
    asyncio.run(main())
```

### 5.2 向量化配置

```python
# collector/processors/embedding.py
from langchain.embeddings import OpenAIEmbeddings

class EmbeddingProcessor:
    def __init__(self):
        self.embedding = OpenAIEmbeddings(
            model='text-embedding-ada-002',
            openai_api_key='your-api-key'
        )
    
    def embed_texts(self, texts):
        return self.embedding.embed_documents(texts)
```

### 5.3 ES索引创建

```bash
# 创建高校索引
curl -X PUT "localhost:9200/colleges" -H "Content-Type: application/json" -d'
{
  "mappings": {
    "properties": {
      "id": { "type": "keyword" },
      "name": { "type": "text", "analyzer": "ik_max_word" },
      "description": { "type": "text", "analyzer": "ik_max_word" },
      "embedding": { "type": "dense_vector", "dims": 1536, "index": true, "similarity": "cosine" }
    }
  }
}'
```

## 6. 测试指南

### 6.1 后端单元测试

```bash
# 运行所有测试
go test ./...

# 运行指定包的测试
go test ./internal/service/...

# 显示测试覆盖率
go test -cover ./...
```

### 6.2 前端单元测试

```bash
# 运行测试
npm run test

# 生成覆盖率报告
npm run test:cov
```

### 6.3 API接口测试

```bash
# 启动后端服务
go run cmd/gateway/main.go

# 使用curl测试
curl -X POST http://localhost:8080/api/v1/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"12345678","phone":"13800138000","code":"123456"}'
```

## 7. 部署指南

### 7.1 Docker Compose本地部署

```yaml
# deploy/docker/docker-compose.yml
version: '3.8'
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: gaokao
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7
    command: redis-server --requirepass redis123

  elasticsearch:
    image: elasticsearch:8.11.0
    environment:
      - discovery.type=single-node
      - ES_JAVA_OPTS=-Xms2g -Xmx2g
    volumes:
      - es_data:/usr/share/elasticsearch/data

  gateway:
    build: ../backend
    ports:
      - "8080:8080"
    depends_on:
      - postgres
      - redis
      - elasticsearch

volumes:
  postgres_data:
  es_data:
```

### 7.2 启动服务

```bash
# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f gateway
```

## 8. 常见问题

### 8.1 数据库连接问题

**问题**: 连接PostgreSQL超时
**解决**: 
1. 检查PostgreSQL是否启动 `docker-compose ps postgres`
2. 检查端口是否冲突 `netstat -an | grep 5432`
3. 检查防火墙设置

### 8.2 Elasticsearch连接问题

**问题**: ES连接失败
**解决**:
1. 检查ES是否启动成功
2. 验证ES用户名密码
3. 检查ES内存设置是否足够

### 8.3 前端API调用跨域问题

**问题**: CORS跨域错误
**解决**:
1. 检查Nginx反向代理配置
2. 确认后端CORS配置允许前端域名
3. 检查OPTIONS预检请求是否正常
