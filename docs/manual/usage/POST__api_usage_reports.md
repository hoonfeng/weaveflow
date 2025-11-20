# 模块：usage

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

