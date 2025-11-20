# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付统计
- 接口方法：GET
- 接口路径：/api/admin/overview/payments_stats
- 接口说明：返回支付状态与渠道分布
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview/payments_stats" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/overview/payments_stats" -H "Authorization: Bearer invalid"
```

### 返回参数
- by_status (array)
- by_provider (array)

- 返回说明：<nil>

### 示例数据
```json
{
  "by_provider": [
    "string"
  ],
  "by_status": [
    "string"
  ]
}
```

