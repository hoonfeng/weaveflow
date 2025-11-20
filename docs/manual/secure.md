# 模块：secure

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 受保护接口示例
- 接口方法：GET
- 接口路径：/api/secure/ping
- 接口说明：需要携带JWT且包含admin角色
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/secure/ping" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/secure/ping" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - msg (string)：消息
- code (int)
- msg (string)

### 错误码
| Code | Message | Description |
|---|---|---|
| E_UNAUTHORIZED | 未授权 | 缺少或无效的JWT |
| E_FORBIDDEN | 无权限 | 需要admin角色 |

#### 错误响应示例
```json
{
  "code": "E_UNAUTHORIZED",
  "msg": "未授权"
}
```

- 返回说明：返回通用包装，data为消息体

### 示例数据
```json
{
  "code": 123,
  "data": {
    "msg": "string"
  },
  "msg": "string"
}
```

