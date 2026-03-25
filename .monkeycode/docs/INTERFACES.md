# 高考志愿填报指导系统 - 接口文档

## 1. API概述

### 1.1 基本信息

- **基础URL**: `https://api.example.com/v1`
- **认证方式**: JWT Bearer Token
- **数据格式**: JSON
- **字符编码**: UTF-8

### 1.2 通用响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

| 字段 | 类型 | 说明 |
|-----|------|-----|
| code | int | 状态码，0表示成功 |
| message | string | 响应信息 |
| data | object | 响应数据 |

### 1.3 错误码定义

| 错误码 | 说明 |
|-------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1002 | 缺少必要参数 |
| 2001 | 用户未登录 |
| 2002 | Token无效 |
| 2003 | 账号已被禁用 |
| 3001 | 资源不存在 |
| 4001 | 服务器内部错误 |
| 4002 | 服务暂不可用 |

---

## 2. 用户模块接口

### 2.1 用户注册

```
POST /api/v1/user/register
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码(8-20位) |
| phone | string | 是 | 手机号 |
| code | string | 是 | 短信验证码 |
| user_type | string | 否 | 用户类型(student/parent)，默认student |

**请求示例**：
```json
{
  "username": "zhangsan",
  "password": "12345678",
  "phone": "13800138000",
  "code": "123456",
  "user_type": "student"
}
```

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user_id": 10001,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 2.2 用户登录

```
POST /api/v1/user/login
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| login_type | string | 是 | 登录方式(phone/password) |
| phone | string | 条件 | 手机号(phone方式必填) |
| password | string | 条件 | 密码(password方式必填) |
| code | string | 条件 | 验证码(phone方式必填) |

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user_id": 10001,
    "username": "zhangsan",
    "phone": "13800138000",
    "user_type": "student",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": "2025-06-26T12:00:00Z"
  }
}
```

### 2.3 获取用户信息

```
GET /api/v1/user/profile
```

**请求头**：
```
Authorization: Bearer <token>
```

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user_id": 10001,
    "username": "zhangsan",
    "phone": "13800138000",
    "user_type": "student",
    "created_at": "2025-03-01T08:00:00Z",
    "profile": {
      "province": "河南省",
      "gaokao_year": 2025,
      "score": 580,
      "rank": 52000,
      "preferred_subjects": ["计算机", "电子信息"],
      "preferred_regions": ["北京", "上海", "浙江"],
      "interest_tags": ["人工智能", "软件开发"]
    }
  }
}
```

### 2.4 更新考生信息

```
PUT /api/v1/user/profile
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| province | string | 否 | 省份 |
| gaokao_year | int | 否 | 高考年份 |
| score | int | 否 | 高考分数 |
| rank | int | 否 | 省内位次 |
| preferred_subjects | []string | 否 | 偏好专业 |
| preferred_regions | []string | 否 | 偏好地区 |
| interest_tags | []string | 否 | 兴趣标签 |

---

## 3. 高校数据接口

### 3.1 高校列表查询

```
GET /api/v1/colleges
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| page | int | 否 | 页码，默认1 |
| page_size | int | 否 | 每页条数，默认20 |
| keyword | string | 否 | 搜索关键词 |
| province | string | 否 | 所在省份 |
| type | string | 否 | 学校类型(综合/理工/师范...) |
| level | string | 否 | 办学层次(本科/专科) |
| rank_min | int | 否 | 排名最小值 |
| rank_max | int | 否 | 排名最大值 |
| min_score | int | 否 | 录取最低分(河南) |

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 1234,
    "page": 1,
    "page_size": 20,
    "list": [
      {
        "id": "10001",
        "name": "清华大学",
        "province": "北京",
        "city": "北京",
        "type": "综合",
        "level": "本科",
        "ranking": {
          "comprehensive": 1,
          "province": 1
        },
        "logo_url": "https://xxx.com/logo.png",
        "admission_score": {
          "science": 680,
          "arts": 660
        }
      }
    ]
  }
}
```

### 3.2 高校详情

```
GET /api/v1/colleges/{id}
```

**路径参数**：
- id: 高校ID

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "10001",
    "name": "清华大学",
    "province": "北京",
    "city": "北京",
    "district": "海淀区",
    "type": "综合",
    "level": "本科",
    "established_year": 1911,
    "departments": 20,
    "faculties": 15000,
    "website": "https://www.tsinghua.edu.cn",
    "description": "清华大学是中国著名高等学府...",
    "rankings": {
      "comprehensive": 1,
      "qs_world": 15,
      "nature_index": 1
    },
    "disciplines": ["计算机科学与技术", "材料科学与工程"],
    "statistics": {
      "enrollment_count": 3800,
      "employment_rate": 98.5,
      "average_salary": 18000
    }
  }
}
```

### 3.3 高校专业列表

```
GET /api/v1/colleges/{id}/majors
```

**路径参数**：
- id: 高校ID

**查询参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页条数 |
| category | string | 否 | 学科门类 |
| admission_year | int | 否 | 招生年份 |

---

## 4. 专业数据接口

### 4.1 专业列表查询

```
GET /api/v1/majors
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页条数 |
| keyword | string | 否 | 搜索关键词 |
| category | string | 否 | 学科门类 |
| degree | string | 否 | 学位类型 |
| salary_level | string | 否 | 薪资等级(high/medium/low) |

### 4.2 专业详情

```
GET /api/v1/majors/{id}
```

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "20001",
    "name": "计算机科学与技术",
    "category": "工学",
    "code": "080901",
    "degree": "工学学士",
    "duration": 4,
    "description": "培养具有良好科学素养...",
    "core_courses": ["数据结构", "算法设计", "操作系统"],
    "employment": {
      "rate": 95.5,
      "avg_salary": 15000,
      "top_industries": ["互联网", "金融科技", "人工智能"]
    },
    "ranking": {
      "overall": 1,
      "schools_count": 500
    }
  }
}
```

---

## 5. 录取数据接口

### 5.1 分数线查询

```
GET /api/v1/scores
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| college_id | string | 条件 | 高校ID(与major_id二选一) |
| major_id | string | 条件 | 专业ID(与college_id二选一) |
| province | string | 否 | 省份，默认河南省 |
| year | int | 否 | 年份，默认2024 |
| batch | string | 否 | 批次(本科一批/本科二批...) |
| category | string | 否 | 科类(理科/文科/综合) |

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "college_id": "10001",
        "college_name": "清华大学",
        "major_id": "20001",
        "major_name": "计算机科学与技术",
        "year": 2024,
        "batch": "本科一批",
        "category": "理科",
        "score": {
          "min": 680,
          "max": 700,
          "avg": 688
        },
        "rank": {
          "min": 100,
          "max": 50
        }
      }
    ]
  }
}
```

### 5.2 录取概率计算

```
GET /api/v1/scores/probability
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| college_id | string | 是 | 高校ID |
| major_id | string | 否 | 专业ID |
| score | int | 是 | 考生分数 |
| rank | int | 是 | 考生位次 |

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "college_name": "清华大学",
    "major_name": "计算机科学与技术",
    "probability": 65.5,
    "level": "冲",
    "analysis": "根据您的高考分数680分和位次120名...",
    "similar_scores": [
      {"year": 2024, "score": 678, "rank": 130, "result": "录取"},
      {"year": 2023, "score": 672, "rank": 150, "result": "录取"}
    ]
  }
}
```

---

## 6. 智能问答接口

### 6.1 AI问答

```
POST /api/v1/qa/chat
```

**请求头**：
```
Authorization: Bearer <token>
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| message | string | 是 | 用户问题 |
| conversation_id | string | 否 | 会话ID(新会话为空) |
| user_context | object | 否 | 用户上下文信息 |

**user_context结构**：
```json
{
  "score": 580,
  "rank": 52000,
  "preferred_subjects": ["计算机", "电子信息"],
  "preferred_regions": ["北京", "上海"],
  "interest_tags": ["人工智能"]
}
```

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "conversation_id": "conv_20250325_001",
    "message_id": "msg_001",
    "answer": "根据您的高考分数580分和省内位次52000名...",
    "sources": [
      {
        "type": "college",
        "id": "10005",
        "name": "北京邮电大学",
        "relevance": 0.95
      }
    ],
    "recommendations": {
      "colleges": [
        {
          "name": "北京邮电大学",
          "probability": 85,
          "reason": "您的分数往年录取概率较高"
        }
      ],
      "majors": [
        {
          "name": "计算机科学与技术",
          "match_score": 90
        }
      ]
    }
  }
}
```

### 6.2 志愿推荐

```
POST /api/v1/qa/recommend
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| score | int | 是 | 高考分数 |
| rank | int | 是 | 省内位次 |
| preferred_subjects | []string | 否 | 偏好专业 |
| preferred_regions | []string | 否 | 偏好地区 |
| batch | string | 否 | 目标批次 |
| plan_type | string | 否 | 志愿类型(平行/顺序) |

**响应示例**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "strategy": {
      "chong": [
        {
          "college_name": "北京航空航天大学",
          "major_name": "计算机科学与技术",
          "probability": 55,
          "score_diff": -15
        }
      ],
      "wen": [
        {
          "college_name": "北京邮电大学",
          "major_name": "计算机科学与技术",
          "probability": 78,
          "score_diff": 5
        }
      ],
      "bao": [
        {
          "college_name": "郑州大学",
          "major_name": "计算机科学与技术",
          "probability": 95,
          "score_diff": 30
        }
      ]
    },
    "analysis": "基于您的分数和位次，推荐..."
  }
}
```

---

## 7. 排名接口

### 7.1 学校排名查询

```
GET /api/v1/rankings/colleges
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| year | int | 否 | 年份，默认2024 |
| ranking_type | string | 否 | 排名类型(comprehensive/subject) |
| category | string | 否 | 分类(理工/综合/师范...) |
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页条数 |

### 7.2 专业排名查询

```
GET /api/v1/rankings/majors
```

**请求参数**：

| 参数名 | 类型 | 必填 | 说明 |
|-------|------|-----|------|
| year | int | 否 | 年份 |
| category | string | 否 | 学科门类 |
| page | int | 否 | 页码 |
| page_size | int | 否 | 每页条数 |
