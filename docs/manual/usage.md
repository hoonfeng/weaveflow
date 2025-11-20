# 模块：usage

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 日用量汇总
- 接口方法：GET
- 接口路径：/api/usage/daily
- 接口说明：按天统计租户用量
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/daily?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/daily" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)：返回码
- msg (string)：返回信息
- data (array)：每日统计
  - item (object)
    - day (string)：日期
    - calls (integer)：调用次数
    - avg_ms (number)：平均耗时毫秒

- 返回说明：返回每日用量，包含调用次数与平均时长

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "avg_ms": 1.23,
      "calls": 123,
      "day": "string"
    }
  ],
  "msg": "string"
}
```

## 用量导出CSV
- 接口方法：GET
- 接口路径：/api/usage/export.csv
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | rangeStart | string | false |  |  |
| query | rangeEnd | string | false |  |  |
| query | groupBy | string | false |  |  |
| query | period | string | false |  |  |
| query | tenant_id | int | false |  |  |
| query | endpoint | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/export.csv?rangeStart=string&rangeEnd=string&groupBy=string&period=string&tenant_id=1&endpoint=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/export.csv" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）
- 响应头：
  - Content-Type：text/csv

### 示例数据
```
{{ csv }}
```

## 月用量汇总
- 接口方法：GET
- 接口路径：/api/usage/monthly
- 接口说明：按月统计租户用量
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/monthly?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/monthly" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)：返回信息
- data (array)：每月统计
  - item (object)
    - avg_ms (number)：平均耗时毫秒
    - month (string)：月份
    - calls (integer)：调用次数
- code (int)：返回码

- 返回说明：返回每月用量，包含调用次数与平均时长

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "avg_ms": 1.23,
      "calls": 123,
      "month": "string"
    }
  ],
  "msg": "string"
}
```

## 用量报表
- 接口方法：POST
- 接口路径：/api/usage/reports
- 接口说明：按日/月聚合调用次数与平均耗时
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | rangeEnd | string | false | 结束日期YYYY-MM-DD |  |
| query | groupBy | string | false | 分组维度(tenant|endpoint|provider) |  |
| query | period | string | false | 聚合周期(daily|monthly) |  |
| query | tenant_id | int | false | 租户ID过滤 |  |
| query | endpoint | string | false | 接口过滤 |  |
| query | rangeStart | string | false | 开始日期YYYY-MM-DD |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/usage/reports?rangeEnd=string&groupBy=string&period=string&tenant_id=1&endpoint=string&rangeStart=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/usage/reports" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)：返回码
- msg (string)：返回信息
- data (array)：报表数据
  - item (object)
    - bucket (string)：时间桶（日/月）
    - calls (integer)：调用次数
    - avg_ms (number)：平均耗时毫秒
    - min_ms (number)：最小耗时毫秒
    - max_ms (number)：最大耗时毫秒
    - tenant_id (integer)：租户ID，可选
    - endpoint (string)：接口名，可选
    - provider (string)：提供方，可选

- 返回说明：返回聚合报表，支持多维度分组

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "avg_ms": 1.23,
      "bucket": "string",
      "calls": 123,
      "endpoint": "string",
      "max_ms": 1.23,
      "min_ms": 1.23,
      "provider": "string",
      "tenant_id": 123
    }
  ],
  "msg": "string"
}
```

