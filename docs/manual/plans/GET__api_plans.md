# 模块：plans

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 套餐列表
- 接口方法：GET
- 接口路径：/api/plans
- 接口说明：列出所有可用套餐
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/plans" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/plans" -H "Authorization: Bearer invalid"
```

### 返回参数
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - monthly_quota (int)：月配额
    - price (number)：单价
- code (int)

- 返回说明：返回套餐列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": 123,
      "monthly_quota": 123,
      "name": "string",
      "price": 1.23
    }
  ],
  "msg": "string"
}
```

