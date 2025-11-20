# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 租户自助注册
- 接口方法：POST
- 接口路径：/api/portal/tenants/register

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | email | string | true |  |  |
| body | password | string | true |  |  |
| body | plan_id | int | false |  |  |
| body | tenant_name | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/tenants/register" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "plan_id": 1,
  "tenant_name": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "plan_id": 1,
  "tenant_name": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/tenants/register"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - token (string)

- 返回说明：返回注册成功后的令牌

### 示例数据
```json
{
  "code": 123,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

