# 模块: service

## 接口: 付费计算示例
- 方法: POST
- 路径: /api/service/compute
- 说明: 使用ApiKey+HMAC签名访问，对调用进行计量
- 认证: apikey

### 请求参数
- [body] x (int) 必填:true 说明:输入X
- [body] y (int) 必填:true 说明:输入Y
- [header] X-Api-Key (string) 必填:true 说明:密钥
- [header] X-Timestamp (string) 必填:true 说明:RFC3339 时间戳
- [header] X-Nonce (string) 必填:true 说明:随机字符串
- [header] X-Signature (string) 必填:true 说明:Base64URL HMAC-SHA256

### 响应
- 状态码: 200
- 返回说明: 返回计算结果z，包含通用包装
- 响应头:
  - Content-Type: application/json

#### 返回参数
- msg (string) — 返回信息
- data (object) — 结果对象
  - z (integer) — x+y 的和
- code (integer) — 返回码

#### 示例数据
```json
{
  "code": 123,
  "data": {
    "z": 123
  },
  "msg": "string"
}
```

