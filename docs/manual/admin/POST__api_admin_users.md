# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建用户
- 接口方法：POST
- 接口路径：/api/admin/users
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | email | string | false |  |  |
| body | password | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/users" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)：数据对象
  - username (string)：用户名
- code (int)：返回码
- msg (string)：返回信息

- 返回说明：返回创建的用户名，包含通用包装

### 示例数据
```json
{
  "code": 123,
  "data": {
    "username": "string"
  },
  "msg": "string"
}
```

