# 模块：pay

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

