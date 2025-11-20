# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付记录列表
- 接口方法：GET
- 接口路径：/api/admin/payments
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | page | int | false |  |  |
| query | size | int | false |  |  |
| query | page_size | int | false |  |  |
| query | order_no | string | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments?page=1&size=1&page_size=1&order_no=string&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/payments" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - txn_id (string)：交易ID
    - order_no (string)：订单号
    - provider (string)：支付渠道
    - status (string)：状态
    - ts (string)：时间

- 返回说明：返回支付记录列表，包含交易ID与状态

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "order_no": "string",
      "provider": "string",
      "status": "string",
      "ts": "string",
      "txn_id": "string"
    }
  ],
  "msg": "string"
}
```

