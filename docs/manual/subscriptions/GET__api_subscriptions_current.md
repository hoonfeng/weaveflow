# 模块：subscriptions

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 当前订阅
- 接口方法：GET
- 接口路径：/api/subscriptions/current
- 接口说明：查询租户当前有效订阅与配额
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/subscriptions/current?tenant_id=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/subscriptions/current" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
{{ rows }}
```

