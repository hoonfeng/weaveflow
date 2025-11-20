# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## Hooks 检视
- 接口方法：POST
- 接口路径：/api/admin/hooks/inspect
- 接口说明：查看当前接口的策略链
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | method | string | false |  |  |
| body | path | string | false |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/hooks/inspect" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "method": "string",
  "path": "string"
}
JSON
```
```json
{
  "method": "string",
  "path": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/hooks/inspect" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - chain (array)：策略链顺序
  - match (object)：匹配信息

- 返回说明：返回策略链与匹配结果

### 示例数据
```json
{
  "code": 123,
  "data": {
    "chain": [
      "string"
    ],
    "match": {}
  },
  "msg": "string"
}
```

