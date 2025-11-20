# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建 API Key
- 接口方法：POST
- 接口路径：/api/portal/apikeys

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true |  |  |
| body | scopes | string | false |  |  |
| body | expired_at | string | false |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "expired_at": "string",
  "name": "string",
  "scopes": "string"
}
JSON
```
```json
{
  "expired_at": "string",
  "name": "string",
  "scopes": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)
  - key (string)
  - secret (string)

- 返回说明：返回创建的明文 key 与 secret

### 示例数据
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

## 删除 API Key
- 接口方法：DELETE
- 接口路径：/api/portal/apikeys/{id}

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/portal/apikeys/{id}"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/portal/apikeys/{id}"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回删除结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## API Key 列表
- 接口方法：GET
- 接口路径：/api/portal/apikeys

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/apikeys"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/apikeys"
```

### 返回参数
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

- 返回说明：返回当前租户的 API Key 列表

### 示例数据
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

## 轮换 API Key 密钥
- 接口方法：POST
- 接口路径：/api/portal/apikeys/{id}/rotate

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys/{id}/rotate"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys/{id}/rotate"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - secret (string)

- 返回说明：返回新 secret

### 示例数据
```json
{
  "code": 123,
  "data": {
    "secret": "string"
  },
  "msg": "string"
}
```

## 启用/禁用 API Key
- 接口方法：POST
- 接口路径：/api/portal/apikeys/{id}/toggle

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | status | string | true |  | enum=[active disabled] |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys/{id}/toggle" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "status": "string"
}
JSON
```
```json
{
  "status": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys/{id}/toggle"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)

- 返回说明：返回更新后的状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string"
  },
  "msg": "string"
}
```

## 租户用户登录
- 接口方法：POST
- 接口路径：/api/portal/login

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | password | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/login" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "password": "string",
  "username": "string"
}
JSON
```
```json
{
  "password": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/login"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - token (string)

- 返回说明：返回令牌

### 示例数据
```json
{
  "code": 123,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

## 角色权限列表
- 接口方法：GET
- 接口路径：/api/portal/roles/{role}/permissions

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/roles/{role}/permissions"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/roles/{role}/permissions"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (string)

- 返回说明：返回该租户下指定角色的权限码列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

## 设置角色权限
- 接口方法：POST
- 接口路径：/api/portal/roles/{role}/permissions

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | codes | array | true | 权限码数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/roles/{role}/permissions" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "codes": [
    "string"
  ]
}
JSON
```
```json
{
  "codes": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/roles/{role}/permissions"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - synced (int)

- 返回说明：返回设置后的权限数量

### 示例数据
```json
{
  "code": 123,
  "data": {
    "synced": 123
  },
  "msg": "string"
}
```

## 创建子账户并绑定到租户
- 接口方法：POST
- 接口路径：/api/portal/users

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | email | string | true |  |  |
| body | password | string | true |  |  |
| body | role | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/users" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "role": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "role": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/users"
```

### 返回参数
- msg (string)
- data (object)
  - uid (int)
- code (int)

- 返回说明：返回创建绑定后的用户ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "uid": 123
  },
  "msg": "string"
}
```

## 删除子账户绑定
- 接口方法：DELETE
- 接口路径：/api/portal/users/{uid}

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/portal/users/{uid}"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/portal/users/{uid}"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回删除结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 子账户列表
- 接口方法：GET
- 接口路径：/api/portal/users

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/users"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/users"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - role (string)
    - created_at (string)
    - uid (int)
    - username (string)
    - email (string)

- 返回说明：返回该租户下的用户列表

### 示例数据
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

## 更新子账户角色
- 接口方法：PUT
- 接口路径：/api/portal/users/{uid}

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | role | string | true |  |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/portal/users/{uid}" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "role": "string"
}
JSON
```
```json
{
  "role": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/portal/users/{uid}"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回更新结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 租户自助注册
- 接口方法：POST
- 接口路径：/api/portal/tenants/register

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | email | string | true |  |  |
| body | password | string | true |  |  |
| body | plan_id | int | false |  |  |
| body | tenant_name | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/tenants/register" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "plan_id": 1,
  "tenant_name": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "plan_id": 1,
  "tenant_name": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/tenants/register"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - token (string)

- 返回说明：返回注册成功后的令牌

### 示例数据
```json
{
  "code": 123,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

