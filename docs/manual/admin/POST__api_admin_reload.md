# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 重载配置
- 接口方法：POST
- 接口路径：/api/admin/reload
- 接口说明：重新加载项目与接口配置并返回差异
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/reload" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/reload" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - added (array)：新增端点
  - removed (array)：移除端点
  - updated (array)：更新端点

- 返回说明：返回重载差异

### 示例数据
```json
{
  "code": 123,
  "data": {
    "added": [
      "string"
    ],
    "removed": [
      "string"
    ],
    "updated": [
      "string"
    ]
  },
  "msg": "string"
}
```

