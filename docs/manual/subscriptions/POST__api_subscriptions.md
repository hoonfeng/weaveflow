# 模块：subscriptions

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 创建订阅
- 接口方法：POST
- 接口路径：/api/subscriptions
- 接口说明：为租户创建或切换订阅套餐
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | plan_id | int | true | 套餐ID |  |
| body | tenant_id | int | true | 租户ID |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/subscriptions" -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "plan_id": 1,
  "tenant_id": 1
}
JSON
```
```json
{
  "plan_id": 1,
  "tenant_id": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/subscriptions" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
map[plan_id:{{ plan_id }} tenant_id:{{ tenant_id }}]
```

