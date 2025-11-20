# DSL 步骤能力总览（与代码实现一致）

## 已实现步骤
- `response`：构建状态码/头/体，支持 `wrap{code,msg,data}`（internal/core/engine.go:39–56）
- `transform`：根据 `mapping` 渲染并写入 `ctx.vars`（internal/core/engine.go:56–62）
- `validate`：单目标校验（类型/必填/范围/正则/枚举/文件大小与类型）（internal/core/engine.go:63–114）
- `cache.set` / `cache.get`：进程内缓存读写（internal/core/engine.go:115–130）
- `sql.query` / `sql.exec`：SQL 查询与执行，支持 `order` 指定参数顺序与 `ctx.vars.sql_log` 记录（internal/core/engine.go:131–176）
- `auth.jwt`：签发 JWT（internal/core/engine.go:177–187）
- `kv.set` / `kv.get`：KV 存取（示例实现为内存 TTLCache）（internal/core/engine.go:188–205）
- `http.request`：外部 HTTP 调用，支持重试/退避与简易熔断 `circuit{key,threshold,openMs,fallback}`（internal/core/engine.go:206–270）
- `upload.save`：保存上传文件到 Blob，返回 `[]string` 文件ID（internal/core/engine.go:270–283；internal/datasource/manager.go:119–158）
- `admin.reload` / `admin.lint` / `admin.docs` / `admin.openapi` / `admin.builtin` / `admin.plugins_usage` / `admin.permissions_scan`（internal/core/engine.go:284–407）
- `hooks.inspect`：查看 Auth/Before/After 有效钩子（internal/core/engine.go:355–365）
- `model.apply`：应用模型（DDL）（internal/core/engine.go:408–417）
- `obs.metrics`：导出 Prometheus 文本（internal/core/engine.go:418–424）
- `plugin.call`：调用插件，支持重试与熔断（internal/core/engine.go:425–468）
- `branch` / `loop`：条件与循环控制（internal/core/engine.go:469–513）
- `vector.search/upsert/delete/ensure/upsert_batch`：向量能力（internal/core/engine.go:514–576）
- `transaction`：显式事务动作 `begin/commit/rollback`（internal/core/engine.go:577–590）

## 事务与限流
- 接口级事务：在接口配置 `transaction: { ds: <name> }`，路由自动插入 `begin/commit`（internal/router/router.go:472–483）
- 端点级限流：支持 `perIp/perTenant/perKey`（internal/router/router.go:82–110）

## 认证与安全
- JWT 与权限校验（roles/permissions）在路由层绑定（internal/router/router.go:65–77, 140–161）
- ApiKey + HMAC：签名字段 `X-Api-Key/X-Timestamp/X-Nonce/X-Signature`，消息为 `METHOD\nPATH\nTS\nNONCE\nBODY_HASH`（internal/router/auth_hmac.go:21–33, 47–53）
- 防重放：`Timestamp` 容差与 `Nonce` TTLCache（internal/router/auth_hmac.go:39–45, 55–59）

## 钩子
- 文件拆分：`configs/hooks/auth.yaml/before.yaml/after.yaml`
- 匹配条件：`module/endpoint/method/path/pathPrefix/pathRegex/interfaces/label/labelsAny/labelsAll`（modules/apphooks/config.go:118–151）
- 参数注入：`params.headers` 与 `bodyHash: sha256`（modules/apphooks/config.go:189–211）

## 说明
- 文档示例应严格遵循上述步骤与参数；未列出的步骤键不属于当前实现