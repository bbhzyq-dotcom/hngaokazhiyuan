# 高考志愿填报指导系统

基于 AI + RAG 知识库的高考志愿填报智能推荐系统，为河南考生提供全国高校、专业录取分数线查询及 AI 智能填报建议。

## 项目概述

本系统采集了全国 **1300+ 所本科高校** 的完整数据，包括：
- 高校基本信息、排名、录取分数线
- **830 个本科专业**（12个学科门类）
- **35,000+ 条高校-专业关联**数据
- 历年录取分数、位次等

## 功能特性

### 1. 高校查询
- 按省份、城市、类型筛选高校
- 查看高校详细信息（排名、专业、师资、就业率等）
- 历年录取分数线查询

### 2. 专业查询
- 覆盖全国普通高校本科全部 830 个专业
- 按学科门类、专业类筛选
- 查看专业培养目标、学习内容、就业方向

### 3. 智能填报推荐（AI + RAG）
- 基于考生成绩、位次、偏好推荐高校和专业
- AI 分析往年录取数据，预测录取概率
- 生成个性化志愿填报方案

### 4. 问答服务
- AI 智能问答，解答填报疑问
- 基于 RAG 知识库检索，确保回答准确性

## 技术架构

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   前端 Vue  │ ── │  Go API    │ ── │ PostgreSQL  │
│  (Web+小程序)│     │  Gateway   │     │   Redis     │
└─────────────┘     └─────────────┘     └─────────────┘
                           │
                    ┌──────┴──────┐
                    │ Elasticsearch│
                    │  (RAG)      │
                    └─────────────┘
```

| 技术栈 | 说明 |
|--------|------|
| 前端 | Vue.js 3 + Vite + Element Plus |
| 后端 | Go + Gin 框架 |
| 数据库 | PostgreSQL + Redis |
| 搜索引擎 | Elasticsearch（RAG 知识库） |
| AI 模型 | 百度文心一言 / 智谱 GLM |

## 数据概览

| 数据类型 | 数量 | 说明 |
|---------|------|------|
| 高校数据 | 1,300+ 所 | 全国31省市本科高校 |
| 专业数据 | 830 个 | 12个学科门类、93个专业类 |
| 高校-专业关联 | 35,000+ 条 | 各高校开设专业 |
| 录取分数 | 20,000+ 条 | 2022-2024年数据 |
| 排名数据 | 2,100+ 条 | 综合排名+学科排名 |

## 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- Python 3.9+
- PostgreSQL 14+
- Redis 6+
- Elasticsearch 8+

### 1. 克隆项目

```bash
git clone https://github.com/bbhzyq-dotcom/hngaokazhiyuan.git
cd hngaokazhiyuan
```

### 2. 启动后端服务

```bash
cd backend

# 安装依赖
go mod tidy

# 配置数据库连接 (修改 config.yaml)
vim config.yaml

# 启动服务
go run cmd/gateway/main.go
```

### 3. 启动前端

```bash
cd frontend

# 安装依赖
npm install

# 开发模式
npm run dev

# 生产构建
npm run build
```

### 4. Docker 部署

```bash
cd deploy/docker

# 启动所有服务
docker-compose up -d
```

## 数据文件说明

### 核心数据文件 (`/data/`)

| 文件 | 说明 | 记录数 |
|------|------|--------|
| `colleges.json` | 高校基本信息 | 1,300+ |
| `majors_845.json` | 专业信息（12个学科门类） | 830 |
| `college_majors_full.json` | 高校-专业关联 | 35,000+ |
| `scores.json` | 录取分数线 | 20,000+ |
| `rankings.json` | 高校排名 | 2,100+ |
| `majors_structure.json` | 专业门类结构 | - |

### 数据采集脚本 (`/collector/`)

| 文件 | 说明 |
|------|------|
| `collector.py` | 数据采集主脚本 |
| `majors_845.py` | 专业数据库生成脚本 |
| `generate_college_majors.py` | 高校-专业关联生成脚本 |

### 采集脚本使用

```bash
cd collector

# 生成专业数据库
python majors_845.py

# 生成高校-专业关联
python generate_college_majors.py

# 采集全部数据
python collector.py
```

## 项目结构

```
hngaokazhiyuan/
├── backend/                    # Go 后端
│   ├── cmd/gateway/main.go    # 主入口
│   ├── internal/
│   │   ├── config/           # 配置管理
│   │   ├── model/             # 数据模型
│   │   ├── repository/        # 数据访问层
│   │   ├── service/          # 业务逻辑
│   │   │   ├── user_service.go      # 用户服务
│   │   │   ├── data_service.go      # 数据服务
│   │   │   ├── rag_service.go      # RAG 知识库
│   │   │   └── qa_service.go        # AI 问答服务
│   │   ├── handler/           # HTTP 处理器
│   │   ├── middleware/        # 中间件（JWT认证）
│   │   └── router/           # 路由配置
│   └── pkg/                   # 公共工具包
│
├── frontend/                   # Vue.js 前端
│   ├── src/
│   │   ├── views/           # 页面组件
│   │   │   ├── Home.vue       # 首页
│   │   │   ├── Colleges.vue  # 高校查询
│   │   │   ├── Majors.vue    # 专业查询
│   │   │   ├── Chat.vue      # AI 问答
│   │   │   └── Profile.vue   # 个人中心
│   │   ├── api/              # API 请求
│   │   ├── stores/           # 状态管理
│   │   └── router/           # 路由配置
│   └── vite.config.js
│
├── collector/                 # 数据采集
│   ├── collector.py           # 主采集脚本
│   ├── majors_845.py         # 专业数据库
│   └── requirements.txt
│
├── data/                     # 采集数据
│   ├── colleges.json         # 高校数据
│   ├── majors_845.json      # 专业数据
│   └── college_majors_full.json  # 关联数据
│
├── deploy/                    # 部署配置
│   ├── docker/docker-compose.yml
│   └── nginx/nginx.conf
│
└── .monkeycode/             # 项目文档
    └── docs/
```

## API 接口

### 用户接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/user/register` | POST | 用户注册 |
| `/api/user/login` | POST | 用户登录 |
| `/api/user/profile` | GET | 获取用户信息 |

### 数据查询接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/colleges` | GET | 高校列表查询 |
| `/api/colleges/:id` | GET | 高校详情 |
| `/api/majors` | GET | 专业列表查询 |
| `/api/majors/:id` | GET | 专业详情 |
| `/api/scores` | GET | 录取分数查询 |
| `/api/rankings` | GET | 排名查询 |

### AI 服务接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/qa/chat` | POST | AI 问答 |
| `/api/qa/recommend` | POST | 志愿推荐 |

## 专业门类

| 学科门类 | 代码 | 专业数量 |
|---------|------|---------|
| 哲学 | 01 | 4 |
| 经济学 | 02 | 25 |
| 法学 | 03 | 40 |
| 教育学 | 04 | 27 |
| 文学 | 05 | 74 |
| 历史学 | 06 | 9 |
| 理学 | 07 | 53 |
| 工学 | 08 | 301 |
| 农学 | 09 | 53 |
| 医学 | 10 | 75 |
| 管理学 | 12 | 105 |
| 艺术学 | 13 | 64 |

**合计：12 个学科门类、93 个专业类、830 个专业**

## 配置文件

### 后端配置 (`backend/config.yaml`)

```yaml
server:
  port: 8080

database:
  host: localhost
  port: 5432
  user: postgres
  password: your_password
  dbname: gaokao_db

redis:
  host: localhost
  port: 6379

elasticsearch:
  host: localhost
  port: 9200

ai:
  api_key: your_api_key
  model: ernie-bot
```

## License

MIT License
