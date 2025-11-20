# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 启用/禁用 API Key
- 接口方法：POST
- 接口路径：/api/portal/apikeys/{id}/toggle

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | status | string | true |  | enum=[active disabled] |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys/{id}/toggle" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "status": "string"
}
JSON
```
```json
{
  "status": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys/{id}/toggle"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)

- 返回说明：返回更新后的状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string"
  },
  "msg": "string"
}
```

