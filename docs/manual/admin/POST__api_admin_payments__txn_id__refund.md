# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付退款
- 接口方法：POST
- 接口路径：/api/admin/payments/{txn_id}/refund
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | txn_id | string | true |  |  |
| body | amount | float | true |  |  |
| body | reason | string | false |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/payments/string/refund" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "amount": 1.23,
  "reason": "string"
}
JSON
```
```json
{
  "amount": 1.23,
  "reason": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/payments/{txn_id}/refund" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - amount (number)：金额
  - status (string)：状态
  - refund_id (string)：退款ID
  - order_no (string)：订单号

- 返回说明：返回退款信息

### 示例数据
```json
{
  "code": 123,
  "data": {
    "amount": 1.23,
    "order_no": "string",
    "refund_id": "string",
    "status": "string"
  },
  "msg": "string"
}
```

