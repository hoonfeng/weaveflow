# 模块：system

说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。

## 接口文档列表
- 接口方法：GET
- 接口路径：/docs

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/docs"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/docs"
```

### 返回参数
（未声明结构，参考示例数据）

- 返回说明：返回所有接口文档端点列表

### 示例数据
```
{{ endpoints }}
```

## 指标导出
- 接口方法：GET
- 接口路径：/metrics

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/metrics"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/metrics"
```

### 返回参数
（未声明结构，参考示例数据）

- 返回说明：返回 Prometheus 文本格式指标
- 响应头：
  - Content-Type：text/plain; version=0.0.4

### 示例数据
```
{{ m }}
```

## OpenAPI 文档导出
- 接口方法：GET
- 接口路径：/openapi.json

### 请求参数
（无显式参数，或仅 Body/Headers 由业务生成）

### 请求示例
```bash
curl -X GET "http://localhost:8080/openapi.json"
```

### 失败请求示例
```bash
curl -X GET "http://localhost:8080/openapi.json"
```

### 返回参数
- openapi (string)：版本
- info (object)：项目信息
- paths (object)：路径定义

- 返回说明：返回项目的 OpenAPI 结构（JSON）
- 响应头：
  - Content-Type：application/json

### 示例数据
```json
{
  "info": {},
  "openapi": "string",
  "paths": {}
}
```

