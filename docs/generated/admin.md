# 模块: admin

## 接口: API Key 轮换
- 方法: POST
- 路径: /api/admin/apikeys/rotate
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:
- [body] name (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回新 Key 与 Secret

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - key (string)
  - secret (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "key": "string",
    "secret": "string"
  },
  "msg": "string"
}
```

## 接口: API Key 启用/禁用
- 方法: POST
- 路径: /api/admin/apikeys/toggle
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:
- [body] name (string) 必填:true 说明:
- [body] enabled (bool) 必填:true 说明:

### 响应
- 状态码: 0
- 返回说明: 返回更新后的状态

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - name (string)
  - status (string) — 状态(active|disabled)
  - tenant_id (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string",
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

## 接口: 依赖健康聚合
- 方法: GET
- 路径: /api/admin/health
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回数据库与插件运行时的健康统计

#### 返回参数
- code (int) — 返回码
- msg (string) — 返回信息
- data (object) — 健康数据
  - sql_main (integer) — 主库探测 1=ok
  - plugin_count (integer) — 插件数量
  - plugin_disabled_count (integer) — 被禁用插件数量

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "plugin_count": 123,
    "plugin_disabled_count": 123,
    "sql_main": 123
  },
  "msg": "string"
}
```

## 接口: Hooks 检视
- 方法: POST
- 路径: /api/admin/hooks/inspect
- 说明: 查看当前接口的策略链
- 认证: jwt
- 角色: admin

### 请求参数
- [body] method (string) 必填:false 说明:
- [body] path (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回策略链与匹配结果

#### 返回参数
- data (object)
  - chain (array) — 策略链顺序
  - match (object) — 匹配信息
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "chain": [
      "string"
    ],
    "match": {}
  },
  "msg": "string"
}
```

## 接口: 订单详情
- 方法: GET
- 路径: /api/admin/orders/{order_no}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回订单详情记录

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - ts (string)
    - order_no (string)
    - tenant_id (int)
    - amount (number)
    - currency (string)
    - provider (string)
    - status (string)

#### 示例数据
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

## 接口: 订单导出
- 方法: GET
- 路径: /api/admin/orders/export
- 认证: jwt
- 角色: admin

### 请求参数
- [query] provider (string) 必填:false 说明:
- [query] status (string) 必填:false 说明:
- [query] tenant_id (int) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回订单列表（JSON），用于下载
- 响应头:
  - Content-Type: application/json

#### 返回参数
- item (object)
  - status (string)
  - ts (string)
  - order_no (string)
  - tenant_id (int)
  - amount (number)
  - currency (string)
  - provider (string)

#### 示例数据
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

## 接口: 订单导出CSV
- 方法: GET
- 路径: /api/admin/orders/export.csv
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:false 说明:
- [query] provider (string) 必填:false 说明:
- [query] status (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回 CSV 文本内容
- 响应头:
  - Content-Type: text/csv

#### 示例数据
```
order_no,tenant_id,amount,currency,provider,status,ts
ord_001,1,99.9,CNY,alipay,succeeded,2025-11-01T10:00:00Z
```

## 接口: 订单列表
- 方法: GET
- 路径: /api/admin/orders
- 认证: jwt
- 角色: admin

### 请求参数
- [query] page_size (int) 必填:false 说明:
- [query] tenant_id (int) 必填:false 说明:
- [query] provider (string) 必填:false 说明:
- [query] status (string) 必填:false 说明:
- [query] page (int) 必填:false 说明:
- [query] size (int) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回订单列表，包含金额与状态

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - order_no (string) — 订单号
    - amount (number) — 金额
    - status (string) — 状态

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "amount": 1.23,
      "order_no": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 概览统计
- 方法: GET
- 路径: /api/admin/overview
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回基础统计（示例字段）

#### 返回参数
- total_users (int)
- total_orders (int)

#### 示例数据
```json
{
  "total_orders": 123,
  "total_users": 123
}
```

## 接口: 支付统计
- 方法: GET
- 路径: /api/admin/overview/payments_stats
- 说明: 返回支付状态与渠道分布
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: <nil>

#### 返回参数
- by_status (array)
- by_provider (array)

#### 示例数据
```json
{
  "by_provider": [
    "string"
  ],
  "by_status": [
    "string"
  ]
}
```

## 接口: 概览趋势
- 方法: GET
- 路径: /api/admin/overview/trends
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回近12月的用户与订单趋势

#### 返回参数
- users (array)
- orders (array)

#### 示例数据
```json
{
  "orders": [
    "string"
  ],
  "users": [
    "string"
  ]
}
```

## 接口: 创建支付通道
- 方法: POST
- 路径: /api/admin/pay/providers
- 认证: jwt
- 角色: admin

### 请求参数
- [body] enabled (bool) 必填:false 说明:是否启用
- [body] config_json (string) 必填:false 说明:JSON 配置
- [body] code (string) 必填:true 说明:渠道代码如 wechat/alipay
- [body] name (string) 必填:true 说明:渠道名称

### 响应
- 状态码: 200
- 返回说明: 返回创建后的记录

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)
  - code (string)
  - name (string)
  - enabled (bool)
  - config_json (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "code": "string",
    "config_json": "string",
    "enabled": true,
    "id": 123,
    "name": "string"
  },
  "msg": "string"
}
```

## 接口: 删除支付通道
- 方法: DELETE
- 路径: /api/admin/pay/providers/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200

#### 示例数据
```
map[]
```

## 接口: 支付通道详情
- 方法: GET
- 路径: /api/admin/pay/providers/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200

#### 示例数据
```
{{ row.0 }}
```

## 接口: 支付通道列表
- 方法: GET
- 路径: /api/admin/pay/providers
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回支付通道配置列表

#### 返回参数
- item (object)
  - enabled (bool)
  - config_json (string)
  - id (int)
  - code (string)
  - name (string)

#### 示例数据
```json
{
  "item": {
    "code": "string",
    "config_json": "string",
    "enabled": true,
    "id": 123,
    "name": "string"
  }
}
```

## 接口: 启用/禁用支付通道
- 方法: POST
- 路径: /api/admin/pay/providers/{id}/toggle
- 认证: jwt
- 角色: admin

### 请求参数
- [path] id (int) 必填:true 说明:
- [body] enabled (bool) 必填:true 说明:

### 响应
- 状态码: 200

#### 示例数据
```
map[enabled:{{ enabled }} id:{{ id }}]
```

## 接口: 更新支付通道配置
- 方法: PUT
- 路径: /api/admin/pay/providers/{id}
- 认证: jwt
- 角色: admin

### 请求参数
- [path] id (int) 必填:true 说明:
- [body] name (string) 必填:false 说明:
- [body] enabled (bool) 必填:false 说明:
- [body] config_json (string) 必填:false 说明:

### 响应
- 状态码: 200

#### 示例数据
```
{{ row.0 }}
```

## 接口: 支付对账
- 方法: POST
- 路径: /api/admin/pay/reconcile
- 认证: jwt
- 角色: admin

### 请求参数
- [body] start (string) 必填:true 说明:开始日期YYYY-MM-DD
- [body] end (string) 必填:true 说明:结束日期YYYY-MM-DD

### 响应
- 状态码: 200

#### 示例数据
```
{{ rows }}
```

## 接口: 支付记录详情
- 方法: GET
- 路径: /api/admin/payments/{txn_id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回支付详情记录

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - order_no (string)
    - provider (string)
    - status (string)
    - ts (string)
    - txn_id (string)

#### 示例数据
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

## 接口: 支付记录导出
- 方法: GET
- 路径: /api/admin/payments/export
- 认证: jwt
- 角色: admin

### 请求参数
- [query] order_no (string) 必填:false 说明:
- [query] provider (string) 必填:false 说明:
- [query] status (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回支付记录列表（JSON），用于下载
- 响应头:
  - Content-Type: application/json

#### 返回参数
- item (object)
  - status (string)
  - ts (string)
  - txn_id (string)
  - order_no (string)
  - provider (string)

#### 示例数据
```json
{
  "item": {
    "order_no": "string",
    "provider": "string",
    "status": "string",
    "ts": "string",
    "txn_id": "string"
  }
}
```

## 接口: 支付记录导出CSV
- 方法: GET
- 路径: /api/admin/payments/export.csv
- 认证: jwt
- 角色: admin

### 请求参数
- [query] provider (string) 必填:false 说明:
- [query] status (string) 必填:false 说明:
- [query] order_no (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回 CSV 文本内容
- 响应头:
  - Content-Type: text/csv

#### 示例数据
```
order_no,provider,status,txn_id,ts
ord_001,alipay,succeeded,tx_001,2025-11-01T10:00:00Z
```

## 接口: 支付记录列表
- 方法: GET
- 路径: /api/admin/payments
- 认证: jwt
- 角色: admin

### 请求参数
- [query] page_size (int) 必填:false 说明:
- [query] order_no (string) 必填:false 说明:
- [query] provider (string) 必填:false 说明:
- [query] status (string) 必填:false 说明:
- [query] page (int) 必填:false 说明:
- [query] size (int) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回支付记录列表，包含交易ID与状态

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - ts (string) — 时间
    - txn_id (string) — 交易ID
    - order_no (string) — 订单号
    - provider (string) — 支付渠道
    - status (string) — 状态

#### 示例数据
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

## 接口: 支付退款
- 方法: POST
- 路径: /api/admin/payments/{txn_id}/refund
- 认证: jwt
- 角色: admin

### 请求参数
- [path] txn_id (string) 必填:true 说明:
- [body] amount (float) 必填:true 说明:
- [body] reason (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回退款信息

#### 返回参数
- data (object)
  - refund_id (string) — 退款ID
  - order_no (string) — 订单号
  - amount (number) — 金额
  - status (string) — 状态
- code (int)
- msg (string)

#### 示例数据
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

## 接口: 创建权限
- 方法: POST
- 路径: /api/admin/permissions
- 认证: jwt
- 角色: admin

### 请求参数
- [body] code (string) 必填:true 说明:权限代码
- [body] name (string) 必填:true 说明:权限名称
- [body] description (string) 必填:false 说明:描述

### 响应
- 状态码: 201
- 返回说明: 返回创建的权限代码

#### 返回参数
- code (string)

#### 示例数据
```json
{
  "code": "string"
}
```

## 接口: 删除权限
- 方法: DELETE
- 路径: /api/admin/permissions/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回删除的权限ID

#### 返回参数
- id (int)

#### 示例数据
```json
{
  "id": 123
}
```

## 接口: 权限列表
- 方法: GET
- 路径: /api/admin/permissions
- 说明: 返回权限代码、名称与描述
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回权限列表

#### 返回参数
- item (object)
  - id (int)
  - code (string)
  - name (string)
  - description (string)

#### 示例数据
```json
{
  "item": {
    "code": "string",
    "description": "string",
    "id": 123,
    "name": "string"
  }
}
```

## 接口: 角色权限列表
- 方法: GET
- 路径: /api/admin/role_permissions
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回角色拥有的权限

#### 返回参数
- item (object)
  - perm_id (int)
  - code (string)
  - name (string)
  - role_id (int)

#### 示例数据
```json
{
  "item": {
    "code": "string",
    "name": "string",
    "perm_id": 123,
    "role_id": 123
  }
}
```

## 接口: 设置角色权限
- 方法: POST
- 路径: /api/admin/roles/{id}/permissions
- 认证: jwt
- 角色: admin

### 请求参数
- [path] id (int) 必填:true 说明:
- [body] perm_ids (array) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回设置后的权限数量

#### 返回参数
- role_id (int)
- count (int)

#### 示例数据
```json
{
  "count": 123,
  "role_id": 123
}
```

## 接口: 同步权限声明入库
- 方法: POST
- 路径: /api/admin/permissions/sync
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回同步的数量

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - scanned (int)
  - synced (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "scanned": 123,
    "synced": 123
  },
  "msg": "string"
}
```

## 接口: 更新权限
- 方法: PUT
- 路径: /api/admin/permissions/{id}
- 认证: jwt
- 角色: admin

### 请求参数
- [body] code (string) 必填:false 说明:权限代码
- [body] name (string) 必填:false 说明:权限名称
- [body] description (string) 必填:false 说明:描述

### 响应
- 状态码: 200
- 返回说明: 返回更新的权限ID

#### 返回参数
- id (int)

#### 示例数据
```json
{
  "id": 123
}
```

## 接口: 创建套餐
- 方法: POST
- 路径: /api/admin/plans
- 认证: jwt
- 角色: admin

### 请求参数
- [body] name (string) 必填:true 说明:
- [body] monthly_quota (int) 必填:false 说明:
- [body] price (number) 必填:true 说明:

### 响应
- 状态码: 201
- 返回说明: 返回创建的套餐名与价格

#### 返回参数
- msg (string)
- data (object)
  - name (string)
  - monthly_quota (int)
  - price (number)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "monthly_quota": 123,
    "name": "string",
    "price": 1.23
  },
  "msg": "string"
}
```

## 接口: 删除套餐
- 方法: DELETE
- 路径: /api/admin/plans/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回删除的套餐ID

#### 返回参数
- msg (string)
- data (object)
  - id (int)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 套餐列表
- 方法: GET
- 路径: /api/admin/plans
- 说明: 返回套餐列表，包含配额与价格
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回套餐列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - monthly_quota (int)
    - price (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": 123,
      "monthly_quota": 123,
      "name": "string",
      "price": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 更新套餐
- 方法: PUT
- 路径: /api/admin/plans/{id}
- 认证: jwt
- 角色: admin

### 请求参数
- [body] name (string) 必填:false 说明:
- [body] monthly_quota (int) 必填:false 说明:
- [body] price (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回更新的套餐ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 动态添加插件
- 方法: POST
- 路径: /api/admin/plugins/add
- 认证: jwt
- 角色: admin

### 请求参数
- [body] timeout (string) 必填:false 说明:超时时间，如300ms
- [body] queueSize (int) 必填:false 说明:队列大小
- [body] functions (array) 必填:true 说明:函数列表
- [body] name (string) 必填:true 说明:插件名称
- [body] windows (string) 必填:true 说明:Windows可执行文件路径
- [body] unix (string) 必填:false 说明:Unix可执行文件路径
- [body] instances (int) 必填:false 说明:实例数量

### 响应
- 状态码: 200
- 返回说明: 返回添加结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 内置插件列表
- 方法: GET
- 路径: /api/admin/plugins/builtin
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回内置插件能力列表

#### 返回参数
- item (object)
  - name (string)
  - desc (string)

#### 示例数据
```json
{
  "item": {
    "desc": "string",
    "name": "string"
  }
}
```

## 接口: 禁用插件
- 方法: POST
- 路径: /api/admin/plugins/disable
- 认证: jwt
- 角色: admin

### 请求参数
- [body] names (array) 必填:true 说明:插件名称数组

### 响应
- 状态码: 200
- 返回说明: 返回禁用结果

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - ok (boolean) — 是否成功

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "ok": true
  },
  "msg": "string"
}
```

## 接口: 启用插件
- 方法: POST
- 路径: /api/admin/plugins/enable
- 认证: jwt
- 角色: admin

### 请求参数
- [body] names (array) 必填:true 说明:插件名称数组

### 响应
- 状态码: 200
- 返回说明: 返回启用结果

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - ok (boolean) — 是否成功

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "ok": true
  },
  "msg": "string"
}
```

## 接口: 外置插件列表
- 方法: GET
- 路径: /api/admin/plugins/external
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回对外提供API服务的插件列表

#### 返回参数
- item (object)
  - name (string)
  - endpoints (array)

#### 示例数据
```json
{
  "item": {
    "endpoints": [
      "string"
    ],
    "name": "string"
  }
}
```

## 接口: 内置插件列表
- 方法: GET
- 路径: /api/admin/plugins/internal
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回系统必要功能的组成插件列表

#### 返回参数
- item (object)
  - name (string)
  - endpoints (array)

#### 示例数据
```json
{
  "item": {
    "endpoints": [
      "string"
    ],
    "name": "string"
  }
}
```

## 接口: 动态移除插件
- 方法: POST
- 路径: /api/admin/plugins/remove
- 认证: jwt
- 角色: admin

### 请求参数
- [body] names (array) 必填:true 说明:插件名称数组

### 响应
- 状态码: 200
- 返回说明: 返回移除结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 重启插件
- 方法: POST
- 路径: /api/admin/plugins/restart
- 认证: jwt
- 角色: admin

### 请求参数
- [body] names (array) 必填:true 说明:插件名称数组

### 响应
- 状态码: 200
- 返回说明: 返回重启结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 启动插件
- 方法: POST
- 路径: /api/admin/plugins/start
- 认证: jwt
- 角色: admin

### 请求参数
- [body] names (array) 必填:true 说明:插件名称数组

### 响应
- 状态码: 200
- 返回说明: 返回启动结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 插件状态
- 方法: GET
- 路径: /api/admin/plugins/status
- 说明: 返回当前已注册插件的状态
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: <nil>

#### 返回参数
- item (object)
  - name (string)
  - enabled (boolean)
  - running (boolean)
  - functions (array)

#### 示例数据
```json
{
  "item": {
    "enabled": true,
    "functions": [
      "string"
    ],
    "name": "string",
    "running": true
  }
}
```

## 接口: 停止插件
- 方法: POST
- 路径: /api/admin/plugins/stop
- 认证: jwt
- 角色: admin

### 请求参数
- [body] names (array) 必填:true 说明:插件名称数组

### 响应
- 状态码: 200
- 返回说明: 返回停止结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 创建角色
- 方法: POST
- 路径: /api/admin/roles
- 认证: jwt
- 角色: admin

### 请求参数
- [body] name (string) 必填:true 说明:角色名称

### 响应
- 状态码: 201
- 返回说明: 返回创建的角色名

#### 返回参数
- data (object)
  - name (string) — 角色名
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string"
  },
  "msg": "string"
}
```

## 接口: 删除角色
- 方法: DELETE
- 路径: /api/admin/roles/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回删除的角色ID

#### 返回参数
- data (object)
  - id (integer) — 角色ID
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 角色列表
- 方法: GET
- 路径: /api/admin/roles
- 说明: 返回全部角色名称与ID
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回角色列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (integer) — 角色ID
    - name (string) — 角色名

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": 123,
      "name": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 更新角色
- 方法: PUT
- 路径: /api/admin/roles/{id}
- 认证: jwt
- 角色: admin

### 请求参数
- [body] name (string) 必填:true 说明:角色名称

### 响应
- 状态码: 200
- 返回说明: 返回更新的角色ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer) — 角色ID

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 取消订阅
- 方法: POST
- 路径: /api/admin/subscriptions/cancel
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回取消订阅的租户ID与状态

#### 返回参数
- msg (string)
- data (object)
  - status (string) — 订阅状态
  - tenant_id (int) — 租户ID
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

## 接口: 变更订阅
- 方法: POST
- 路径: /api/admin/subscriptions/change
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:
- [body] plan_id (int) 必填:true 说明:
- [body] status (string) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回变更的租户ID与订阅状态

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - tenant_id (int) — 租户ID
  - status (string) — 订阅状态

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

## 接口: 创建订阅
- 方法: POST
- 路径: /api/admin/subscriptions
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:
- [body] plan_id (int) 必填:true 说明:

### 响应
- 状态码: 201
- 返回说明: 返回租户订阅创建结果

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - tenant_id (int)
  - plan_id (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "plan_id": 123,
    "tenant_id": 123
  },
  "msg": "string"
}
```

## 接口: 当前订阅查询
- 方法: GET
- 路径: /api/admin/subscriptions/current
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回指定租户的当前订阅

#### 返回参数
- data (object)
  - end_at (string)
  - id (int)
  - tenant_id (int)
  - plan_id (int)
  - status (string)
  - start_at (string)
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "end_at": "string",
    "id": 123,
    "plan_id": 123,
    "start_at": "string",
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

## 接口: 订阅列表
- 方法: GET
- 路径: /api/admin/subscriptions
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回订阅列表

#### 返回参数
- msg (string)
- data (array)
  - item (object)
    - tenant_id (int) — 租户ID
    - plan_id (int) — 计划ID
    - plan_name (string) — 计划名称
    - status (string) — 状态
    - id (int) — 订阅ID
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": 123,
      "plan_id": 123,
      "plan_name": "string",
      "status": "string",
      "tenant_id": 123
    }
  ],
  "msg": "string"
}
```

## 接口: 接口规范检查
- 方法: GET
- 路径: /api/admin/lint
- 说明: 返回当前已加载接口的规范问题列表
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回接口配置的检查结果列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - level (string) — 级别(error|warn|info)
    - message (string) — 描述
    - path (string) — 位置

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "level": "string",
      "message": "string",
      "path": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 模型迁移
- 方法: POST
- 路径: /api/admin/migrate
- 说明: 按模型配置对齐数据库结构
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回迁移的模型表名列表

#### 返回参数
- code (int)
- msg (string)
- data (array) — 表名列表
  - item (string) — 表名

#### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

## 接口: 清理HMAC非重放随机值
- 方法: POST
- 路径: /api/admin/nonces/purge
- 说明: 删除早于TTL的nonce记录
- 认证: jwt
- 角色: admin

### 请求参数
- [query] ttlSeconds (int) 必填:false 说明:TTL秒，默认900

### 响应
- 状态码: 200

#### 示例数据
```
map[ttlSeconds:{{ ttlSeconds }}]
```

## 接口: 队列状态
- 方法: GET
- 路径: /api/admin/queue/status
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回队列运行状态

#### 返回参数
- data (object)
  - workers (int) — 工作线程数量
  - pending (int) — 等待任务数
  - running (int) — 运行中任务数
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "pending": 123,
    "running": 123,
    "workers": 123
  },
  "msg": "string"
}
```

## 接口: 重载配置
- 方法: POST
- 路径: /api/admin/reload
- 说明: 重新加载项目与接口配置并返回差异
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回重载差异

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - removed (array) — 移除端点
  - updated (array) — 更新端点
  - added (array) — 新增端点

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "added": [
      "string"
    ],
    "removed": [
      "string"
    ],
    "updated": [
      "string"
    ]
  },
  "msg": "string"
}
```

## 接口: 远程调用示例（本地调试）
- 方法: GET
- 路径: /api/admin/remote/debug
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回插件运行时状态

#### 返回参数
- msg (string)
- data (object) — 插件状态原样透传
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 远程调用测试（熔断回退）
- 方法: GET
- 路径: /api/admin/remote/test
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回远程调用结果或回退内容

#### 返回参数
- code (int)
- msg (string)
- data (object) — 远程响应或fallback

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 租户API Key列表
- 方法: GET
- 路径: /api/admin/tenants/{id}/apikeys
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回租户的 API Key

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - status (string)
    - expired_at (string)
    - scopes (string)
    - created_at (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "expired_at": "string",
      "id": 123,
      "name": "string",
      "scopes": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 创建租户
- 方法: POST
- 路径: /api/admin/tenants
- 认证: jwt
- 角色: admin

### 请求参数
- [body] domain (string) 必填:false 说明:域名
- [body] contact_email (string) 必填:true 说明:联系邮箱
- [body] description (string) 必填:false 说明:描述
- [body] status (string) 必填:false 说明:状态
- [body] name (string) 必填:true 说明:租户名称

### 响应
- 状态码: 201
- 返回说明: 返回创建的租户名和状态

#### 返回参数
- data (object)
  - name (string) — 租户名
  - status (string) — 状态
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string",
    "status": "string"
  },
  "msg": "string"
}
```

## 接口: 删除租户
- 方法: DELETE
- 路径: /api/admin/tenants/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回删除的租户ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer) — 租户ID

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 租户详情
- 方法: GET
- 路径: /api/admin/tenants/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回单个租户详情

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - domain (string)
  - contact_email (string)
  - status (string)
  - description (string)
  - created_at (string)
  - id (int)
  - name (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "contact_email": "string",
    "created_at": "string",
    "description": "string",
    "domain": "string",
    "id": 123,
    "name": "string",
    "status": "string"
  },
  "msg": "string"
}
```

## 接口: 查询租户限流
- 方法: GET
- 路径: /api/admin/tenant_limits
- 说明: 查询指定租户的限流设置
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200
- 返回说明: 返回租户限流设置列表（通常为单条）

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - tenant_id (int) — 租户ID
    - rps (number) — 每秒请求限制
    - burst (int) — 突发容量
    - updated_at (string) — 更新时间

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "burst": 123,
      "rps": 1.23,
      "tenant_id": 123,
      "updated_at": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 设置租户限流
- 方法: POST
- 路径: /api/admin/tenant_limits/upsert
- 说明: 设定租户级rps与burst
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:租户ID
- [body] rps (float) 必填:true 说明:每秒请求数
- [body] burst (int) 必填:true 说明:突发

### 响应
- 状态码: 200
- 返回说明: 返回更新后的限流设置

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - tenant_id (int)
  - rps (number)
  - burst (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "burst": 123,
    "rps": 1.23,
    "tenant_id": 123
  },
  "msg": "string"
}
```

## 接口: 租户列表
- 方法: GET
- 路径: /api/admin/tenants
- 说明: 分页返回租户列表（示例为全部返回）
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回租户列表

#### 返回参数
- data (array)
  - item (object)
    - status (string) — 状态
    - created_at (string) — 创建时间
    - id (integer) — 租户ID
    - name (string) — 租户名
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "id": 123,
      "name": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 租户订阅列表
- 方法: GET
- 路径: /api/admin/tenants/{id}/subscriptions
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回租户订阅记录

#### 返回参数
- data (array)
  - item (object)
    - id (int)
    - plan_id (int)
    - status (string)
    - start_at (string)
    - end_at (string)
    - next_billing_at (string)
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "end_at": "string",
      "id": 123,
      "next_billing_at": "string",
      "plan_id": 123,
      "start_at": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 更新租户
- 方法: PUT
- 路径: /api/admin/tenants/{id}
- 认证: jwt
- 角色: admin

### 请求参数
- [body] name (string) 必填:false 说明:租户名称
- [body] domain (string) 必填:false 说明:域名
- [body] contact_email (string) 必填:false 说明:联系邮箱
- [body] description (string) 必填:false 说明:描述
- [body] status (string) 必填:false 说明:状态

### 响应
- 状态码: 200
- 返回说明: 返回更新的租户ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer) — 租户ID

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 租户月度用量统计
- 方法: GET
- 路径: /api/admin/tenants/{id}/usage/monthly
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回最近12个月按月统计的调用次数

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - ym (string)
    - cnt (int)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "cnt": 123,
      "ym": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 绑定租户成员
- 方法: POST
- 路径: /api/admin/tenants/{id}/users
- 认证: jwt
- 角色: admin

### 请求参数
- [body] user_id (int) 必填:true 说明:用户ID
- [body] role (string) 必填:true 说明:成员角色

### 响应
- 状态码: 200
- 返回说明: 返回绑定结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 租户成员列表
- 方法: GET
- 路径: /api/admin/tenants/{id}/users
- 说明: 返回指定租户的成员列表
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回租户成员及其角色

#### 返回参数
- msg (string)
- data (array)
  - item (object)
    - email (string)
    - role (string)
    - created_at (string)
    - uid (int)
    - username (string)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "email": "string",
      "role": "string",
      "uid": 123,
      "username": "string"
    }
  ],
  "msg": "string"
}
```

## 接口: 解绑租户成员
- 方法: DELETE
- 路径: /api/admin/tenants/{id}/users/{uid}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回解绑结果

#### 返回参数
- code (int)
- msg (string)
- data (object)

#### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 接口: 创建用户
- 方法: POST
- 路径: /api/admin/users
- 认证: jwt
- 角色: admin

### 请求参数
- [body] username (string) 必填:true 说明:
- [body] email (string) 必填:false 说明:
- [body] password (string) 必填:true 说明:

### 响应
- 状态码: 201
- 返回说明: 返回创建的用户名，包含通用包装

#### 返回参数
- code (int) — 返回码
- msg (string) — 返回信息
- data (object) — 数据对象
  - username (string) — 用户名

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "username": "string"
  },
  "msg": "string"
}
```

## 接口: 删除用户
- 方法: DELETE
- 路径: /api/admin/users/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回删除的用户ID，包含通用包装

#### 返回参数
- data (object)
  - id (integer) — 用户ID
- code (int) — 返回码
- msg (string) — 返回信息

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 接口: 用户详情
- 方法: GET
- 路径: /api/admin/users/{id}
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回单个用户详情

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)
  - username (string)
  - email (string)
  - created_at (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "created_at": "string",
    "email": "string",
    "id": 123,
    "username": "string"
  },
  "msg": "string"
}
```

## 接口: 用户列表
- 方法: GET
- 路径: /api/admin/users
- 认证: jwt
- 角色: admin

### 请求参数
- [query] page (int) 必填:false 说明:
- [query] page_size (int) 必填:false 说明:
- [query] search (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回用户列表（分页与搜索）

#### 返回参数
- msg (string)
- data (object)
  - items (array)
    - item (object)
      - id (integer)
      - username (string)
      - email (string)
      - role (string)
      - status (string)
      - created_at (string)
  - total (int)
  - page (int)
  - page_size (int)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "items": [
      {
        "created_at": "string",
        "email": "string",
        "id": 123,
        "role": "string",
        "status": "string",
        "username": "string"
      }
    ],
    "page": 123,
    "page_size": 123,
    "total": 123
  },
  "msg": "string"
}
```

## 接口: 当前管理员信息
- 方法: GET
- 路径: /api/admin/users/me
- 说明: 返回当前登录管理员的基础信息
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: <nil>

#### 返回参数
- msg (string)
- data (object)
  - username (string)
  - roles (array)
  - perms (array)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "perms": [
      "string"
    ],
    "roles": [
      "string"
    ],
    "username": "string"
  },
  "msg": "string"
}
```

## 接口: 绑定用户角色
- 方法: POST
- 路径: /api/admin/user_roles/bind
- 认证: jwt
- 角色: admin

### 请求参数
- [body] user_id (int) 必填:true 说明:
- [body] role_id (int) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回绑定的用户与角色ID

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - user_id (int)
  - role_id (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "role_id": 123,
    "user_id": 123
  },
  "msg": "string"
}
```

## 接口: 用户角色列表
- 方法: GET
- 路径: /api/admin/user_roles
- 认证: jwt
- 角色: admin

### 请求参数
- [query] user_id (int) 必填:false 说明:
- [query] role_id (int) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回用户角色列表

#### 返回参数
- msg (string)
- data (array)
  - item (object)
    - user_id (int)
    - role_id (int)
    - role_name (string)
- code (int)

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "role_id": 123,
      "role_name": "string",
      "user_id": 123
    }
  ],
  "msg": "string"
}
```

## 接口: 设置用户角色
- 方法: POST
- 路径: /api/admin/users/{id}/roles
- 认证: jwt
- 角色: admin

### 请求参数
- [path] id (int) 必填:true 说明:
- [body] roles (array) 必填:true 说明:角色ID数组

### 响应
- 状态码: 200
- 返回说明: 返回设置后的角色数量

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - user_id (int)
  - roles_count (int)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "roles_count": 123,
    "user_id": 123
  },
  "msg": "string"
}
```

## 接口: 解绑用户角色
- 方法: POST
- 路径: /api/admin/user_roles/unbind
- 认证: jwt
- 角色: admin

### 请求参数
- [body] user_id (int) 必填:true 说明:
- [body] role_id (int) 必填:true 说明:

### 响应
- 状态码: 200
- 返回说明: 返回解绑的用户与角色ID

#### 返回参数
- data (object)
  - user_id (int)
  - role_id (int)
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "role_id": 123,
    "user_id": 123
  },
  "msg": "string"
}
```

## 接口: 更新用户
- 方法: PUT
- 路径: /api/admin/users/{id}
- 认证: jwt
- 角色: admin

### 请求参数
- [body] password (string) 必填:false 说明:
- [body] username (string) 必填:false 说明:
- [body] email (string) 必填:false 说明:

### 响应
- 状态码: 200
- 返回说明: 返回更新的用户ID，包含通用包装

#### 返回参数
- msg (string) — 返回信息
- data (object)
  - id (integer) — 用户ID
- code (int) — 返回码

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

