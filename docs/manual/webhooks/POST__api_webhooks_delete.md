# 模块：webhooks

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 删除Webhook端点
- 接口方法：POST
- 接口路径：/api/webhooks/delete
- 接口说明：删除指定端点
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | id | int | true | 端点ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks/delete" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
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
curl -X POST "http://localhost:8080/api/webhooks/delete" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)

### 错误码
| Code | Message | Description |
|---|---|---|
| E_NOT_FOUND | 端点不存在 | ID 无效 |

#### 错误响应示例
```json
{
  "code": "E_NOT_FOUND",
  "msg": "端点不存在"
}
```

- 返回说明：返回删除的端点ID

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

