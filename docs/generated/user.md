# 模块: user

## 接口: 获取用户信息
- 方法: GET
- 路径: /api/user/{id}
- 说明: 获取指定用户的基本资料

### 请求参数
- [path] id (int) 必填:true 说明:用户ID
- [header] Authorization (string) 必填:true 说明:认证令牌

### 响应
- 状态码: 200
- 返回说明: 返回用户的基本信息
- 响应头:
  - X-Trace: 请求追踪标识

#### 返回参数
- data (object) — 用户数据
  - id (int) — 用户ID
  - name (string) — 用户名
- code (int) — 状态码
- msg (string) — 描述

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123,
    "name": "string"
  },
  "msg": "string"
}
```

