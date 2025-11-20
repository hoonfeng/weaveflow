# 常见问题（FAQ）

- Q：设置了限流但无效？
  - A：确认端点 `rateLimit.rps>0` 且与 `burst` 合理；若启用 `perIp|perTenant|perKey`，确保请求头存在对应键。
- Q：ApiKey 签名失败？
  - A：检查 `X-Timestamp` 是否在允许偏差内（默认 300s）；`X-Nonce` 是否重复；`bodyHash` 是否按预期计算。
- Q：事务未生效？
  - A：仅在配置 `transaction:{ ds }` 时路由自动插入；检查步骤中是否出现 `sql.*`。
- Q：上传解析失败？
  - A：非 `multipart/form-data` 时，框架将原始体视为单文件 `infile`；确保 `request.file.files` 字段声明与大小/类型约束正确。
- Q：OpenAPI 未包含请求体？
  - A：将参数来源声明在 `body/form/file` 中；路径/查询/头参数才会进入 `parameters`。