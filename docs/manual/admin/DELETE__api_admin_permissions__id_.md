# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 删除权限
- 接口方法：DELETE
- 接口路径：/api/admin/permissions/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/permissions/{id}" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/admin/permissions/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- id (int)

- 返回说明：返回删除的权限ID

### 示例数据
```json
{
  "id": 123
}
```

