# 钩子（Auth/Before/After）

- 目标：通过配置在鉴权前、执行前、执行后插入逻辑（通常调用插件）。
- 配置文件：`configs/hooks/*.yaml`（全局）；或在接口中用 `hooks` 字段覆盖/禁用。

## 全局钩子示例
```yaml
auth:
  - kind: plugin
    name: global_auth_check
    match: { module: admin }
    params: { plugin: auth, function: check, headers: { user: X-User-Id } }

before:
  - kind: plugin
    name: preflight
    match: { pathPrefix: "/api/" }
    params: { plugin: metrics, function: preflight }

after:
  - kind: plugin
    name: audit
    match: { method: [GET,POST], module: admin }
    params: { plugin: metrics, function: audit }
```

## 接口级覆盖与禁用
```yaml
hooks:
  disable: [global_auth_check]
  before:
    - { kind: plugin, name: local_pre, match: { path: "/api/service/compute" }, params: { plugin: metrics, function: preflight } }
```