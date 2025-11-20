# 模块：string

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 文本反转
- 接口方法：POST
- 接口路径：/api/text/reverse
- 接口说明：通过插件实现字符串反转
- 认证方式：apikey

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| body | text | string | true | 待反转文本 | minLen=1 |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/text/reverse" -H "X-API-Key: <key>" -H "X-Signature: <hmac>" -H "Content-Type: application/json" -d @- << 'JSON'
{
  "text": "string"
}
JSON
```
```json
{
  "text": "string"
}
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/text/reverse" -H "X-API-Key: <key>" -H "X-Signature: invalid"
```

### 认证说明（HMAC-SHA256）
必需请求头：`X-Api-Key`、`X-Timestamp(RFC3339)`、`X-Nonce`、`X-Signature`。
计算步骤：
1) 计算 BodyHash：`base64url(sha256(body))`（空体则为空字符串）
2) 构造消息：
```
POST
/api/text/reverse
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
- data (string)：反转结果

- 返回说明：返回反转后的文本，包含通用包装
- 响应头：
  - Content-Type：application/json

### 示例数据
```json
{
  "code": 123,
  "data": "string",
  "msg": "string"
}
```

