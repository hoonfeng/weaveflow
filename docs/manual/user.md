# 模块：user

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 获取用户信息
- 接口方法：GET
- 接口路径：/api/user/{id}
- 接口说明：获取指定用户的基本资料

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| path | id | int | true | 用户ID | min=1 |
| header | Authorization | string | true | 认证令牌 |  |

### 请求示例
```bash
curl -X GET "http://localhost:8080/api/user/1" -H "Authorization: string"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/api/user/{id}"
```

### 返回参数
- code (int)：状态码
- msg (string)：描述
- data (object)：用户数据
  - id (int)：用户ID
  - name (string)：用户名

### 错误码
| Code | Message | Description |
|---|---|---|
| E_NOT_FOUND | 用户不存在 | ID 无效 |

#### 错误响应示例
```json
{
  "code": "E_NOT_FOUND",
  "msg": "用户不存在"
}
```

- 返回说明：返回用户的基本信息
- 响应头：
  - X-Trace：请求追踪标识

### 示例数据
```json
{
  "code": 123,
  "data": {
    "id": 123,
    "name": "string"
  },
  "msg": "string"
}
```

