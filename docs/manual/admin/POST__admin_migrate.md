# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 模型迁移
- 接口方法：POST
- 接口路径：/admin/migrate
- 接口说明：按模型配置对齐数据库结构
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/admin/migrate" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/admin/migrate" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)：表名列表
  - item (string)：表名

### 错误码
| Code | Message | Description |
|---|---|---|
| E_MIGRATE_FAILED | 迁移失败 | 数据库不可达或模型错误 |

#### 错误响应示例
```json
{
  "code": "E_MIGRATE_FAILED",
  "msg": "迁移失败"
}
```

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

