# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建套餐
- 接口方法：POST
- 接口路径：/api/admin/plans
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | name | string | true |  |  |
| body | monthly_quota | int | false |  |  |
| body | price | number | true |  |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plans" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "monthly_quota": 1,
  "name": "string",
  "price": 1.23
}
JSON
```
```json
{
  "monthly_quota": 1,
  "name": "string",
  "price": 1.23
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/plans" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (object)
  - name (string)
  - monthly_quota (int)
  - price (number)

- 返回说明：返回创建的套餐名与价格

### 示例数据
```json
{
  "code": 123,
  "data": {
    "monthly_quota": 123,
    "name": "string",
    "price": 1.23
  },
  "msg": "string"
}
```

