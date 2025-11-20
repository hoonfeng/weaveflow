# 文档与 OpenAPI（基于配置）

- 目的：不写代码，直接从 `configs/interfaces` 生成接口手册与 OpenAPI。

## 生成接口手册（逐接口）
- 在项目根目录执行：
```
 go run ./cmd/manualdoc --out docs/manual
```
- 产物：`docs/manual`（模块汇总与逐接口文档）。

## 生成模块文档与 OpenAPI
- 在项目根目录执行：
```
 go run ./cmd/docgen --out docs/generated
```
- 产物：`
  - `docs/generated/<module>.md` 模块汇总
  - `docs/generated/openapi.<module>.json` OpenAPI 3.0 规范

## 前端展示
- Portal 页面已对接 `/docs` 与 `/openapi.json`，具体参见页面配置。

## 文档字段建议
- 在接口的 `docs` 与 `response` 中补充 `title|description|headers|schema|example`，提升文档质量。