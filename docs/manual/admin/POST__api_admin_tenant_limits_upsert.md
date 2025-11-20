# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 设置租户限流
- 接口方法：POST
- 接口路径：/api/admin/tenant_limits/upsert
- 接口说明：设定租户级rps与burst
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | rps | float | true | 每秒请求数 |  |
| body | burst | int | true | 突发 |  |
| body | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenant_limits/upsert" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "burst": 1,
  "rps": 1.23,
  "tenant_id": 1
}
JSON
```
```json
{
  "burst": 1,
  "rps": 1.23,
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenant_limits/upsert" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - rps (number)
  - burst (int)
  - tenant_id (int)

- 返回说明：返回更新后的限流设置

### 示例数据
```json
{
  "code": 123,
  "data": {
    "burst": 123,
    "rps": 1.23,
    "tenant_id": 123
  },
  "msg": "string"
}
```

