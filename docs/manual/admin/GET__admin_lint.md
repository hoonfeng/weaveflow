# 模块：admin

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 接口规范检查
- 接口方法：GET
- 接口路径：/admin/lint
- 接口说明：返回当前已加载接口的规范问题列表
- 认证方式：jwt
- 访问角色：admin

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/admin/lint" -H "Authorization: Bearer <token>"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/admin/lint" -H "Authorization: Bearer invalid"
```

### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - path (string)：位置
    - level (string)：级别(error|warn|info)
    - message (string)：描述

### 错误码
| Code | Message | Description |
|---|---|---|
| E_LINT_FAILED | 检查失败 | 配置文件无法解析 |

#### 错误响应示例
```json
{
  "code": "E_LINT_FAILED",
  "msg": "检查失败"
}
```

- 返回说明：返回接口配置的检查结果列表

### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "level": "string",
      "message": "string",
      "path": "string"
    }
  ],
  "msg": "string"
}
```

