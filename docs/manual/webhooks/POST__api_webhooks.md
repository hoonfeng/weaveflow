# 模块：webhooks

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建Webhook端点
- 接口方法：POST
- 接口路径：/api/webhooks
- 接口说明：为租户添加Webhook接收地址
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true | 租户ID |  |
| body | url | string | true | 回调地址 |  |
| body | secret | string | false | 签名密钥 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "secret": "string",
  "tenant_id": 1,
  "url": "string"
}
JSON
```
```json
{
  "secret": "string",
  "tenant_id": 1,
  "url": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/webhooks" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：端点ID
  - url (string)：回调URL

### 错误码
| Code | Message | Description |
|---|---|---|
| E_ENDPOINT_EXISTS | 端点已存在 | URL 重复 |

#### 错误响应示例
```json
{
  "code": "E_ENDPOINT_EXISTS",
  "msg": "端点已存在"
}
```

- 返回说明：返回创建的端点ID与URL

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123,
    "url": "string"
  },
  "msg": "string"
}
```

