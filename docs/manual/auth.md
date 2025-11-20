# 模块：auth

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 用户登录
- 接口方法：POST
- 接口路径：/api/auth/login
- 接口说明：校验用户名与密码并签发JWT

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | username | string | true | 用户名 | minLen=1 |
| body | password | string | true | 密码 | minLen=6 |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/auth/login" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "password": "string",
  "username": "string"
}
JSON
```
```json
{
  "password": "string",
  "username": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/auth/login"
```

### 返回参数
- code (integer)：返回码
- msg (string)：返回信息
- data (object)：数据对象
  - token (string)：令牌字符串

### 错误码
| Code | Message | Description |
|---|---|---|
| E_AUTH_INVALID | 用户名或密码错误 | 认证失败 |
| E_AUTH_LOCKED | 账户已锁定 | 请联系管理员 |

#### 错误响应示例
```json
{
  "code": "E_AUTH_INVALID",
  "msg": "用户名或密码错误"
}
```

- 返回说明：返回签发的令牌与通用包装结构
- 响应头：
  - Content-Type：application/json

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

