# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付记录详情
- 接口方法：GET
- 接口路径：/api/admin/payments/{txn_id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments/{txn_id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments/{txn_id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (array)
  - item (object)
    - txn_id (string)
    - order_no (string)
    - provider (string)
    - status (string)
    - ts (string)
- code (int)
- msg (string)

- 返回说明：返回支付详情记录

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "order_no": "string",
      "provider": "string",
      "status": "string",
      "ts": "string",
      "txn_id": "string"
    }
  ],
  "msg": "string"
}
```

