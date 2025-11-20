# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 设置角色权限
- 接口方法：POST
- 接口路径：/api/admin/roles/{id}/permissions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | perm_ids | array | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles/1/permissions" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "perm_ids": [
    "string"
  ]
}
JSON
```
```json
{
  "perm_ids": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles/{id}/permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- role_id (int)
- count (int)

- 返回说明：返回设置后的权限数量

### 示例数据
```json
{
  "count": 123,
  "role_id": 123
}
```

