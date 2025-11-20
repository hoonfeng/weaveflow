# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 依赖健康聚合
- 接口方法：GET
- 接口路径：/admin/health
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/admin/health" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/admin/health" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)：返回码
- msg (string)：返回信息
- data (object)：健康数据
  - plugin_disabled_count (integer)：被禁用插件数量
  - sql_main (integer)：主库探测 1=ok
  - plugin_count (integer)：插件数量

- 返回说明：返回数据库与插件运行时的健康统计

### 示例数据
```json
{
  "code": 123,
  "data": {
    "plugin_count": 123,
    "plugin_disabled_count": 123,
    "sql_main": 123
  },
  "msg": "string"
}
```

