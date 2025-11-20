# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## API Key 列表
- 接口方法：GET
- 接口路径：/api/portal/apikeys

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/apikeys"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/portal/apikeys"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - status (string)
    - expired_at (string)
    - scopes (string)
    - created_at (string)

- 返回说明：返回当前租户的 API Key 列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "created_at": "string",
      "expired_at": "string",
      "id": 123,
      "name": "string",
      "scopes": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

