# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建权限
- 接口方法：POST
- 接口路径：/api/admin/permissions
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | code | string | true | 权限代码 |  |
| body | name | string | true | 权限名称 |  |
| body | description | string | false | 描述 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/permissions" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "code": "string",
  "description": "string",
  "name": "string"
}
JSON
```
```json
{
  "code": "string",
  "description": "string",
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/permissions" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (string)

- 返回说明：返回创建的权限代码

### 示例数据
```json
{
  "code": "string"
}
```

