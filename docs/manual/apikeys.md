# 模块：apikeys

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建API Key
- 接口方法：POST
- 接口路径：/api/apikeys
- 接口说明：为指定租户生成API Key（只展示一次）
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true | 租户ID |  |
| body | name | string | true | 名称 | minLen=1 |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/apikeys" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "name": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "name": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/apikeys" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - key (string)：Key 前缀 ak_
  - secret (string)：Secret 仅展示一次

- 返回说明：返回生成的 Key 与 Secret

### 示例数据
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

## 列出API Keys
- 接口方法：GET
- 接口路径：/api/apikeys
- 接口说明：按租户列出所有API Key
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/apikeys?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/apikeys" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (array)
  - item (object)
    - id (integer)：Key ID
    - name (string)：名称
    - status (string)：状态
    - expired_at (string)：过期时间
    - created_at (string)：创建时间
- code (int)
- msg (string)

- 返回说明：返回 API Key 列表

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
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 撤销API Key
- 接口方法：POST
- 接口路径：/api/apikeys/revoke
- 接口说明：将密钥状态更新为撤销
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | id | int | true | 密钥ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/apikeys/revoke" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
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
curl -X POST "http://localhost:8080/api/apikeys/revoke" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - id (integer)：密钥ID
- code (int)
- msg (string)

- 返回说明：返回撤销的密钥ID

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

