# 模块: secure

## 接口: 受保护接口示例
- 方法: GET
- 路径: /api/secure/ping
- 说明: 需要携带JWT且包含admin角色
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回通用包装，data为消息体

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - msg (string) — 消息

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "msg": "string"
  },
  "msg": "string"
}
```

