# 模块：vector

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 向量检索示例
- 接口方法：POST
- 接口路径：/api/vector/demo
- 接口说明：演示向量 upsert 与 search

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | id | string | true | 元素ID |  |
| body | vec | array | true | 输入向量 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/demo" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "id": "string",
  "vec": [
    "string"
  ]
}
JSON
```
```json
{
  "id": "string",
  "vec": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/demo"
```

### 返回参数
- data (array)
  - item (object)
    - id (string)
    - score (number)
    - payload (object)
- code (int)
- msg (string)

- 返回说明：返回检索结果列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": "string",
      "payload": {},
      "score": 1.23
    }
  ],
  "msg": "string"
}
```

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

## Qdrant 检索示例
- 接口方法：POST
- 接口路径：/api/vector/qdrant/search
- 接口说明：基于 Qdrant 进行向量检索（需配置 datasources.vector.qdrant）

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | vec | array | true | 查询向量 |  |
| body | topK | int | false | 返回数量 |  |
| body | options | object | false | 额外检索选项(filter/params) |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/qdrant/search" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "options": {
    "key": "value"
  },
  "topK": 1,
  "vec": [
    "string"
  ]
}
JSON
```
```json
{
  "options": {
    "key": "value"
  },
  "topK": 1,
  "vec": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/vector/qdrant/search"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (string)：向量ID
    - score (number)：相似度分数
    - payload (object)：元数据

### 错误码
| Code | Message | Description |
|---|---|---|
| E_VECTOR_NOT_AVAILABLE | 向量服务不可用 | 请检查 qdrant 配置 |
| E_COLLECTION_NOT_FOUND | 集合不存在 | 请先初始化集合 |

#### 错误响应示例
```json
{
  "code": "E_VECTOR_NOT_AVAILABLE",
  "msg": "向量服务不可用"
}
```

- 返回说明：返回检索结果列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": "string",
      "payload": {},
      "score": 1.23
    }
  ],
  "msg": "string"
}
```

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

