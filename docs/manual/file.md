# 模块：file

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 文件上传
- 接口方法：POST
- 接口路径：/api/files/upload
- 接口说明：接收并保存上传的文件到本地存储

### 请求参数
| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |
|---|---|---|---|---|---|
| file | files | <nil> | true | 上传的文件列表 | maxSize=32MB, types=[image/png image/jpeg] |

### 请求示例
```bash
curl -X POST "http://localhost:8080/api/files/upload"
```

### 失败请求示例
```bash
curl -X POST "http://localhost:8080/api/files/upload"
```

### 返回参数
- code (int)：状态码
- msg (string)：描述
- data (array)：文件ID列表
  - item (string)：文件ID

### 错误码
| Code | Message | Description |
|---|---|---|
| E_FILE_TOO_LARGE | 文件过大 | 超过最大 32MB |
| E_FILE_TYPE_DENIED | 文件类型不允许 | 仅允许 image/png |

#### 错误响应示例
```json
{
  "code": "E_FILE_TOO_LARGE",
  "msg": "文件过大"
}
```

- 返回说明：返回保存后的文件标识列表
- 响应头：
  - X-Uploaded：true

### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

