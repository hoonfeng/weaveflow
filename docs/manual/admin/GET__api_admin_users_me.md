# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 当前管理员信息
- 接口方法：GET
- 接口路径：/api/admin/users/me
- 接口说明：返回当前登录管理员的基础信息
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users/me" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users/me" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - username (string)
  - roles (array)
  - perms (array)

- 返回说明：<nil>

### 示例数据
```json
{
  "code": 123,
  "data": {
    "perms": [
      "string"
    ],
    "roles": [
      "string"
    ],
    "username": "string"
  },
  "msg": "string"
}
```

