# 模块: portal

## 接口: 创建 API Key
- 方法: POST
- 路径: /api/portal/apikeys

### 请求参数
- [body] name (string) 必填:true 说明:
- [body] scopes (string) 必填:false 说明:
- [body] expired_at (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回创建的明文 key 与 secret

#### 返回参数
- data (object)
  - id (int)
  - key (string)
  - secret (string)
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123,
    "key": "string",
    "secret": "string"
  },
  "msg": "string"
}
```

## 接口: 删除 API Key
- 方法: DELETE
- 路径: /api/portal/apikeys/{id}

### 响应
- 状态码: 200
- 返回说明: 返回删除结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: API Key 列表
- 方法: GET
- 路径: /api/portal/apikeys

### 响应
- 状态码: 200
- 返回说明: 返回当前租户的 API Key 列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - status (string)
    - expired_at (string)
    - scopes (string)
    - created_at (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "expired_at": "string",
      "id": 123,
      "name": "string",
      "scopes": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 轮换 API Key 密钥
- 方法: POST
- 路径: /api/portal/apikeys/{id}/rotate

### 响应
- 状态码: 200
- 返回说明: 返回新 secret

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - secret (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "secret": "string"
  },
  "msg": "string"
}
```

## 接口: 启用/禁用 API Key
- 方法: POST
- 路径: /api/portal/apikeys/{id}/toggle

### 请求参数
- [body] status (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回更新后的状态

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string"
  },
  "msg": "string"
}
```

## 接口: 租户用户登录
- 方法: POST
- 路径: /api/portal/login

### 请求参数
- [body] username (string) 必填:true 说明:
- [body] password (string) 必填:true 说明:

### 响应
- 状态码: 0
- 返回说明: 返回令牌

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - token (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

## 接口: 角色权限列表
- 方法: GET
- 路径: /api/portal/roles/{role}/permissions

### 响应
- 状态码: 200
- 返回说明: 返回该租户下指定角色的权限码列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

## 接口: 设置角色权限
- 方法: POST
- 路径: /api/portal/roles/{role}/permissions

### 请求参数
- [body] codes (array) 必填:true 说明:权限码数组

### 响应
- 状态码: 200
- 返回说明: 返回设置后的权限数量

#### 返回参数
- data (object)
  - synced (int)
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "synced": 123
  },
  "msg": "string"
}
```

## 接口: 创建子账户并绑定到租户
- 方法: POST
- 路径: /api/portal/users

### 请求参数
- [body] email (string) 必填:true 说明:
- [body] password (string) 必填:true 说明:
- [body] role (string) 必填:true 说明:
- [body] username (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回创建绑定后的用户ID

#### 返回参数
- data (object)
  - uid (int)
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "uid": 123
  },
  "msg": "string"
}
```

## 接口: 删除子账户绑定
- 方法: DELETE
- 路径: /api/portal/users/{uid}

### 响应
- 状态码: 200
- 返回说明: 返回删除结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 子账户列表
- 方法: GET
- 路径: /api/portal/users

### 响应
- 状态码: 200
- 返回说明: 返回该租户下的用户列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - uid (int)
    - username (string)
    - email (string)
    - role (string)
    - created_at (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "email": "string",
      "role": "string",
      "uid": 123,
      "username": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 更新子账户角色
- 方法: PUT
- 路径: /api/portal/users/{uid}

### 请求参数
- [body] role (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回更新结果

#### 返回参数
- msg (string)
- data (object)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 租户自助注册
- 方法: POST
- 路径: /api/portal/tenants/register

### 请求参数
- [body] tenant_name (string) 必填:true 说明:
- [body] username (string) 必填:true 说明:
- [body] email (string) 必填:true 说明:
- [body] password (string) 必填:true 说明:
- [body] plan_id (int) 必填:false 说明:

### 响应
- 状态码: 0
- 返回说明: 返回注册成功后的令牌

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - token (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

