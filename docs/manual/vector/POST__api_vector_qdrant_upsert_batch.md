# 模块：vector

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## Qdrant 批量插入
- 接口方法：POST
- 接口路径：/api/vector/qdrant/upsert_batch
- 接口说明：批量写入点（需配置 datasources.vector.qdrant）

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | collection | string | true | 集合名 |  |
| body | items | array | true | 列表 [{id,vec,meta}] |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/qdrant/upsert_batch" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "collection": "string",
  "items": [
    "string"
  ]
}
JSON
```
```json
{
  "collection": "string",
  "items": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/qdrant/upsert_batch"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)：状态

- 返回说明：返回批量写入状态

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

