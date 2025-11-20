# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 更新子账户角色
- 接口方法：PUT
- 接口路径：/api/portal/users/{uid}

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | role | string | true |  |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/portal/users/{uid}" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "role": "string"
}
JSON
```
```json
{
  "role": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/portal/users/{uid}"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回更新结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

