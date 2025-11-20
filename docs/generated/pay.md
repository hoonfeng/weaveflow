# 模块: pay

## 接口: 支付回调（支付宝）
- 方法: POST
- 路径: /api/pay/callback/alipay

### 请求参数
- [body] amount (string) 必填:false 说明:
- [body] currency (string) 必填:false 说明:
- [body] order_no (string) 必填:true 说明:
- [body] status (string) 必填:true 说明:
- [body] txn_id (string) 必填:false 说明:
- [body] sig (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回订单号与支付状态

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - order_no (string)
  - status (string)

#### 示例数据
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

## 接口: 支付回调（Stripe）
- 方法: POST
- 路径: /api/pay/callback/stripe
- 说明: 支付成功/失败回调，更新订单并入账

### 请求参数
- [body] order_no (string) 必填:true 说明:订单号
- [body] status (string) 必填:true 说明:状态(succeeded/failed)
- [body] txn_id (string) 必填:false 说明:交易号

### 响应
- 状态码: 200
- 返回说明: 返回订单号与支付状态

#### 返回参数
- data (object)
  - order_no (string)
  - status (string)
- code (int)
- msg (string)

#### 示例数据
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

## 接口: 支付回调（微信）
- 方法: POST
- 路径: /api/pay/callback/wechat

### 请求参数
- [body] amount (string) 必填:false 说明:
- [body] currency (string) 必填:false 说明:
- [body] order_no (string) 必填:true 说明:
- [body] status (string) 必填:true 说明:
- [body] txn_id (string) 必填:false 说明:
- [body] sig (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回订单号与支付状态

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - order_no (string)
  - status (string)

#### 示例数据
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

## 接口: 创建订单
- 方法: POST
- 路径: /api/pay/checkout
- 说明: 为租户创建支付订单
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:租户ID
- [body] amount (float) 必填:true 说明:金额
- [body] currency (string) 必填:false 说明:货币
- [body] provider (string) 必填:true 说明:支付渠道

### 响应
- 状态码: 201
- 返回说明: 返回创建的订单号与支付链接

#### 返回参数
- msg (string)
- data (object)
  - order_no (string) — 订单号
  - pay_url (string) — 支付跳转链接
- code (int)

#### 示例数据
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

