# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 变更订阅
- 接口方法：POST
- 接口路径：/api/admin/subscriptions/change
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true |  |  |
| body | plan_id | int | true |  |  |
| body | status | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions/change" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "plan_id": 1,
  "status": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "plan_id": 1,
  "status": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions/change" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)：订阅状态
  - tenant_id (int)：租户ID

- 返回说明：返回变更的租户ID与订阅状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

