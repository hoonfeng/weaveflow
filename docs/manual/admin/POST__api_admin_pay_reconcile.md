# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 支付对账
- 接口方法：POST
- 接口路径：/api/admin/pay/reconcile
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | start | string | true | 开始日期YYYY-MM-DD |  |
| body | end | string | true | 结束日期YYYY-MM-DD |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/reconcile" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "end": "string",
  "start": "string"
}
JSON
```
```json
{
  "end": "string",
  "start": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/admin/pay/reconcile" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
{{ rows }}
```

