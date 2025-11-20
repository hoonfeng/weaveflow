# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 更新支付通道配置
- 接口方法：PUT
- 接口路径：/api/admin/pay/providers/{id}
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true |  |  |
| body | name | string | false |  |  |
| body | enabled | bool | false |  |  |
| body | config_json | string | false |  |  |

### 请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/pay/providers/1" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "config_json": "string",
  "enabled": true,
  "name": "string"
}
JSON
```
```json
{
  "config_json": "string",
  "enabled": true,
  "name": "string"
}
```

### 失败请求示例
```bash
curl -X PUT "http://localhost:8080/api/admin/pay/providers/{id}" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
{{ row.0 }}
```

