# WeaveFlow 织流

一个以配置与 DSL 为核心的 Go API 后端框架：用 YAML 编织步骤、数据源、钩子与插件，快速构建可运行的接口。

## 特性
- 接口配置化：`configs/interfaces*/*.yaml` 描述端点、认证、参数与步骤流水线
- 步骤 DSL：内置 `sql.query/sql.exec/transform/branch/response/model.apply` 等能力
- 模板函数：丰富的字符串、哈希、集合、时间函数用于数据加工
- 认证与权限：JWT/ApiKey/HMAC 等，支持角色与权限 DSL
- 插件运行时：`plugins/` 进程式插件（goproc），热重启与并发实例管理
- 数据源适配：SQL/KV/Blob/Vector 等，统一封装
- 模型迁移：`configs/models/` + 启动自动迁移或管理端接口迁移
- 数据初始化：`configs/init/*.yaml` 幂等种子脚本
- 静态资源：可托管管理端 UI 或单页应用
- 热重载：监控配置变更并动态重建路由

## 快速开始
- 环境
  - Go 1.21+
  - MySQL 8.0+（或按需替换数据源）
- 获取
  - `git clone <your-repo-url>`
- 配置
  - 在系统环境或 `.env` 中设置 `DB_DSN`（示例：`user:pass@tcp(127.0.0.1:3306)/ifaceconf_dev?parseTime=true`）
  - `configs/project.yaml` 读取 `datasources.sql.main.dsn: ${DB_DSN:}`（为空则跳过数据库初始化）
- 启动
  - `go run ./cmd/server/main.go`
- 验证
  - 最小接口目录：`configs/interfaces_minimal`
  - 登录示例：`POST /api/auth/login`
  - 文档：`GET /api/system/openapi` 或访问 `docs/`

## 仓库结构
- `cmd/server/`：入口程序
- `internal/`：核心实现（配置、路由、引擎、数据源、模型、插件、指标等）
- `configs/`：项目配置、最小接口、模型与种子脚本
- `plugins/`：插件示例源码
- `webui/admin/`：管理端 UI 源码（构建产物已忽略）
- `third_party/goproc/`：插件运行时依赖（含 LICENSE）
- `docs/`：学习路径与参考手册（生成文档建议在 CI 中产出）
- `test/`：集成测试

## 学习路径
- 从 `docs/学习路径/第一阶段` 开始，依次进阶到第二/第三/第四阶段
- 重点阅读：
  - `第二阶段/01-步骤DSL深入学习.md`
  - `第二阶段/04-模型配置与自动迁移.md`
  - `第二阶段/05-数据初始化与种子脚本.md`
  - `参考手册/02-内置能力清单.md` 与 `参考手册/01-模板函数参考.md`

## 配置要点
- `configs/project.yaml`
  - `datasources.sql.main.dsn` 使用 `${DB_DSN:}`（避免提交敏感默认值）
  - `init.scripts` 声明种子脚本（幂等）
  - `static.items` 托管静态资源（可选）
- 模型迁移
  - 启动自动迁移：入口在 `cmd/server/main.go`
  - 管理接口：`POST /api/admin/migrate`
- 步骤 DSL 示例
  - `configs/init/minimal.yaml` 展示“查询→变换→分支→执行”的幂等脚本模式

## 命名与定位
- 框架名称：WeaveFlow（中文：织流）
- 含义：以 DSL 将“步骤、数据源、钩子与插件”编织为可运行的接口流，突出编排与组合能力

## 许可证
- 建议使用 Apache-2.0 License（兼顾商用与贡献者权益）

## CI 建议
- GitHub Actions：
  - Go 构建与测试：`go build ./cmd/server`、`go test ./...`
  - 前端（可选）：`webui/admin` 执行 `npm ci && npm run build`
  - OpenAPI 文档：在 CI 生成并上传构件而非提交大文件

## 贡献
- Issue 与 PR 欢迎提交：修复、文档改进、插件示例、数据源适配等

## 致谢
- 依赖与第三方：`chi`、`mysql-driver`、`yaml`、`goproc` 等