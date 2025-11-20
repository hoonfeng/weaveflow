# 模块: string

## 接口: 文本反转
- 方法: POST
- 路径: /api/text/reverse
- 说明: 通过插件实现字符串反转
- 认证: apikey

### 请求参数
- [body] text (string) 必填:true 说明:待反转文本

### 响应
- 状态码: 200
- 返回说明: 返回反转后的文本，包含通用包装
- 响应头:
  - Content-Type: application/json

#### 返回参数
- code (integer) — 返回码
- msg (string) — 返回信息
- data (string) — 反转结果

#### 示例数据
```json
{
  "code": 123,
  "data": "string",
  "msg": "string"
}
```

