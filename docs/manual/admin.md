# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## API Key 轮换
- 接口方法：POST
- 接口路径：/api/admin/apikeys/rotate
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true |  |  |
| body | name | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/apikeys/rotate" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "name": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "name": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/apikeys/rotate" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (object)
  - key (string)
  - secret (string)
- code (int)

- 返回说明：返回新 Key 与 Secret

### 示例数据
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

## API Key 启用/禁用
- 接口方法：POST
- 接口路径：/api/admin/apikeys/toggle
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true |  |  |
| body | name | string | true |  |  |
| body | enabled | bool | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/apikeys/toggle" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "enabled": true,
  "name": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "enabled": true,
  "name": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/apikeys/toggle" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - tenant_id (int)
  - name (string)
  - status (string)：状态(active|disabled)

- 返回说明：返回更新后的状态

### 示例数据
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

## 依赖健康聚合
- 接口方法：GET
- 接口路径：/api/admin/health
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/health" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/health" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)：返回信息
- data (object)：健康数据
  - sql_main (integer)：主库探测 1=ok
  - plugin_count (integer)：插件数量
  - plugin_disabled_count (integer)：被禁用插件数量
- code (int)：返回码

- 返回说明：返回数据库与插件运行时的健康统计

### 示例数据
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

## Hooks 检视
- 接口方法：POST
- 接口路径：/api/admin/hooks/inspect
- 接口说明：查看当前接口的策略链
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | method | string | false |  |  |
| body | path | string | false |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/hooks/inspect" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "method": "string",
  "path": "string"
}
JSON
```
```json
{
  "method": "string",
  "path": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/hooks/inspect" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - chain (array)：策略链顺序
  - match (object)：匹配信息

- 返回说明：返回策略链与匹配结果

### 示例数据
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

## 订单导出CSV
- 接口方法：GET
- 接口路径：/api/admin/orders/export.csv
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
curl -X GET "http://localhost:8080/api/admin/orders/export.csv?tenant_id=1&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/export.csv" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

- 返回说明：返回 CSV 文本内容
- 响应头：
  - Content-Type：text/csv

### 示例数据
```
order_no,tenant_id,amount,currency,provider,status,ts
ord_001,1,99.9,CNY,alipay,succeeded,2025-11-01T10:00:00Z
```

## 订单列表
- 接口方法：GET
- 接口路径：/api/admin/orders
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |
| query | page | int | false |  |  |
| query | size | int | false |  |  |
| query | page_size | int | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders?tenant_id=1&provider=string&status=string&page=1&size=1&page_size=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - order_no (string)：订单号
    - amount (number)：金额
    - status (string)：状态

- 返回说明：返回订单列表，包含金额与状态

### 示例数据
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

## 概览统计
- 接口方法：GET
- 接口路径：/api/admin/overview
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview" -H "Authorization: Bearer invalid"
```

### 返回参数
- total_users (int)
- total_orders (int)

- 返回说明：返回基础统计（示例字段）

### 示例数据
```json
{
  "total_orders": 123,
  "total_users": 123
}
```

## 支付统计
- 接口方法：GET
- 接口路径：/api/admin/overview/payments_stats
- 接口说明：返回支付状态与渠道分布
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview/payments_stats" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview/payments_stats" -H "Authorization: Bearer invalid"
```

### 返回参数
- by_status (array)
- by_provider (array)

- 返回说明：<nil>

### 示例数据
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

## 概览趋势
- 接口方法：GET
- 接口路径：/api/admin/overview/trends
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview/trends" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview/trends" -H "Authorization: Bearer invalid"
```

### 返回参数
- users (array)
- orders (array)

- 返回说明：返回近12月的用户与订单趋势

### 示例数据
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

## 创建支付通道
- 接口方法：POST
- 接口路径：/api/admin/pay/providers
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | code | string | true | 渠道代码如 wechat/alipay |  |
| body | name | string | true | 渠道名称 |  |
| body | enabled | bool | false | 是否启用 |  |
| body | config_json | string | false | JSON 配置 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/providers" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "code": "string",
  "config_json": "string",
  "enabled": true,
  "name": "string"
}
JSON
```
```json
{
  "code": "string",
  "config_json": "string",
  "enabled": true,
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/providers" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - enabled (bool)
  - config_json (string)
  - id (int)
  - code (string)
  - name (string)

- 返回说明：返回创建后的记录

### 示例数据
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

## 删除支付通道
- 接口方法：DELETE
- 接口路径：/api/admin/pay/providers/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/pay/providers/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/pay/providers/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
map[]
```

## 支付通道详情
- 接口方法：GET
- 接口路径：/api/admin/pay/providers/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/pay/providers/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/pay/providers/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
{{ row.0 }}
```

## 支付通道列表
- 接口方法：GET
- 接口路径：/api/admin/pay/providers
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/pay/providers" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/pay/providers" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - id (int)
  - code (string)
  - name (string)
  - enabled (bool)
  - config_json (string)

- 返回说明：返回支付通道配置列表

### 示例数据
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

## 启用/禁用支付通道
- 接口方法：POST
- 接口路径：/api/admin/pay/providers/{id}/toggle
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | enabled | bool | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/providers/1/toggle" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "enabled": true
}
JSON
```
```json
{
  "enabled": true
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/providers/{id}/toggle" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
map[enabled:{{ enabled }} id:{{ id }}]
```

## 更新支付通道配置
- 接口方法：PUT
- 接口路径：/api/admin/pay/providers/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | name | string | false |  |  |
| body | enabled | bool | false |  |  |
| body | config_json | string | false |  |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/pay/providers/1" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "config_json": "string",
  "enabled": true,
  "name": "string"
}
JSON
```
```json
{
  "config_json": "string",
  "enabled": true,
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/pay/providers/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
{{ row.0 }}
```

## 支付对账
- 接口方法：POST
- 接口路径：/api/admin/pay/reconcile
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | start | string | true | 开始日期YYYY-MM-DD |  |
| body | end | string | true | 结束日期YYYY-MM-DD |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/reconcile" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "end": "string",
  "start": "string"
}
JSON
```
```json
{
  "end": "string",
  "start": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/reconcile" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
{{ rows }}
```

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

## 支付记录导出
- 接口方法：GET
- 接口路径：/api/admin/payments/export
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | order_no | string | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments/export?order_no=string&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments/export" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - txn_id (string)
  - order_no (string)
  - provider (string)
  - status (string)
  - ts (string)

- 返回说明：返回支付记录列表（JSON），用于下载
- 响应头：
  - Content-Type：application/json

### 示例数据
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

## 支付记录导出CSV
- 接口方法：GET
- 接口路径：/api/admin/payments/export.csv
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | order_no | string | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments/export.csv?order_no=string&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments/export.csv" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

- 返回说明：返回 CSV 文本内容
- 响应头：
  - Content-Type：text/csv

### 示例数据
```
order_no,provider,status,txn_id,ts
ord_001,alipay,succeeded,tx_001,2025-11-01T10:00:00Z
```

## 支付记录列表
- 接口方法：GET
- 接口路径：/api/admin/payments
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | page | int | false |  |  |
| query | size | int | false |  |  |
| query | page_size | int | false |  |  |
| query | order_no | string | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments?page=1&size=1&page_size=1&order_no=string&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - txn_id (string)：交易ID
    - order_no (string)：订单号
    - provider (string)：支付渠道
    - status (string)：状态
    - ts (string)：时间

- 返回说明：返回支付记录列表，包含交易ID与状态

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

## 创建权限
- 接口方法：POST
- 接口路径：/api/admin/permissions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | code | string | true | 权限代码 |  |
| body | name | string | true | 权限名称 |  |
| body | description | string | false | 描述 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/permissions" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "code": "string",
  "description": "string",
  "name": "string"
}
JSON
```
```json
{
  "code": "string",
  "description": "string",
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (string)

- 返回说明：返回创建的权限代码

### 示例数据
```json
{
  "code": "string"
}
```

## 删除权限
- 接口方法：DELETE
- 接口路径：/api/admin/permissions/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/permissions/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/permissions/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- id (int)

- 返回说明：返回删除的权限ID

### 示例数据
```json
{
  "id": 123
}
```

## 权限列表
- 接口方法：GET
- 接口路径：/api/admin/permissions
- 接口说明：返回权限代码、名称与描述
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/permissions" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - id (int)
  - code (string)
  - name (string)
  - description (string)

- 返回说明：返回权限列表

### 示例数据
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

## 角色权限列表
- 接口方法：GET
- 接口路径：/api/admin/role_permissions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/role_permissions" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/role_permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - perm_id (int)
  - code (string)
  - name (string)
  - role_id (int)

- 返回说明：返回角色拥有的权限

### 示例数据
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

## 设置角色权限
- 接口方法：POST
- 接口路径：/api/admin/roles/{id}/permissions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | perm_ids | array | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles/1/permissions" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "perm_ids": [
    "string"
  ]
}
JSON
```
```json
{
  "perm_ids": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles/{id}/permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- role_id (int)
- count (int)

- 返回说明：返回设置后的权限数量

### 示例数据
```json
{
  "count": 123,
  "role_id": 123
}
```

## 同步权限声明入库
- 接口方法：POST
- 接口路径：/api/admin/permissions/sync
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/permissions/sync" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/permissions/sync" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - scanned (int)
  - synced (int)
- code (int)
- msg (string)

- 返回说明：返回同步的数量

### 示例数据
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

## 更新权限
- 接口方法：PUT
- 接口路径：/api/admin/permissions/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | code | string | false | 权限代码 |  |
| body | name | string | false | 权限名称 |  |
| body | description | string | false | 描述 |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/permissions/{id}" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "code": "string",
  "description": "string",
  "name": "string"
}
JSON
```
```json
{
  "code": "string",
  "description": "string",
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/permissions/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- id (int)

- 返回说明：返回更新的权限ID

### 示例数据
```json
{
  "id": 123
}
```

## 创建套餐
- 接口方法：POST
- 接口路径：/api/admin/plans
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true |  |  |
| body | monthly_quota | int | false |  |  |
| body | price | number | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plans" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "monthly_quota": 1,
  "name": "string",
  "price": 1.23
}
JSON
```
```json
{
  "monthly_quota": 1,
  "name": "string",
  "price": 1.23
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plans" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - name (string)
  - monthly_quota (int)
  - price (number)

- 返回说明：返回创建的套餐名与价格

### 示例数据
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

## 删除套餐
- 接口方法：DELETE
- 接口路径：/api/admin/plans/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/plans/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/plans/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)

- 返回说明：返回删除的套餐ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 套餐列表
- 接口方法：GET
- 接口路径：/api/admin/plans
- 接口说明：返回套餐列表，包含配额与价格
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plans" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plans" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - monthly_quota (int)
    - price (string)

- 返回说明：返回套餐列表

### 示例数据
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

## 更新套餐
- 接口方法：PUT
- 接口路径：/api/admin/plans/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | false |  |  |
| body | monthly_quota | int | false |  |  |
| body | price | string | false |  |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/plans/{id}" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "monthly_quota": 1,
  "name": "string",
  "price": "string"
}
JSON
```
```json
{
  "monthly_quota": 1,
  "name": "string",
  "price": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/plans/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)

- 返回说明：返回更新的套餐ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 动态添加插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/add
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true | 插件名称 |  |
| body | windows | string | true | Windows可执行文件路径 |  |
| body | unix | string | false | Unix可执行文件路径 |  |
| body | instances | int | false | 实例数量 |  |
| body | timeout | string | false | 超时时间，如300ms |  |
| body | queueSize | int | false | 队列大小 |  |
| body | functions | array | true | 函数列表 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/add" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "functions": [
    "string"
  ],
  "instances": 1,
  "name": "string",
  "queueSize": 1,
  "timeout": "string",
  "unix": "string",
  "windows": "string"
}
JSON
```
```json
{
  "functions": [
    "string"
  ],
  "instances": 1,
  "name": "string",
  "queueSize": 1,
  "timeout": "string",
  "unix": "string",
  "windows": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/add" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回添加结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 内置插件列表
- 接口方法：GET
- 接口路径：/api/admin/plugins/builtin
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/builtin" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/builtin" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - name (string)
  - desc (string)

- 返回说明：返回内置插件能力列表

### 示例数据
```json
{
  "item": {
    "desc": "string",
    "name": "string"
  }
}
```

## 禁用插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/disable
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | names | array | true | 插件名称数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/disable" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "names": [
    "string"
  ]
}
JSON
```
```json
{
  "names": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/disable" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - ok (boolean)：是否成功

- 返回说明：返回禁用结果

### 示例数据
```json
{
  "code": 123,
  "data": {
    "ok": true
  },
  "msg": "string"
}
```

## 启用插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/enable
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | names | array | true | 插件名称数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/enable" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "names": [
    "string"
  ]
}
JSON
```
```json
{
  "names": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/enable" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (object)
  - ok (boolean)：是否成功
- code (int)

- 返回说明：返回启用结果

### 示例数据
```json
{
  "code": 123,
  "data": {
    "ok": true
  },
  "msg": "string"
}
```

## 外置插件列表
- 接口方法：GET
- 接口路径：/api/admin/plugins/external
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/external" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/external" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - name (string)
  - endpoints (array)

- 返回说明：返回对外提供API服务的插件列表

### 示例数据
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

## 内置插件列表
- 接口方法：GET
- 接口路径：/api/admin/plugins/internal
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/internal" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/internal" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - name (string)
  - endpoints (array)

- 返回说明：返回系统必要功能的组成插件列表

### 示例数据
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

## 动态移除插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/remove
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | names | array | true | 插件名称数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/remove" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "names": [
    "string"
  ]
}
JSON
```
```json
{
  "names": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/remove" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回移除结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 重启插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/restart
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | names | array | true | 插件名称数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/restart" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "names": [
    "string"
  ]
}
JSON
```
```json
{
  "names": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/restart" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回重启结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 启动插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/start
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | names | array | true | 插件名称数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/start" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "names": [
    "string"
  ]
}
JSON
```
```json
{
  "names": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/start" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回启动结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 插件状态
- 接口方法：GET
- 接口路径：/api/admin/plugins/status
- 接口说明：返回当前已注册插件的状态
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/status" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/plugins/status" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - running (boolean)
  - functions (array)
  - name (string)
  - enabled (boolean)

- 返回说明：<nil>

### 示例数据
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

## 停止插件
- 接口方法：POST
- 接口路径：/api/admin/plugins/stop
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | names | array | true | 插件名称数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/stop" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "names": [
    "string"
  ]
}
JSON
```
```json
{
  "names": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plugins/stop" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回停止结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 创建角色
- 接口方法：POST
- 接口路径：/api/admin/roles
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true | 角色名称 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "name": "string"
}
JSON
```
```json
{
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (object)
  - name (string)：角色名
- code (int)

- 返回说明：返回创建的角色名

### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string"
  },
  "msg": "string"
}
```

## 删除角色
- 接口方法：DELETE
- 接口路径：/api/admin/roles/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/roles/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/roles/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：角色ID

- 返回说明：返回删除的角色ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 角色列表
- 接口方法：GET
- 接口路径：/api/admin/roles
- 接口说明：返回全部角色名称与ID
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/roles" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/roles" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (integer)：角色ID
    - name (string)：角色名

- 返回说明：返回角色列表

### 示例数据
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

## 更新角色
- 接口方法：PUT
- 接口路径：/api/admin/roles/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true | 角色名称 |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/roles/{id}" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "name": "string"
}
JSON
```
```json
{
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/roles/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：角色ID

- 返回说明：返回更新的角色ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 取消订阅
- 接口方法：POST
- 接口路径：/api/admin/subscriptions/cancel
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions/cancel" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "tenant_id": 1
}
JSON
```
```json
{
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions/cancel" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - tenant_id (int)：租户ID
  - status (string)：订阅状态

- 返回说明：返回取消订阅的租户ID与状态

### 示例数据
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

## 变更订阅
- 接口方法：POST
- 接口路径：/api/admin/subscriptions/change
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true |  |  |
| body | plan_id | int | true |  |  |
| body | status | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions/change" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "plan_id": 1,
  "status": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "plan_id": 1,
  "status": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions/change" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)：订阅状态
  - tenant_id (int)：租户ID

- 返回说明：返回变更的租户ID与订阅状态

### 示例数据
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

## 创建订阅
- 接口方法：POST
- 接口路径：/api/admin/subscriptions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | plan_id | int | true |  |  |
| body | tenant_id | int | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "plan_id": 1,
  "tenant_id": 1
}
JSON
```
```json
{
  "plan_id": 1,
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/subscriptions" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - plan_id (int)
  - tenant_id (int)
- code (int)
- msg (string)

- 返回说明：返回租户订阅创建结果

### 示例数据
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

## 当前订阅查询
- 接口方法：GET
- 接口路径：/api/admin/subscriptions/current
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions/current?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions/current" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)
  - tenant_id (int)
  - plan_id (int)
  - status (string)
  - start_at (string)
  - end_at (string)

- 返回说明：返回指定租户的当前订阅

### 示例数据
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

## 订阅列表
- 接口方法：GET
- 接口路径：/api/admin/subscriptions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/subscriptions" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - status (string)：状态
    - id (int)：订阅ID
    - tenant_id (int)：租户ID
    - plan_id (int)：计划ID
    - plan_name (string)：计划名称

- 返回说明：返回订阅列表

### 示例数据
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

## 接口规范检查
- 接口方法：GET
- 接口路径：/api/admin/lint
- 接口说明：返回当前已加载接口的规范问题列表
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/lint" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/lint" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - level (string)：级别(error|warn|info)
    - message (string)：描述
    - path (string)：位置

- 返回说明：返回接口配置的检查结果列表

### 示例数据
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

## 模型迁移
- 接口方法：POST
- 接口路径：/api/admin/migrate
- 接口说明：按模型配置对齐数据库结构
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/migrate" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/migrate" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (array)：表名列表
  - item (string)：表名
- code (int)

- 返回说明：返回迁移的模型表名列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

## 清理HMAC非重放随机值
- 接口方法：POST
- 接口路径：/api/admin/nonces/purge
- 接口说明：删除早于TTL的nonce记录
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | ttlSeconds | int | false | TTL秒，默认900 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/nonces/purge?ttlSeconds=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/nonces/purge" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
map[ttlSeconds:{{ ttlSeconds }}]
```

## 队列状态
- 接口方法：GET
- 接口路径：/api/admin/queue/status
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/queue/status" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/queue/status" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - running (int)：运行中任务数
  - workers (int)：工作线程数量
  - pending (int)：等待任务数
- code (int)
- msg (string)

- 返回说明：返回队列运行状态

### 示例数据
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

## 重载配置
- 接口方法：POST
- 接口路径：/api/admin/reload
- 接口说明：重新加载项目与接口配置并返回差异
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/reload" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/reload" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - added (array)：新增端点
  - removed (array)：移除端点
  - updated (array)：更新端点

- 返回说明：返回重载差异

### 示例数据
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

## 远程调用示例（本地调试）
- 接口方法：GET
- 接口路径：/api/admin/remote/debug
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/remote/debug" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/remote/debug" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)：插件状态原样透传

- 返回说明：返回插件运行时状态

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 远程调用测试（熔断回退）
- 接口方法：GET
- 接口路径：/api/admin/remote/test
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/remote/test" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/remote/test" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)：远程响应或fallback

- 返回说明：返回远程调用结果或回退内容

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 租户API Key列表
- 接口方法：GET
- 接口路径：/api/admin/tenants/{id}/apikeys
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/apikeys" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/apikeys" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - expired_at (string)
    - scopes (string)
    - created_at (string)
    - id (int)
    - name (string)
    - status (string)

- 返回说明：返回租户的 API Key

### 示例数据
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

## 创建租户
- 接口方法：POST
- 接口路径：/api/admin/tenants
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | domain | string | false | 域名 |  |
| body | contact_email | string | true | 联系邮箱 |  |
| body | description | string | false | 描述 |  |
| body | status | string | false | 状态 | default=active |
| body | name | string | true | 租户名称 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenants" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "contact_email": "string",
  "description": "string",
  "domain": "string",
  "name": "string",
  "status": "string"
}
JSON
```
```json
{
  "contact_email": "string",
  "description": "string",
  "domain": "string",
  "name": "string",
  "status": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenants" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - name (string)：租户名
  - status (string)：状态

- 返回说明：返回创建的租户名和状态

### 示例数据
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

## 删除租户
- 接口方法：DELETE
- 接口路径：/api/admin/tenants/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：租户ID

- 返回说明：返回删除的租户ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 租户详情
- 接口方法：GET
- 接口路径：/api/admin/tenants/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)
  - description (string)
  - created_at (string)
  - id (int)
  - name (string)
  - domain (string)
  - contact_email (string)

- 返回说明：返回单个租户详情

### 示例数据
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

## 查询租户限流
- 接口方法：GET
- 接口路径：/api/admin/tenant_limits
- 接口说明：查询指定租户的限流设置
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenant_limits?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenant_limits" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - tenant_id (int)：租户ID
    - rps (number)：每秒请求限制
    - burst (int)：突发容量
    - updated_at (string)：更新时间

- 返回说明：返回租户限流设置列表（通常为单条）

### 示例数据
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

## 设置租户限流
- 接口方法：POST
- 接口路径：/api/admin/tenant_limits/upsert
- 接口说明：设定租户级rps与burst
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | rps | float | true | 每秒请求数 |  |
| body | burst | int | true | 突发 |  |
| body | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenant_limits/upsert" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "burst": 1,
  "rps": 1.23,
  "tenant_id": 1
}
JSON
```
```json
{
  "burst": 1,
  "rps": 1.23,
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenant_limits/upsert" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - rps (number)
  - burst (int)
  - tenant_id (int)

- 返回说明：返回更新后的限流设置

### 示例数据
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

## 租户列表
- 接口方法：GET
- 接口路径：/api/admin/tenants
- 接口说明：分页返回租户列表（示例为全部返回）
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (integer)：租户ID
    - name (string)：租户名
    - status (string)：状态
    - created_at (string)：创建时间

- 返回说明：返回租户列表

### 示例数据
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

## 租户订阅列表
- 接口方法：GET
- 接口路径：/api/admin/tenants/{id}/subscriptions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/subscriptions" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/subscriptions" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - next_billing_at (string)
    - id (int)
    - plan_id (int)
    - status (string)
    - start_at (string)
    - end_at (string)

- 返回说明：返回租户订阅记录

### 示例数据
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

## 更新租户
- 接口方法：PUT
- 接口路径：/api/admin/tenants/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | false | 租户名称 |  |
| body | domain | string | false | 域名 |  |
| body | contact_email | string | false | 联系邮箱 |  |
| body | description | string | false | 描述 |  |
| body | status | string | false | 状态 |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "contact_email": "string",
  "description": "string",
  "domain": "string",
  "name": "string",
  "status": "string"
}
JSON
```
```json
{
  "contact_email": "string",
  "description": "string",
  "domain": "string",
  "name": "string",
  "status": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：租户ID

- 返回说明：返回更新的租户ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 租户月度用量统计
- 接口方法：GET
- 接口路径：/api/admin/tenants/{id}/usage/monthly
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/usage/monthly" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/usage/monthly" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (array)
  - item (object)
    - ym (string)
    - cnt (int)
- code (int)

- 返回说明：返回最近12个月按月统计的调用次数

### 示例数据
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

## 绑定租户成员
- 接口方法：POST
- 接口路径：/api/admin/tenants/{id}/users
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | user_id | int | true | 用户ID |  |
| body | role | string | true | 成员角色 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenants/{id}/users" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "role": "string",
  "user_id": 1
}
JSON
```
```json
{
  "role": "string",
  "user_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenants/{id}/users" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回绑定结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 租户成员列表
- 接口方法：GET
- 接口路径：/api/admin/tenants/{id}/users
- 接口说明：返回指定租户的成员列表
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/users" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/tenants/{id}/users" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - uid (int)
    - username (string)
    - email (string)
    - role (string)
    - created_at (string)

- 返回说明：返回租户成员及其角色

### 示例数据
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

## 解绑租户成员
- 接口方法：DELETE
- 接口路径：/api/admin/tenants/{id}/users/{uid}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/tenants/{id}/users/{uid}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/tenants/{id}/users/{uid}" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (object)
- code (int)

- 返回说明：返回解绑结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

## 创建用户
- 接口方法：POST
- 接口路径：/api/admin/users
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | email | string | false |  |  |
| body | password | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)：数据对象
  - username (string)：用户名
- code (int)：返回码
- msg (string)：返回信息

- 返回说明：返回创建的用户名，包含通用包装

### 示例数据
```json
{
  "code": 123,
  "data": {
    "username": "string"
  },
  "msg": "string"
}
```

## 删除用户
- 接口方法：DELETE
- 接口路径：/api/admin/users/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/users/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/users/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)：返回码
- msg (string)：返回信息
- data (object)
  - id (integer)：用户ID

- 返回说明：返回删除的用户ID，包含通用包装

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

## 用户详情
- 接口方法：GET
- 接口路径：/api/admin/users/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (object)
  - id (int)
  - username (string)
  - email (string)
  - created_at (string)
- code (int)

- 返回说明：返回单个用户详情

### 示例数据
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

## 用户列表
- 接口方法：GET
- 接口路径：/api/admin/users
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | page | int | false |  |  |
| query | page_size | int | false |  |  |
| query | search | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users?page=1&page_size=1&search=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
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

- 返回说明：返回用户列表（分页与搜索）

### 示例数据
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

## 当前管理员信息
- 接口方法：GET
- 接口路径：/api/admin/users/me
- 接口说明：返回当前登录管理员的基础信息
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users/me" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users/me" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - username (string)
  - roles (array)
  - perms (array)

- 返回说明：<nil>

### 示例数据
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

## 绑定用户角色
- 接口方法：POST
- 接口路径：/api/admin/user_roles/bind
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | user_id | int | true |  |  |
| body | role_id | int | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/user_roles/bind" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "role_id": 1,
  "user_id": 1
}
JSON
```
```json
{
  "role_id": 1,
  "user_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/user_roles/bind" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - user_id (int)
  - role_id (int)

- 返回说明：返回绑定的用户与角色ID

### 示例数据
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

## 用户角色列表
- 接口方法：GET
- 接口路径：/api/admin/user_roles
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | user_id | int | false |  |  |
| query | role_id | int | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/user_roles?user_id=1&role_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/user_roles" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - user_id (int)
    - role_id (int)
    - role_name (string)

- 返回说明：返回用户角色列表

### 示例数据
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

## 设置用户角色
- 接口方法：POST
- 接口路径：/api/admin/users/{id}/roles
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | roles | array | true | 角色ID数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users/1/roles" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "roles": [
    "string"
  ]
}
JSON
```
```json
{
  "roles": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users/{id}/roles" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - roles_count (int)
  - user_id (int)
- code (int)
- msg (string)

- 返回说明：返回设置后的角色数量

### 示例数据
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

## 解绑用户角色
- 接口方法：POST
- 接口路径：/api/admin/user_roles/unbind
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | role_id | int | true |  |  |
| body | user_id | int | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/user_roles/unbind" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "role_id": 1,
  "user_id": 1
}
JSON
```
```json
{
  "role_id": 1,
  "user_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/user_roles/unbind" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - role_id (int)
  - user_id (int)

- 返回说明：返回解绑的用户与角色ID

### 示例数据
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

## 更新用户
- 接口方法：PUT
- 接口路径：/api/admin/users/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | false |  |  |
| body | email | string | false |  |  |
| body | password | string | false |  |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/users/{id}" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/users/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)：返回码
- msg (string)：返回信息
- data (object)
  - id (integer)：用户ID

- 返回说明：返回更新的用户ID，包含通用包装

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

