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

- [ ] 4. 设计并创建 PostgreSQL 数据表
  - [ ] 4.1 创建用户相关表（users, student_profiles）
    - users 表：id, username, password_hash, phone, user_type, created_at
    - student_profiles 表：user_id, province, gaokao_year, score, rank, preferred_subjects, preferred_regions, interest_tags

  - [ ] 4.2 创建高校相关表（colleges）
    - id, name, province, city, district, type, level, established_year, departments, faculties, website, description, rankings, disciplines, statistics

  - [ ] 4.3 创建专业相关表（majors）
    - id, name, category, code, degree, duration, description, core_courses, employment

  - [ ] 4.4 创建录取分数表（admission_scores）
    - college_id, major_id, province, year, batch, category, science_score, arts_score, science_rank_min, science_rank_max, arts_rank_min, arts_rank_max

  - [ ] 4.5 创建对话历史表（conversations, messages）
    - conversations: id, user_id, created_at
    - messages: id, conversation_id, role, content, created_at

- [ ] 5. 创建 Elasticsearch 索引
  - [ ] 5.1 创建 colleges 索引
    - 设置 mappings（id, name, province, city, type, level, rankings, majors, description, embedding）
    - 配置 ik_max_word 中文分词器

  - [ ] 5.2 创建 majors 索引
    - 设置 mappings（id, name, category, degree, duration, employment, college_ids, embedding）

  - [ ] 5.3 创建 admission_scores 索引
    - 设置 mappings（college_id, major_id, province, year, batch, scores, ranks）

  - [ ] 5.4 创建 knowledge_base 索引（报考知识、规划知识）
    - 设置 mappings（type, title, content, tags, embedding）

- [ ] 6. 实现 Repository 层
  - [ ] 6.1 实现 UserRepository
    - Create, FindByID, FindByPhone, Update 方法

  - [ ] 6.2 实现 CollegeRepository
    - Create, FindByID, FindAll, Search, FindByConditions 方法

  - [ ] 6.3 实现 MajorRepository
    - Create, FindByID, FindAll, Search, FindByConditions 方法

  - [ ] 6.4 实现 ScoreRepository
    - Create, FindByCollege, FindByMajor, FindByConditions, FindProbability 方法

## 阶段三：用户服务与认证

- [ ] 7. 实现用户服务核心功能
  - [ ] 7.1 实现用户注册接口 POST /api/v1/user/register
    - 参数校验（username, password, phone, code）
    - 密码哈希存储
    - 生成 JWT Token
    - 返回 user_id 和 token

  - [ ] 7.2 实现用户登录接口 POST /api/v1/user/login
    - 支持密码登录和短信验证码登录
    - 验证密码/验证码
    - 生成 JWT Token
    - 返回用户信息和过期时间

  - [ ] 7.3 实现获取用户信息 GET /api/v1/user/profile
    - JWT Token 验证
    - 关联查询考生信息
    - 返回用户完整信息

  - [ ] 7.4 实现更新考生信息 PUT /api/v1/user/profile
    - 更新分数、位次、偏好等
    - 参数校验

  - [ ] 7.5 实现更新考生偏好 PUT /api/v1/user/preferences
    - 更新 preferred_subjects, preferred_regions, interest_tags

- [ ] 8. 实现 JWT 认证中间件
  - 解析和验证 JWT Token
  - 注入用户信息到 Context
  - 处理 Token 刷新逻辑

- [ ] 9. 实现短信验证码服务（模拟）
  - 生成6位验证码
  - 存储到 Redis（5分钟过期）
  - 验证验证码

## 阶段四：数据服务 API

- [ ] 10. 实现高校数据接口
  - [ ] 10.1 GET /api/v1/colleges - 高校列表查询
    - 分页、关键词搜索、条件过滤
    - 返回高校基本信息

  - [ ] 10.2 GET /api/v1/colleges/:id - 高校详情
    - 返回高校完整信息（排名、师资、学科等）

  - [ ] 10.3 GET /api/v1/colleges/:id/majors - 高校专业列表
    - 分页、学科门类过滤

- [ ] 11. 实现专业数据接口
  - [ ] 11.1 GET /api/v1/majors - 专业列表查询
    - 分页、关键词搜索、学科门类过滤

  - [ ] 11.2 GET /api/v1/majors/:id - 专业详情
    - 返回专业完整信息（核心课程、就业情况等）

- [ ] 12. 实现录取数据接口
  - [ ] 12.1 GET /api/v1/scores - 分数线查询
    - 按高校ID或专业ID查询
    - 支持年份、省份、批次、科类过滤

  - [ ] 12.2 GET /api/v1/scores/probability - 录取概率计算
    - 基于考生分数和位次
    - 分析历史录取数据
    - 返回录取概率和建议

- [ ] 13. 实现排名接口
  - [ ] 13.1 GET /api/v1/rankings/colleges - 学校排名查询
    - 按年份、排名类型、分类过滤

  - [ ] 13.2 GET /api/v1/rankings/majors - 专业排名查询
    - 按年份、学科门类过滤

## 阶段五：RAG 知识库构建

- [ ] 14. 实现数据向量化模块
  - [ ] 14.1 实现文本向量化接口
    - 调用国产大模型 Embedding API
    - 支持批量向量化

  - [ ] 14.2 实现高校/专业描述向量化
    - 生成文本描述
    - 调用向量化接口
    - 存储向量到 ES

- [ ] 15. 实现 RAG 检索模块
  - [ ] 15.1 实现混合检索策略
    - 关键词检索（BM25）
    - 向量相似度检索
    - 结果融合

  - [ ] 15.2 实现检索结果重排序
    - 使用交叉编码器重排序
    - 选择 Top-K 相关文档

  - [ ] 15.3 实现上下文组装
    - 限制上下文长度
    - 格式化检索结果

- [ ] 16. 实现知识库数据导入
  - [ ] 16.1 导入高校知识数据
    - 学校概况、历史沿革、师资力量、学科优势

  - [ ] 16.2 导入专业知识数据
    - 专业介绍、培养方案、课程设置、就业方向

  - [ ] 16.3 导入录取知识数据
    - 历年录取分数、录取位次、录取规则

  - [ ] 16.4 导入报考知识数据
    - 志愿填报技巧、冲稳保策略、专业选择方法

  - [ ] 16.5 导入规划知识数据
    - 职业规划建议、专业发展前景、行业分析

## 阶段六：问答服务与 AI 集成

- [ ] 17. 实现问答服务核心功能
  - [ ] 17.1 POST /api/v1/qa/chat - AI 问答接口
    - 解析用户问题，提取考生上下文
    - RAG 检索相关知识
    - 调用大模型 API 生成答案
    - 返回答案、来源、推荐

  - [ ] 17.2 GET /api/v1/qa/history - 对话历史
    - 返回用户历史对话

  - [ ] 17.3 POST /api/v1/qa/recommend - 志愿推荐
    - 基于考生分数和偏好
    - 生成冲稳保策略
    - 返回推荐院校和专业

  - [ ] 17.4 GET /api/v1/qa/colleges - 推荐院校
    - 根据考生情况推荐院校

  - [ ] 17.5 GET /api/v1/qa/majors - 推荐专业
    - 根据考生情况推荐专业

- [ ] 18. 实现 AI 大模型集成
  - [ ] 18.1 实现模型客户端
    - 支持文心一言/通义千问 API
    - 实现 Chat Completions 接口

  - [ ] 18.2 实现提示词模板
    - 系统提示词（角色设定）
    - 用户提示词（问题 + 上下文）
    -few-shot 示例

  - [ ] 18.3 实现答案后处理
    - 解析模型返回
    - 格式化输出
    - 错误处理

## 阶段七：前端开发

- [ ] 19. 实现前端基础页面
  - [ ] 19.1 登录/注册页面
    - 手机号登录
    - 账号密码登录
    - 验证码登录

  - [ ] 19.2 个人中心页面
    - 考生信息管理
    - 偏好设置
    - 高考信息录入

- [ ] 20. 实现高校/专业查询页面
  - [ ] 20.1 高校列表页
    - 搜索和过滤
    - 高校卡片展示
    - 分页加载

  - [ ] 20.2 高校详情页
    - 学校介绍、排名、师资
    - 专业设置
    - 录取分数线

  - [ ] 20.3 专业列表页
    - 搜索和过滤
    - 专业卡片展示

  - [ ] 20.4 专业详情页
    - 专业介绍
    - 核心课程
    - 就业前景

- [ ] 21. 实现智能问答页面
  - [ ] 21.1 AI 问答界面
    - 对话列表
    - 消息输入
    - 院校/专业推荐卡片

  - [ ] 21.2 志愿推荐界面
    - 考生信息输入
    - 冲稳保策略展示
    - 推荐结果列表

- [ ] 22. 实现分数线查询页面
  - [ ] 22.1 分数线查询
    - 按高校/专业查询
    - 年份、省份、批次过滤
    - 录取概率计算

- [ ] 23. 实现排名榜单页面
  - [ ] 23.1 学校排名
    - 综合排名、专业排名
    - 分类筛选

  - [ ] 23.2 专业排名
    - 学科门类筛选

## 阶段八：数据采集模块

- [ ] 24. 实现数据采集爬虫
  - [ ] 24.1 高校信息爬虫
    - 从高校官网/阳光高考采集
    - 学校名称、地区、类型、排名等

  - [ ] 24.2 专业信息爬虫
    - 专业名称、学科门类、学位类型

  - [ ] 24.3 录取分数爬虫
    - 历年录取分数、位次

- [ ] 25. 实现数据处理和导入
  - [ ] 25.1 数据清洗
    - 格式化、去重、校验

  - [ ] 25.2 数据存储
    - 导入 PostgreSQL
    - 导入 Elasticsearch

- [ ] 26. 实现采集调度
  - [ ] 26.1 增量更新策略
    - 定期增量更新
    - 全量更新策略

  - [ ] 26.2 任务调度
    - 定时任务配置
    - 采集状态监控

## 阶段九：API 网关与部署

- [ ] 27. 实现 API 网关
  - [ ] 27.1 路由配置
    - 路径路由
    - 服务发现

  - [ ] 27.2 认证鉴权
    - JWT Token 验证
    - 限流熔断

  - [ ] 27.3 日志监控
    - 请求日志
    - 错误监控

- [ ] 28. 部署配置
  - [ ] 28.1 云服务器部署
    - Docker Compose 编排
    - Nginx 配置

  - [ ] 28.2 微信小程序配置
    - 域名白名单
    - API 请求配置

- [ ] 29. 系统测试
  - [ ] 29.1 单元测试
    - Service 层测试
    - Repository 层测试

  - [ ] 29.2 API 测试
    - 接口功能测试
    - 边界条件测试

  - [ ] 29.3 集成测试
    - 前后端联调
    - RAG 检索测试

## 检查点

- [x] 检查点1：确保基础项目结构搭建完成
- [ ] 检查点2：确保数据模型和存储层实现完成
- [ ] 检查点3：确保用户服务和认证实现完成
- [ ] 检查点4：确保数据服务 API 实现完成
- [ ] 检查点5：确保 RAG 知识库构建完成
- [ ] 检查点6：确保问答服务和 AI 集成完成
- [ ] 检查点7：确保前端页面开发完成
- [ ] 检查点8：确保数据采集模块完成
- [ ] 检查点9：确保部署配置完成，系统可运行
