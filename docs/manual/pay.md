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

## 支付回调（微信）
- 接口方法：POST
- 接口路径：/api/pay/callback/wechat

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | order_no | string | true |  |  |
| body | status | string | true |  |  |
| body | txn_id | string | false |  |  |
| body | sig | string | true |  |  |
| body | amount | string | false |  |  |
| body | currency | string | false |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/callback/wechat" -H "Content-Type: application/json" -d @- << 'JSON'
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
curl -X POST "http://localhost:8080/api/pay/callback/wechat"
```

### 返回参数
- msg (string)
- data (object)
  - order_no (string)
  - status (string)
- code (int)

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

## 创建订单
- 接口方法：POST
- 接口路径：/api/pay/checkout
- 接口说明：为租户创建支付订单
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true | 租户ID |  |
| body | amount | float | true | 金额 |  |
| body | currency | string | false | 货币 |  |
| body | provider | string | true | 支付渠道 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/checkout" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "amount": 1.23,
  "currency": "string",
  "provider": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "amount": 1.23,
  "currency": "string",
  "provider": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/pay/checkout" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - order_no (string)：订单号
  - pay_url (string)：支付跳转链接

### 错误码
| Code | Message | Description |
|---|---|---|
| E_PROVIDER_UNSUPPORTED | 不支持的支付渠道 | 请选择 alipay/wechat/stripe |
| E_AMOUNT_INVALID | 金额不合法 | 金额必须大于0 |

#### 错误响应示例
```json
{
  "code": "E_PROVIDER_UNSUPPORTED",
  "msg": "不支持的支付渠道"
}
```

- 返回说明：返回创建的订单号与支付链接

### 示例数据
```json
{
  "code": 123,
  "data": {
    "order_no": "string",
    "pay_url": "string"
  },
  "msg": "string"
}
```

