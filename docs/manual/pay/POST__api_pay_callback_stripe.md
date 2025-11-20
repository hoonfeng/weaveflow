# 模块：pay

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付回调（Stripe）
- 接口方法：POST
- 接口路径：/api/pay/callback/stripe
- 接口说明：支付成功/失败回调，更新订单并入账

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | order_no | string | true | 订单号 |  |
| body | status | string | true | 状态(succeeded/failed) |  |
| body | txn_id | string | false | 交易号 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/callback/stripe" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "order_no": "string",
  "status": "string",
  "txn_id": "string"
}
JSON
```
```json
{
  "order_no": "string",
  "status": "string",
  "txn_id": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/callback/stripe"
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

