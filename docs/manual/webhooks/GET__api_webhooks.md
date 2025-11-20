# 模块：webhooks

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 列出Webhook端点
- 接口方法：GET
- 接口路径：/api/webhooks
- 接口说明：按租户查询Webhook端点
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/webhooks?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/webhooks" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - url (string)
    - status (string)
    - created_at (string)

### 错误码
| Code | Message | Description |
|---|---|---|
| E_DB_UNAVAILABLE | 数据库不可用 | 查询失败 |

#### 错误响应示例
```json
{
  "code": "E_DB_UNAVAILABLE",
  "msg": "数据库不可用"
}
```

- 返回说明：返回端点列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "id": 123,
      "status": "string",
      "url": "string"
    }
  ],
  "msg": "string"
}
```

