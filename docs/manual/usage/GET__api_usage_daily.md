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

