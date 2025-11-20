# 模块：pay

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付回调（支付宝）
- 接口方法：POST
- 接口路径：/api/pay/callback/alipay

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | amount | string | false |  |  |
| body | currency | string | false |  |  |
| body | order_no | string | true |  |  |
| body | status | string | true |  |  |
| body | txn_id | string | false |  |  |
| body | sig | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/callback/alipay" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "amount": "string",
  "currency": "string",
  "order_no": "string",
  "sig": "string",
  "status": "string",
  "txn_id": "string"
}
JSON
```
```json
{
  "amount": "string",
  "currency": "string",
  "order_no": "string",
  "sig": "string",
  "status": "string",
  "txn_id": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/callback/alipay"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - order_no (string)
  - status (string)

- 返回说明：返回订单号与支付状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "order_no": "string",
    "status": "string"
  },
  "msg": "string"
}
```

