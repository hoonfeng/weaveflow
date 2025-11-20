# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 订单详情
- 接口方法：GET
- 接口路径：/api/admin/orders/{order_no}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/{order_no}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/{order_no}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - provider (string)
    - status (string)
    - ts (string)
    - order_no (string)
    - tenant_id (int)
    - amount (number)
    - currency (string)

- 返回说明：返回订单详情记录

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "amount": 1.23,
      "currency": "string",
      "order_no": "string",
      "provider": "string",
      "status": "string",
      "tenant_id": 123,
      "ts": "string"
    }
  ],
  "msg": "string"
}
```

