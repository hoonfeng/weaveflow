# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 订阅列表
- 接口方法：GET
- 接口路径：/api/admin/subscriptions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - status (string)：状态
    - id (int)：订阅ID
    - tenant_id (int)：租户ID
    - plan_id (int)：计划ID
    - plan_name (string)：计划名称

- 返回说明：返回订阅列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": 123,
      "plan_id": 123,
      "plan_name": "string",
      "status": "string",
      "tenant_id": 123
    }
  ],
  "msg": "string"
}
```

