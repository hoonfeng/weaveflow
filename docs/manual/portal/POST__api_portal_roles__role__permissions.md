# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 设置角色权限
- 接口方法：POST
- 接口路径：/api/portal/roles/{role}/permissions

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | codes | array | true | 权限码数组 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/roles/{role}/permissions" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "codes": [
    "string"
  ]
}
JSON
```
```json
{
  "codes": [
    "string"
  ]
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/portal/roles/{role}/permissions"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - synced (int)

- 返回说明：返回设置后的权限数量

### 示例数据
```json
{
  "code": 123,
  "data": {
    "synced": 123
  },
  "msg": "string"
}
```

