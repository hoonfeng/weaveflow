# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 清理HMAC非重放随机值
- 接口方法：POST
- 接口路径：/admin/nonces/purge
- 接口说明：删除早于TTL的nonce记录
- 认证方式：jwt
- 访问角色：admin

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| query | ttlSeconds | int | false | TTL秒，默认900 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/admin/nonces/purge?ttlSeconds=1" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/admin/nonces/purge" -H "Authorization: Bearer invalid"
```

### 返回参数
（未声明结构，参考示例数据）

### 示例数据
```
map[ttlSeconds:{{ ttlSeconds }}]
```

