# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 租户订阅列表
- 接口方法：GET
- 接口路径：/api/admin/tenants/{id}/subscriptions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/subscriptions" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/subscriptions" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - next_billing_at (string)
    - id (int)
    - plan_id (int)
    - status (string)
    - start_at (string)
    - end_at (string)

- 返回说明：返回租户订阅记录

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "end_at": "string",
      "id": 123,
      "next_billing_at": "string",
      "plan_id": 123,
      "start_at": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

