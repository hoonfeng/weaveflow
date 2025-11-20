# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 队列状态
- 接口方法：GET
- 接口路径：/api/admin/queue/status
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/queue/status" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/queue/status" -H "Authorization: Bearer invalid"
```

### 返回参数
- data (object)
  - running (int)：运行中任务数
  - workers (int)：工作线程数量
  - pending (int)：等待任务数
- code (int)
- msg (string)

- 返回说明：返回队列运行状态

### 示例数据
```json
{
  "code": 123,
  "data": {
    "pending": 123,
    "running": 123,
    "workers": 123
  },
  "msg": "string"
}
```

