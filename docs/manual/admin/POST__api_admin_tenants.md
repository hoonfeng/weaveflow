# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建租户
- 接口方法：POST
- 接口路径：/api/admin/tenants
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | domain | string | false | 域名 |  |
| body | contact_email | string | true | 联系邮箱 |  |
| body | description | string | false | 描述 |  |
| body | status | string | false | 状态 | default=active |
| body | name | string | true | 租户名称 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/tenants" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
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
curl -X POST "http://localhost:8080/api/admin/tenants" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - name (string)：租户名
  - status (string)：状态

- 返回说明：返回创建的租户名和状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string",
    "status": "string"
  },
  "msg": "string"
}
```

