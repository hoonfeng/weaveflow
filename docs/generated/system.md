# 模块: system

## 接口: 接口文档列表
- 方法: GET
- 路径: /docs

### 响应
- 状态码: 200
- 返回说明: 返回所有接口文档端点列表

#### 示例数据
```
{{ endpoints }}
```

## 接口: 指标导出
- 方法: GET
- 路径: /metrics

### 响应
- 状态码: 200
- 返回说明: 返回 Prometheus 文本格式指标
- 响应头:
  - Content-Type: text/plain; version=0.0.4

#### 示例数据
```
{{ m }}
```

## 接口: OpenAPI 文档导出
- 方法: GET
- 路径: /openapi.json

### 响应
- 状态码: 200
- 返回说明: 返回项目的 OpenAPI 结构（JSON）
- 响应头:
  - Content-Type: application/json

#### 返回参数
- openapi (string) — 版本
- info (object) — 项目信息
- paths (object) — 路径定义

#### 示例数据
```json
{
  "info": {},
  "openapi": "string",
  "paths": {}
}
```

