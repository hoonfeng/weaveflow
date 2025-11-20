# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 更新租户
- 接口方法：PUT
- 接口路径：/api/admin/tenants/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | false | 租户名称 |  |
| body | domain | string | false | 域名 |  |
| body | contact_email | string | false | 联系邮箱 |  |
| body | description | string | false | 描述 |  |
| body | status | string | false | 状态 |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "contact_email": "string",
  "description": "string",
  "domain": "string",
  "name": "string",
  "status": "string"
}
JSON
```
```json
{
  "contact_email": "string",
  "description": "string",
  "domain": "string",
  "name": "string",
  "status": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/tenants/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (integer)：租户ID

- 返回说明：返回更新的租户ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123
  },
  "msg": "string"
}
```

