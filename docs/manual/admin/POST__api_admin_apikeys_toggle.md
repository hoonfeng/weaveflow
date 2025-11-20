# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## API Key 启用/禁用
- 接口方法：POST
- 接口路径：/api/admin/apikeys/toggle
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | tenant_id | int | true |  |  |
| body | name | string | true |  |  |
| body | enabled | bool | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/apikeys/toggle" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "enabled": true,
  "name": "string",
  "tenant_id": 1
}
JSON
```
```json
{
  "enabled": true,
  "name": "string",
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/apikeys/toggle" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - tenant_id (int)
  - name (string)
  - status (string)：状态(active|disabled)

- 返回说明：返回更新后的状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string",
    "status": "string",
    "tenant_id": 123
  },
  "msg": "string"
}
```

