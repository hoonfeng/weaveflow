# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 订单导出CSV
- 接口方法：GET
- 接口路径：/api/admin/orders/export.csv
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | false |  |  |
| query | provider | string | false |  |  |
| query | status | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/export.csv?tenant_id=1&provider=string&status=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/admin/orders/export.csv" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

- 返回说明：返回 CSV 文本内容
- 响应头：
  - Content-Type：text/csv

### 示例数据
```
order_no,tenant_id,amount,currency,provider,status,ts
ord_001,1,99.9,CNY,alipay,succeeded,2025-11-01T10:00:00Z
```

