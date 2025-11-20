# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 绑定用户角色
- 接口方法：POST
- 接口路径：/api/admin/user_roles/bind
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | user_id | int | true |  |  |
| body | role_id | int | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/user_roles/bind" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "role_id": 1,
  "user_id": 1
}
JSON
```
```json
{
  "role_id": 1,
  "user_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/user_roles/bind" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - user_id (int)
  - role_id (int)

- 返回说明：返回绑定的用户与角色ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "role_id": 123,
    "user_id": 123
  },
  "msg": "string"
}
```

