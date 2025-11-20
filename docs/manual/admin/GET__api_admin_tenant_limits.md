# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

