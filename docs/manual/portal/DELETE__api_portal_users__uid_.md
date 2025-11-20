# 模块：portal

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 删除子账户绑定
- 接口方法：DELETE
- 接口路径：/api/portal/users/{uid}

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X DELETE "http://localhost:8080/api/portal/users/{uid}"
```

### 失败请求示例
```bash
curl -X DELETE "http://localhost:8080/api/portal/users/{uid}"
```

### 返回参数
- code (int)
- msg (string)
- data (object)

- 返回说明：返回删除结果

### 示例数据
```json
{
  "code": 123,
  "data": {},
  "msg": "string"
}
```

