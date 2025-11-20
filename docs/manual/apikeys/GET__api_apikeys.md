# 模块：apikeys

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 列出API Keys
- 接口方法：GET
- 接口路径：/api/apikeys
- 接口说明：按租户列出所有API Key
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/apikeys?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/apikeys" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (array)
  - item (object)
    - id (integer)：Key ID
    - name (string)：名称
    - status (string)：状态
    - expired_at (string)：过期时间
    - created_at (string)：创建时间
- code (int)
- msg (string)

- 返回说明：返回 API Key 列表

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
      "status": "string"
    }
  ],
  "msg": "string"
}
```

