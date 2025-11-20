# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

