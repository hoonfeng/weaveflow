# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 角色权限列表
- 接口方法：GET
- 接口路径：/api/admin/role_permissions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/role_permissions" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/role_permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- item (object)
  - perm_id (int)
  - code (string)
  - name (string)
  - role_id (int)

- 返回说明：返回角色拥有的权限

### 示例数据
```json
{
  "item": {
    "code": "string",
    "name": "string",
    "perm_id": 123,
    "role_id": 123
  }
}
```

