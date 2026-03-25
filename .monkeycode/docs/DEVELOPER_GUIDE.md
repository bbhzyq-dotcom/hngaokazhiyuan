# 高考志愿填报指导系统 - 开发指南

## 1. 项目结构

```
gaokao-advisor/
├── backend/                      # 后端Go服务
│   ├── cmd/                      # 入口目录
│   │   └── gateway/              # API网关入口
│   ├── internal/                 # 内部包
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
│   ├── config.yaml              # 配置文件
│   └── go.mod
│
├── frontend/                     # 前端Vue项目
│   ├── src/
│   │   ├── api/                # API调用
│   │   ├── components/         # 公共组件
│   │   ├── views/              # 页面
│   │   ├── stores/             # Pinia状态
│   │   ├── router/             # 路由配置
│   │   └── main.js
│   ├── vite.config.js
│   └── package.json
│
├── deploy/                      # 部署配置
│   └── docker/
│       └── docker-compose.yml
│
└── docs/                        # 项目文档
```

## 2. 开发环境要求

| 软件 | 版本 | 说明 |
|-----|------|------|
| Go | 1.21+ | 后端运行环境 |
| Node.js | 18+ | 前端构建工具 |
| Python | 3.10+ | 数据采集脚本 |
| Docker | 24+ | 容器化支持 |

## 3. 快速开始

### 3.1 启动依赖服务

```bash
cd deploy/docker
docker-compose up -d
```

### 3.2 启动后端

```bash
cd backend
go mod tidy
go run cmd/gateway/main.go
```

### 3.3 启动前端

```bash
cd frontend
npm install
npm run dev
```

## 4. 后端开发规范

### 4.1 项目初始化

```bash
cd backend
go mod init github.com/gaokao-advisor/backend
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/redis/go-redis/v9
go get github.com/elastic/go-elasticsearch/v8
go get github.com/golang-jwt/jwt/v5
```

### 4.2 代码分层结构

```
internal/
├── handler/     # HTTP处理层，参数校验，调用service
├── service/     # 业务逻辑，事务管理
├── repository/  # 数据访问，数据库和ES操作
├── model/       # 数据模型，DTO定义
└── middleware/  # 中间件，认证、日志、限流
```

### 4.3 API响应封装

```go
// pkg/response/response.go
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(200, Response{Code: 0, Message: "success", Data: data})
}
```

## 5. 前端开发规范

### 5.1 项目初始化

```bash
cd frontend
npm create vite@latest . -- --template vue
npm install vue-router@4 pinia element-plus @element-plus/icons-vue axios
```

### 5.2 目录结构

```
src/
├── api/              # API接口定义
├── components/       # 公共组件
├── views/            # 页面
├── stores/           # Pinia状态
└── router/           # 路由
```

### 5.3 API调用封装

```javascript
// src/api/request.js
import axios from 'axios'

const request = axios.create({
  baseURL: '/api/v1',
  timeout: 30000
})

request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default request
```

## 6. 部署指南

### 6.1 Docker Compose本地部署

```bash
cd deploy/docker
docker-compose up -d
```

启动服务：
- PostgreSQL: localhost:5432
- Redis: localhost:6379
- Elasticsearch: localhost:9200

## 7. 常见问题

### 7.1 数据库连接问题

检查PostgreSQL是否启动：`docker-compose ps postgres`

### 7.2 前端API调用跨域问题

检查vite.config.js中的代理配置是否正确
