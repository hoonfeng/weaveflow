# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

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

