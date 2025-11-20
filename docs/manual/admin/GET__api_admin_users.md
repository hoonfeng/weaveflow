# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 用户列表
- 接口方法：GET
- 接口路径：/api/admin/users
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | page | int | false |  |  |
| query | page_size | int | false |  |  |
| query | search | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users?page=1&page_size=1&search=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/users" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - items (array)
    - item (object)
      - id (integer)
      - username (string)
      - email (string)
      - role (string)
      - status (string)
      - created_at (string)
  - total (int)
  - page (int)
  - page_size (int)

- 返回说明：返回用户列表（分页与搜索）

### 示例数据
```json
{
  "code": 123,
  "data": {
    "items": [
      {
        "created_at": "string",
        "email": "string",
        "id": 123,
        "role": "string",
        "status": "string",
        "username": "string"
      }
    ],
    "page": 123,
    "page_size": 123,
    "total": 123
  },
  "msg": "string"
}
```

