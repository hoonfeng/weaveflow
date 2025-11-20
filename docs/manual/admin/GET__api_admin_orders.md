# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 订单列表
- 接口方法：GET
- 接口路径：/api/admin/orders
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |
| query | page | int | false |  |  |
| query | size | int | false |  |  |
| query | page_size | int | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders?tenant_id=1&provider=string&status=string&page=1&size=1&page_size=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - order_no (string)：订单号
    - amount (number)：金额
    - status (string)：状态

- 返回说明：返回订单列表，包含金额与状态

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "amount": 1.23,
      "order_no": "string",
      "status": "string"
    }
  ],
  "msg": "string"
}
```

