# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建子账户并绑定到租户
- 接口方法：POST
- 接口路径：/api/portal/users

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true |  |  |
| body | email | string | true |  |  |
| body | password | string | true |  |  |
| body | role | string | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/users" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "email": "string",
  "password": "string",
  "role": "string",
  "username": "string"
}
JSON
```
```json
{
  "email": "string",
  "password": "string",
  "role": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/users"
```

### 返回参数
- msg (string)
- data (object)
  - uid (int)
- code (int)

- 返回说明：返回创建绑定后的用户ID

### 示例数据
```json
{
  "code": 123,
  "data": {
    "uid": 123
  },
  "msg": "string"
}
```

