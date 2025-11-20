# 模块：apikeys

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 撤销API Key
- 接口方法：POST
- 接口路径：/api/apikeys/revoke
- 接口说明：将密钥状态更新为撤销
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | id | int | true | 密钥ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/apikeys/revoke" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "id": 1
}
JSON
```
```json
{
  "id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/apikeys/revoke" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - id (integer)：密钥ID
- code (int)
- msg (string)

- 返回说明：返回撤销的密钥ID

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

