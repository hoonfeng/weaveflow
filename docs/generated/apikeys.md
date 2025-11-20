# 模块: apikeys

## 接口: 创建API Key
- 方法: POST
- 路径: /api/apikeys
- 说明: 为指定租户生成API Key（只展示一次）
- 认证: jwt
- 角色: admin

### 请求参数
- [body] name (string) 必填:true 说明:名称
- [body] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 201
- 返回说明: 返回生成的 Key 与 Secret

#### 返回参数
- msg (string)
- data (object)
  - secret (string) — Secret 仅展示一次
  - key (string) — Key 前缀 ak_
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "key": "string",
    "secret": "string"
  },
  "msg": "string"
}
```

## 接口: 列出API Keys
- 方法: GET
- 路径: /api/apikeys
- 说明: 按租户列出所有API Key
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200
- 返回说明: 返回 API Key 列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (integer) — Key ID
    - name (string) — 名称
    - status (string) — 状态
    - expired_at (string) — 过期时间
    - created_at (string) — 创建时间

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
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 撤销API Key
- 方法: POST
- 路径: /api/apikeys/revoke
- 说明: 将密钥状态更新为撤销
- 认证: jwt
- 角色: admin

### 请求参数
- [body] id (int) 必填:true 说明:密钥ID

### 响应
- 状态码: 200
- 返回说明: 返回撤销的密钥ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer) — 密钥ID

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

