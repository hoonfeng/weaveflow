# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建角色
- 接口方法：POST
- 接口路径：/api/admin/roles
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true | 角色名称 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "name": "string"
}
JSON
```
```json
{
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/roles" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (object)
  - name (string)：角色名
- code (int)

- 返回说明：返回创建的角色名

### 示例数据
```json
{
  "code": 123,
  "data": {
    "name": "string"
  },
  "msg": "string"
}
```

