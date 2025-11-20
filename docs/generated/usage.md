# 模块: usage

## 接口: 日用量汇总
- 方法: GET
- 路径: /api/usage/daily
- 说明: 按天统计租户用量
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200
- 返回说明: 返回每日用量，包含调用次数与平均时长

#### 返回参数
- code (int) — 返回码
- msg (string) — 返回信息
- data (array) — 每日统计
  - item (object)
    - avg_ms (number) — 平均耗时毫秒
    - day (string) — 日期
    - calls (integer) — 调用次数

#### 示例数据
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

## 接口: 用量导出CSV
- 方法: GET
- 路径: /api/usage/export.csv
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:false 说明:
- [query] endpoint (string) 必填:false 说明:
- [query] rangeStart (string) 必填:false 说明:
- [query] rangeEnd (string) 必填:false 说明:
- [query] groupBy (string) 必填:false 说明:
- [query] period (string) 必填:false 说明:

### 响应
- 状态码: 200
- 响应头:
  - Content-Type: text/csv

#### 示例数据
```
{{ csv }}
```

## 接口: 月用量汇总
- 方法: GET
- 路径: /api/usage/monthly
- 说明: 按月统计租户用量
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200
- 返回说明: 返回每月用量，包含调用次数与平均时长

#### 返回参数
- code (int) — 返回码
- msg (string) — 返回信息
- data (array) — 每月统计
  - item (object)
    - month (string) — 月份
    - calls (integer) — 调用次数
    - avg_ms (number) — 平均耗时毫秒

#### 示例数据
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

## 接口: 用量报表
- 方法: POST
- 路径: /api/usage/reports
- 说明: 按日/月聚合调用次数与平均耗时
- 认证: jwt
- 角色: admin

### 请求参数
- [query] rangeStart (string) 必填:false 说明:开始日期YYYY-MM-DD
- [query] rangeEnd (string) 必填:false 说明:结束日期YYYY-MM-DD
- [query] groupBy (string) 必填:false 说明:分组维度(tenant|endpoint|provider)
- [query] period (string) 必填:false 说明:聚合周期(daily|monthly)
- [query] tenant_id (int) 必填:false 说明:租户ID过滤
- [query] endpoint (string) 必填:false 说明:接口过滤

### 响应
- 状态码: 200
- 返回说明: 返回聚合报表，支持多维度分组

#### 返回参数
- code (int) — 返回码
- msg (string) — 返回信息
- data (array) — 报表数据
  - item (object)
    - bucket (string) — 时间桶（日/月）
    - calls (integer) — 调用次数
    - avg_ms (number) — 平均耗时毫秒
    - min_ms (number) — 最小耗时毫秒
    - max_ms (number) — 最大耗时毫秒
    - tenant_id (integer) — 租户ID，可选
    - endpoint (string) — 接口名，可选
    - provider (string) — 提供方，可选

#### 示例数据
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

