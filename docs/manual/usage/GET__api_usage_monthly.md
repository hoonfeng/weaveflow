# 模块：usage

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

