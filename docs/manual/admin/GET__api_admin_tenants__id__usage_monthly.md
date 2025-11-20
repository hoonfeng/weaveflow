# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

