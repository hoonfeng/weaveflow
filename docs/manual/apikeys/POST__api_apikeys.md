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

