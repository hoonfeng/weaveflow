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

