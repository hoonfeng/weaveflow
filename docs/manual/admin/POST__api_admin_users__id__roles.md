# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 设置用户角色
- 接口方法：POST
- 接口路径：/api/admin/users/{id}/roles
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | roles | array | true | 角色ID数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users/1/roles" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "roles": [
    "string"
  ]
}
JSON
```
```json
{
  "roles": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users/{id}/roles" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - roles_count (int)
  - user_id (int)
- code (int)
- msg (string)

- 返回说明：返回设置后的角色数量

### 示例数据
```json
{
  "code": 123,
  "data": {
    "roles_count": 123,
    "user_id": 123
  },
  "msg": "string"
}
```

