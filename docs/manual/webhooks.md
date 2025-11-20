# 模块：webhooks

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建Webhook端点
- 接口方法：POST
- 接口路径：/api/webhooks
- 接口说明：为租户添加Webhook接收地址
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true | 租户ID |  |
| body | url | string | true | 回调地址 |  |
| body | secret | string | false | 签名密钥 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "secret": "string",
  "tenant_id": 1,
  "url": "string"
}
JSON
```
```json
{
  "secret": "string",
  "tenant_id": 1,
  "url": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：端点ID
  - url (string)：回调URL

### 错误码
| Code | Message | Description |
|---|---|---|
| E_ENDPOINT_EXISTS | 端点已存在 | URL 重复 |

#### 错误响应示例
```json
{
  "code": "E_ENDPOINT_EXISTS",
  "msg": "端点已存在"
}
```

- 返回说明：返回创建的端点ID与URL

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123,
    "url": "string"
  },
  "msg": "string"
}
```

## 删除Webhook端点
- 接口方法：POST
- 接口路径：/api/webhooks/delete
- 接口说明：删除指定端点
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | id | int | true | 端点ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks/delete" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "id": 1
}
JSON
```
```json
{
  "id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks/delete" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)

### 错误码
| Code | Message | Description |
|---|---|---|
| E_NOT_FOUND | 端点不存在 | ID 无效 |

#### 错误响应示例
```json
{
  "code": "E_NOT_FOUND",
  "msg": "端点不存在"
}
```

- 返回说明：返回删除的端点ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 列出Webhook端点
- 接口方法：GET
- 接口路径：/api/webhooks
- 接口说明：按租户查询Webhook端点
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/webhooks?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/webhooks" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - url (string)
    - status (string)
    - created_at (string)

### 错误码
| Code | Message | Description |
|---|---|---|
| E_DB_UNAVAILABLE | 数据库不可用 | 查询失败 |

#### 错误响应示例
```json
{
  "code": "E_DB_UNAVAILABLE",
  "msg": "数据库不可用"
}
```

- 返回说明：返回端点列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "id": 123,
      "status": "string",
      "url": "string"
    }
  ],
  "msg": "string"
}
```

## Webhook任务列表
- 接口方法：GET
- 接口路径：/api/webhooks/tasks
- 接口说明：查询指定租户的Webhook任务
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/webhooks/tasks?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/webhooks/tasks" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - event (string)
    - status (string)
    - retries (int)
    - next_try_at (string)
    - created_at (string)
    - url (string)
    - id (int)

### 错误码
| Code | Message | Description |
|---|---|---|
| E_DB_UNAVAILABLE | 数据库不可用 | 查询失败 |

#### 错误响应示例
```json
{
  "code": "E_DB_UNAVAILABLE",
  "msg": "数据库不可用"
}
```

- 返回说明：返回任务列表与端点URL

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "event": "string",
      "id": 123,
      "next_try_at": "string",
      "retries": 123,
      "status": "string",
      "url": "string"
    }
  ],
  "msg": "string"
}
```

## 重试Webhook任务
- 接口方法：POST
- 接口路径：/api/webhooks/tasks/retry
- 接口说明：将任务状态设置为retry并立即调度
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | id | int | true | 任务ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks/tasks/retry" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "id": 1
}
JSON
```
```json
{
  "id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks/tasks/retry" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：任务ID

### 错误码
| Code | Message | Description |
|---|---|---|
| E_TASK_NOT_FOUND | 任务不存在 | ID 无效 |

#### 错误响应示例
```json
{
  "code": "E_TASK_NOT_FOUND",
  "msg": "任务不存在"
}
```

- 返回说明：返回重试的任务ID，包含通用包装

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

