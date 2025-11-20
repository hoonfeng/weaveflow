# 模块：vector

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## Qdrant 集合初始化
- 接口方法：POST
- 接口路径：/api/vector/qdrant/init
- 接口说明：创建集合（需配置 datasources.vector.qdrant）

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | size | int | true | 向量维度 |  |
| body | metric | string | false | 距离度量(Cosine|Euclid|Dot) |  |
| body | collection | string | true | 集合名 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/qdrant/init" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "collection": "string",
  "metric": "string",
  "size": 1
}
JSON
```
```json
{
  "collection": "string",
  "metric": "string",
  "size": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/qdrant/init"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string)：状态

- 返回说明：返回集合创建状态

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

