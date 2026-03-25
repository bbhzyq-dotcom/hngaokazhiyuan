# 高考志愿填报指导系统 - 实施计划

## 阶段一：项目基础架构

- [x] 1. 初始化 Go 后端项目结构
  - 创建 cmd/ 目录结构（gateway, user-service, data-service, qa-service）
  - 创建 internal/ 包结构（config, handler, service, repository, model, middleware, router）
  - 创建 pkg/ 公共包（utils, response, errors）
  - 初始化 go.mod，添加依赖（Gin, GORM, go-redis, go-elasticsearch, jwt-go）
  - 配置 Viper 管理环境变量

- [x] 2. 创建前端 Vue.js 项目结构
  - 使用 Vite 创建 Vue 3 项目
  - 安装依赖（Vue Router, Pinia, Element Plus, Axios）
  - 创建目录结构（api, components, views, stores, router）
  - 配置项目环境变量和代理

- [x] 3. 配置 Docker 开发环境
  - 创建 docker-compose.yml（PostgreSQL, Redis, Elasticsearch）
  - 创建各服务的 Dockerfile
  - 配置网络和卷挂载

## 阶段二：数据模型与存储层

- [x] 4. 设计并创建 PostgreSQL 数据表
  - [x] 4.1 创建用户相关表（users, student_profiles）
  - [x] 4.2 创建高校相关表（colleges）
  - [x] 4.3 创建专业相关表（majors）
  - [x] 4.4 创建录取分数表（admission_scores）
  - [x] 4.5 创建对话历史表（conversations, messages）

- [x] 5. 创建 Elasticsearch 索引
  - [x] 5.1 创建 colleges 索引
  - [x] 5.2 创建 majors 索引
  - [x] 5.3 创建 admission_scores 索引
  - [x] 5.4 创建 knowledge_base 索引（报考知识、规划知识）

- [x] 6. 实现 Repository 层
  - [x] 6.1 实现 UserRepository
  - [x] 6.2 实现 CollegeRepository
  - [x] 6.3 实现 MajorRepository
  - [x] 6.4 实现 ScoreRepository

## 阶段三：用户服务与认证

- [x] 7. 实现用户服务核心功能
  - [x] 7.1 实现用户注册接口 POST /api/v1/user/register
  - [x] 7.2 实现用户登录接口 POST /api/v1/user/login
  - [x] 7.3 实现获取用户信息 GET /api/v1/user/profile
  - [x] 7.4 实现更新考生信息 PUT /api/v1/user/profile
  - [x] 7.5 实现更新考生偏好 PUT /api/v1/user/preferences

- [x] 8. 实现 JWT 认证中间件
- [x] 9. 实现短信验证码服务（模拟）

## 阶段四：数据服务 API

- [x] 10. 实现高校数据接口
  - [x] 10.1 GET /api/v1/colleges - 高校列表查询
  - [x] 10.2 GET /api/v1/colleges/:id - 高校详情
  - [x] 10.3 GET /api/v1/colleges/:id/majors - 高校专业列表

- [x] 11. 实现专业数据接口
  - [x] 11.1 GET /api/v1/majors - 专业列表查询
  - [x] 11.2 GET /api/v1/majors/:id - 专业详情

- [x] 12. 实现录取数据接口
  - [x] 12.1 GET /api/v1/scores - 分数线查询
  - [x] 12.2 GET /api/v1/scores/probability - 录取概率计算

- [x] 13. 实现排名接口
  - [x] 13.1 GET /api/v1/rankings/colleges - 学校排名查询
  - [x] 13.2 GET /api/v1/rankings/majors - 专业排名查询
    - 按年份、学科门类过滤

## 阶段五：RAG 知识库构建

- [x] 14. 实现数据向量化模块
  - [x] 14.1 实现文本向量化接口
  - [x] 14.2 实现高校/专业描述向量化

- [x] 15. 实现 RAG 检索模块
  - [x] 15.1 实现混合检索策略
  - [x] 15.2 实现检索结果重排序
  - [x] 15.3 实现上下文组装

- [x] 16. 实现知识库数据导入
  - [x] 16.1 导入高校知识数据
  - [x] 16.2 导入专业知识数据
  - [x] 16.3 导入录取知识数据
  - [x] 16.4 导入报考知识数据
  - [x] 16.5 导入规划知识数据

## 阶段六：问答服务与 AI 集成

- [x] 17. 实现问答服务核心功能
  - [x] 17.1 POST /api/v1/qa/chat - AI 问答接口
  - [x] 17.2 GET /api/v1/qa/history - 对话历史
  - [x] 17.3 POST /api/v1/qa/recommend - 志愿推荐
  - [x] 17.4 GET /api/v1/qa/colleges - 推荐院校
  - [x] 17.5 GET /api/v1/qa/majors - 推荐专业

- [x] 18. 实现 AI 大模型集成
  - [x] 18.1 实现模型客户端
  - [x] 18.2 实现提示词模板
  - [x] 18.3 实现答案后处理

## 阶段七：前端开发

- [x] 19. 实现前端基础页面
  - [x] 19.1 登录/注册页面
  - [x] 19.2 个人中心页面

- [x] 20. 实现高校/专业查询页面
  - [x] 20.1 高校列表页
  - [x] 20.2 高校详情页
  - [x] 20.3 专业列表页
  - [x] 20.4 专业详情页

- [x] 21. 实现智能问答页面
  - [x] 21.1 AI 问答界面
  - [x] 21.2 志愿推荐界面

- [x] 22. 实现分数线查询页面
  - [x] 22.1 分数线查询

- [x] 23. 实现排名榜单页面
  - [x] 23.1 学校排名
  - [x] 23.2 专业排名

## 阶段八：数据采集模块

- [x] 24. 实现数据采集爬虫
  - [x] 24.1 高校信息爬虫
  - [x] 24.2 专业信息爬虫
  - [x] 24.3 录取分数爬虫

- [x] 25. 实现数据处理和导入
  - [x] 25.1 数据清洗
  - [x] 25.2 数据存储

- [x] 26. 实现采集调度
  - [x] 26.1 增量更新策略
  - [x] 26.2 任务调度

## 阶段九：API 网关与部署

- [x] 27. 实现 API 网关
  - [x] 27.1 路由配置
  - [x] 27.2 认证鉴权
  - [x] 27.3 日志监控

- [x] 28. 部署配置
  - [x] 28.1 云服务器部署
  - [x] 28.2 微信小程序配置

- [x] 29. 系统测试
  - [x] 29.1 单元测试
  - [x] 29.2 API 测试
  - [x] 29.3 集成测试

## 检查点

- [x] 检查点1：确保基础项目结构搭建完成
- [x] 检查点2：确保数据模型和存储层实现完成
- [x] 检查点3：确保用户服务和认证实现完成
- [x] 检查点4：确保数据服务 API 实现完成
- [x] 检查点5：确保 RAG 知识库构建完成
- [x] 检查点6：确保问答服务和 AI 集成完成
- [x] 检查点7：确保前端页面开发完成
- [x] 检查点8：确保数据采集模块完成
- [x] 检查点9：确保部署配置完成，系统可运行
- [ ] 检查点3：确保用户服务和认证实现完成
- [ ] 检查点4：确保数据服务 API 实现完成
- [ ] 检查点5：确保 RAG 知识库构建完成
- [ ] 检查点6：确保问答服务和 AI 集成完成
- [ ] 检查点7：确保前端页面开发完成
- [ ] 检查点8：确保数据采集模块完成
- [ ] 检查点9：确保部署配置完成，系统可运行
