# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建 API Key
- 接口方法：POST
- 接口路径：/api/portal/apikeys

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true |  |  |
| body | scopes | string | false |  |  |
| body | expired_at | string | false |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "expired_at": "string",
  "name": "string",
  "scopes": "string"
}
JSON
```
```json
{
  "expired_at": "string",
  "name": "string",
  "scopes": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/apikeys"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - id (int)
  - key (string)
  - secret (string)

- 返回说明：返回创建的明文 key 与 secret

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123,
    "key": "string",
    "secret": "string"
  },
  "msg": "string"
}
```

