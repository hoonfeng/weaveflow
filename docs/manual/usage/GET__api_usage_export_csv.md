# 模块：usage

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 用量导出CSV
- 接口方法：GET
- 接口路径：/api/usage/export.csv
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | rangeStart | string | false |  |  |
| query | rangeEnd | string | false |  |  |
| query | groupBy | string | false |  |  |
| query | period | string | false |  |  |
| query | tenant_id | int | false |  |  |
| query | endpoint | string | false |  |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/export.csv?rangeStart=string&rangeEnd=string&groupBy=string&period=string&tenant_id=1&endpoint=string" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/usage/export.csv" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）
- 响应头：
  - Content-Type：text/csv

### 示例数据
```
{{ csv }}
```

