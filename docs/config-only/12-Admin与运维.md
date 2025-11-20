# Admin 与运维（零代码）

## 常用能力（步骤）
```yaml
- admin.docs:            { out: docs }          # 列出端点
- admin.openapi:         { out: spec }          # 生成 OpenAPI
- admin.permissions_scan:{ out: perms }         # 扫描权限声明
- admin.plugins_usage:   { out: usage }         # 分析插件使用情况
- admin.lint:            { out: lint }          # 接口配置检查
- admin.reload:          { out: reload }        # 热重载配置
```

## 运行指标
- 以 Prometheus 文本导出：
```yaml
- obs.metrics: { out: metrics }
```
- 路由标签自动记录模块，便于聚合分析。

## 插件状态（通过步骤）
- 在接口中使用 `plugins.status` 步骤将插件状态写入 `out`，例如：
```yaml
- plugins.status: { out: plugins }
- response: { status: 200, body: "{{ plugins }}" }
```