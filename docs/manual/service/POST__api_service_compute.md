# 模块：service

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 付费计算示例
- 接口方法：POST
- 接口路径：/api/service/compute
- 接口说明：使用ApiKey+HMAC签名访问，对调用进行计量
- 认证方式：apikey

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | x | int | true | 输入X |  |
| body | y | int | true | 输入Y |  |
| header | X-Api-Key | string | true | 密钥 |  |
| header | X-Timestamp | string | true | RFC3339 时间戳 |  |
| header | X-Nonce | string | true | 随机字符串 |  |
| header | X-Signature | string | true | Base64URL HMAC-SHA256 |  |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/service/compute" -H "X-Api-Key: string" -H "X-Timestamp: string" -H "X-Nonce: string" -H "X-Signature: string" -H "X-API-Key: <key>" -H "X-Signature: <hmac>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "x": 1,
  "y": 1
}
JSON
```
```json
{
  "x": 1,
  "y": 1
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/service/compute" -H "X-API-Key: <key>" -H "X-Signature: invalid"
```

### 认证说明（HMAC-SHA256）
必需请求头：`X-Api-Key`、`X-Timestamp(RFC3339)`、`X-Nonce`、`X-Signature`。
计算步骤：
1) 计算 BodyHash：`base64url(sha256(body))`（空体则为空字符串）
2) 构造消息：
```
POST
/api/service/compute
<timestamp>
<nonce>
<bodyHash>
```

3) 计算签名：`base64url( HMAC-SHA256(secret, message) )`

#### 代码示例（Go）
```go
package main
import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/base64"
)
func sign(secret, method, path, ts, nonce, bodyHash string) string {
  msg := method + "\n" + path + "\n" + ts + "\n" + nonce + "\n" + bodyHash
  mac := hmac.New(sha256.New, []byte(secret))
  mac.Write([]byte(msg))
  return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
```

#### 代码示例（JavaScript/Node.js）
```js
const crypto = require('crypto')
function sign(secret, method, path, ts, nonce, bodyHash){
  const msg = method+'\n'+path+'\n'+ts+'\n'+nonce+'\n'+bodyHash
  const h = crypto.createHmac('sha256', secret).update(msg).digest()
  return h.toString('base64url')
}
```

#### 代码示例（Python）
```python
import hmac, hashlib, base64
def sign(secret: str, method: str, path: str, ts: str, nonce: str, body_hash: str) -> str:
    msg = "\n".join([method, path, ts, nonce, body_hash]).encode()
    dig = hmac.new(secret.encode(), msg, hashlib.sha256).digest()
    return base64.urlsafe_b64encode(dig).decode().rstrip('=')
```

### 返回参数
- code (integer)：返回码
- msg (string)：返回信息
- data (object)：结果对象
  - z (integer)：x+y 的和

### 错误码
| Code | Message | Description |
|---|---|---|
| E_SIGNATURE_INVALID | 签名无效 | HMAC 校验失败 |
| E_QUOTA_EXCEEDED | 超出配额 | 请升级套餐或稍后再试 |

#### 错误响应示例
```json
{
  "code": "E_SIGNATURE_INVALID",
  "msg": "签名无效"
}
```

- 返回说明：返回计算结果z，包含通用包装
- 响应头：
  - Content-Type：application/json

### 示例数据
```json
{
  "code": 123,
  "data": {
    "z": 123
  },
  "msg": "string"
}
```

