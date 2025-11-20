# 模块: file

## 接口: 文件上传
- 方法: POST
- 路径: /api/files/upload
- 说明: 接收并保存上传的文件到本地存储

### 请求参数
- [file] files (<nil>) 必填:true 说明:上传的文件列表

### 响应
- 状态码: 201
- 返回说明: 返回保存后的文件标识列表
- 响应头:
  - X-Uploaded: true

#### 返回参数
- code (int) — 状态码
- msg (string) — 描述
- data (array) — 文件ID列表
  - item (string) — 文件ID

#### 示例数据
```json
{
  "code": 123,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

