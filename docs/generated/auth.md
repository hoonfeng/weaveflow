# 模块: auth

## 接口: 用户登录
- 方法: POST
- 路径: /api/auth/login
- 说明: 校验用户名与密码并签发JWT

### 请求参数
- [body] username (string) 必填:true 说明:用户名
- [body] password (string) 必填:true 说明:密码

### 响应
- 状态码: 0
- 返回说明: 返回签发的令牌与通用包装结构
- 响应头:
  - Content-Type: application/json

#### 返回参数
- code (integer) — 返回码
- msg (string) — 返回信息
- data (object) — 数据对象
  - token (string) — 令牌字符串

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "token": "string"
  },
  "msg": "string"
}
```

