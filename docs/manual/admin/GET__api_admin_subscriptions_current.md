# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 当前订阅查询
- 接口方法：GET
- 接口路径：/api/admin/subscriptions/current
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions/current?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions/current" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)
  - tenant_id (int)
  - plan_id (int)
  - status (string)
  - start_at (string)
  - end_at (string)

- 返回说明：返回指定租户的当前订阅

### 示例数据
```json
{
  "code": 123,
  "data": {
    "end_at": "string",
    "id": 123,
    "plan_id": 123,
    "start_at": "string",
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

