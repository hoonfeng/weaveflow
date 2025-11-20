# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 订单导出
- 接口方法：GET
- 接口路径：/api/admin/orders/export
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/export?tenant_id=1&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/export" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - order_no (string)
  - tenant_id (int)
  - amount (number)
  - currency (string)
  - provider (string)
  - status (string)
  - ts (string)

- 返回说明：返回订单列表（JSON），用于下载
- 响应头：
  - Content-Type：application/json

### 示例数据
```json
{
  "item": {
    "amount": 1.23,
    "currency": "string",
    "order_no": "string",
    "provider": "string",
    "status": "string",
    "tenant_id": 123,
    "ts": "string"
  }
}
```

