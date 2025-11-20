# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 模型迁移
- 接口方法：POST
- 接口路径：/api/admin/migrate
- 接口说明：按模型配置对齐数据库结构
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/migrate" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/migrate" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (array)：表名列表
  - item (string)：表名
- code (int)

- 返回说明：返回迁移的模型表名列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

