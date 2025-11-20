# 模块: webhooks

## 接口: 创建Webhook端点
- 方法: POST
- 路径: /api/webhooks
- 说明: 为租户添加Webhook接收地址
- 认证: jwt
- 角色: admin

### 请求参数
- [body] secret (string) 必填:false 说明:签名密钥
- [body] tenant_id (int) 必填:true 说明:租户ID
- [body] url (string) 必填:true 说明:回调地址

### 响应
- 状态码: 201
- 返回说明: 返回创建的端点ID与URL

#### 返回参数
- msg (string)
- data (object)
  - id (integer) — 端点ID
  - url (string) — 回调URL
- code (int)

#### 示例数据
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

## 接口: 删除Webhook端点
- 方法: POST
- 路径: /api/webhooks/delete
- 说明: 删除指定端点
- 认证: jwt
- 角色: admin

### 请求参数
- [body] id (int) 必填:true 说明:端点ID

### 响应
- 状态码: 200
- 返回说明: 返回删除的端点ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 列出Webhook端点
- 方法: GET
- 路径: /api/webhooks
- 说明: 按租户查询Webhook端点
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200
- 返回说明: 返回端点列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - url (string)
    - status (string)
    - created_at (string)

#### 示例数据
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

## 接口: Webhook任务列表
- 方法: GET
- 路径: /api/webhooks/tasks
- 说明: 查询指定租户的Webhook任务
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200
- 返回说明: 返回任务列表与端点URL

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - url (string)
    - id (int)
    - event (string)
    - status (string)
    - retries (int)
    - next_try_at (string)
    - created_at (string)

#### 示例数据
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

## 接口: 重试Webhook任务
- 方法: POST
- 路径: /api/webhooks/tasks/retry
- 说明: 将任务状态设置为retry并立即调度
- 认证: jwt
- 角色: admin

### 请求参数
- [body] id (int) 必填:true 说明:任务ID

### 响应
- 状态码: 200
- 返回说明: 返回重试的任务ID，包含通用包装

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer) — 任务ID

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

