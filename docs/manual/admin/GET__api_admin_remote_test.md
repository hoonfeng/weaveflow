# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 远程调用测试（熔断回退）
- 接口方法：GET
- 接口路径：/api/admin/remote/test
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/remote/test" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/remote/test" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)：远程响应或fallback

- 返回说明：返回远程调用结果或回退内容

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

