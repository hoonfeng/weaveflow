# 模块: vector

## 接口: 向量检索示例
- 方法: POST
- 路径: /api/vector/demo
- 说明: 演示向量 upsert 与 search

### 请求参数
- [body] id (string) 必填:true 说明:元素ID
- [body] vec (array) 必填:true 说明:输入向量

### 响应
- 状态码: 200
- 返回说明: 返回检索结果列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (string)
    - score (number)
    - payload (object)

#### 示例数据
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

## 接口: Qdrant 集合初始化
- 方法: POST
- 路径: /api/vector/qdrant/init
- 说明: 创建集合（需配置 datasources.vector.qdrant）

### 请求参数
- [body] size (int) 必填:true 说明:向量维度
- [body] metric (string) 必填:false 说明:距离度量(Cosine|Euclid|Dot)
- [body] collection (string) 必填:true 说明:集合名

### 响应
- 状态码: 200
- 返回说明: 返回集合创建状态

#### 返回参数
- data (object)
  - status (string) — 状态
- code (int)
- msg (string)

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string"
  },
  "msg": "string"
}
```

## 接口: Qdrant 检索示例
- 方法: POST
- 路径: /api/vector/qdrant/search
- 说明: 基于 Qdrant 进行向量检索（需配置 datasources.vector.qdrant）

### 请求参数
- [body] vec (array) 必填:true 说明:查询向量
- [body] topK (int) 必填:false 说明:返回数量
- [body] options (object) 必填:false 说明:额外检索选项(filter/params)

### 响应
- 状态码: 200
- 返回说明: 返回检索结果列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (string) — 向量ID
    - score (number) — 相似度分数
    - payload (object) — 元数据

#### 示例数据
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

## 接口: Qdrant 批量插入
- 方法: POST
- 路径: /api/vector/qdrant/upsert_batch
- 说明: 批量写入点（需配置 datasources.vector.qdrant）

### 请求参数
- [body] collection (string) 必填:true 说明:集合名
- [body] items (array) 必填:true 说明:列表 [{id,vec,meta}]

### 响应
- 状态码: 200
- 返回说明: 返回批量写入状态

#### 返回参数
- code (int)
- msg (string)
- data (object)
  - status (string) — 状态

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "status": "string"
  },
  "msg": "string"
}
```

