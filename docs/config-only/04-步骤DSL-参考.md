# 步骤 DSL 参考（零代码）

- 通过 `steps:` 声明执行链；每步输出可写入 `ctx.vars`（指定 `out`）。
- 核心类型如下（仅列配置使用，结合示例即可落地）。

## 步骤执行原理

1. **上下文变量 (ctx.vars)**：每个步骤的输出可以通过 `out` 字段写入上下文变量
2. **执行顺序**：步骤按顺序执行，前一步的输出可作为后一步的输入
3. **错误处理**：任何步骤失败都会终止执行并返回错误响应
4. **模板函数**：支持使用 `{{ }}` 语法调用模板函数处理数据

## 步骤字段通用说明

所有步骤都支持以下通用字段：

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 否 | - | 输出变量名，将步骤结果保存到上下文变量中 |
| `if` | string | 否 | - | 条件表达式，只有表达式为真时才执行该步骤 |
| `name` | string | 否 | - | 步骤名称，用于调试和日志记录 |
| `timeout` | duration | 否 | - | 步骤执行超时时间，如 `30s`, `1m` |

## 响应 (response)
```yaml
- response:
    status: 200
    headers: { "X-Trace": "{{ uuid }}" }
    wrap: { code: 0, msg: ok, data: "{{ data }}" }
```

### response 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `status` | int | 是 | - | HTTP状态码，如200, 201, 400, 404, 500等 |
| `headers` | object | 否 | {} | 响应头，支持模板函数 |
| `body` | any | 否 | - | 响应体内容，支持模板函数 |
| `wrap` | object | 否 | - | 包装响应结构，自动添加code/msg/data字段 |
| `wrap.code` | int | 否 | 0 | 包装结构的状态码 |
| `wrap.msg` | string | 否 | "ok" | 包装结构的消息 |
| `wrap.data` | any | 否 | - | 包装结构的数据内容 |

**示例**：
```yaml
- response:
    status: 201
    headers:
      Location: "/api/users/{{ user.id }}"
      X-Request-ID: "{{ uuid }}"
    body: 
      id: "{{ user.id }}"
      name: "{{ user.name }}"

- response:
    status: 200
    wrap:
      code: 0
      msg: "success"
      data: "{{ users }}"
```

## 校验 (validate)
```yaml
- validate: { target: body.text, required: true, minLen: 1, maxLen: 10000 }
- validate: { target: file.files, maxSize: 32MB, types: ["image/png", "image/jpeg"] }
```

### validate 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `target` | string | 是 | - | 要校验的目标字段路径，如 `body.username`, `query.page`, `file.avatar` |
| `required` | bool | 否 | false | 是否必填 |
| `type` | string | 否 | - | 字段类型：`int`, `string`, `bool`, `float`, `array`, `object` |
| `min` | number | 否 | - | 最小值（数字类型） |
| `max` | number | 否 | - | 最大值（数字类型） |
| `minLen` | int | 否 | - | 最小长度（字符串、数组类型） |
| `maxLen` | int | 否 | - | 最大长度（字符串、数组类型） |
| `pattern` | string | 否 | - | 正则表达式模式 |
| `format` | string | 否 | - | 格式校验：`email`, `url`, `uuid`, `date`, `date-time` |
| `enum` | array | 否 | - | 枚举值列表 |
| `maxSize` | string | 否 | - | 最大文件大小：`B`, `KB`, `MB`, `GB` |
| `types` | array[string] | 否 | - | 允许的MIME类型列表（文件类型） |
| `message` | string | 否 | - | 自定义错误消息 |
| `code` | string | 否 | - | 自定义错误代码 |

**示例**：
```yaml
- validate:
    target: body.email
    required: true
    type: string
    format: email
    message: "邮箱格式不正确"
    code: "E_INVALID_EMAIL"

- validate:
    target: query.page
    type: int
    min: 1
    max: 1000

- validate:
    target: file.avatar
    required: true
    maxSize: "5MB"
    types: ["image/jpeg", "image/png"]
```

## 映射 (transform)
```yaml
- transform:
    mapping:
      user_id: "{{ toint(query.id) }}"
      now_ms:  "{{ now ms }}"
```

### transform 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `mapping` | object | 是 | - | 映射规则对象，键为变量名，值为模板表达式 |
| `mapping.*` | string | 是 | - | 模板表达式，支持所有模板函数 |

**示例**：
```yaml
- transform:
    mapping:
      # 类型转换
      user_id: "{{ toint(query.id) }}"
      is_admin: "{{ tobool(user.role == 'admin') }}"
      
      # 时间处理
      current_time: "{{ now }}"
      timestamp_ms: "{{ now ms }}"
      
      # 字符串操作
      search_term: "{{ lower(query.search) }}"
      email_domain: "{{ split(user.email, '@')[1] }}"
      
      # 数学运算
      total_price: "{{ mul(body.quantity, body.price) }}"
      discount_price: "{{ mul(total_price, sub(1, body.discount)) }}"
      
      # 数据结构操作
      user_names: "{{ map(users, 'name') }}"
      active_users: "{{ filter(users, 'status == \"active\"') }}"
```

## SQL 操作

### sql.query - SQL查询
```yaml
- sql.query:
    ds: main
    sql: "SELECT id,name FROM users WHERE id = ?"
    params: { id: "{{ toint(path.id) }}" }
    order: [id]
    out: user
```

#### sql.query 字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 数据源名称，在project.yaml中配置 |
| `sql` | string | 是 | - | SQL查询语句，支持参数占位符 `?` |
| `params` | object | 否 | {} | 参数对象，键值对形式 |
| `order` | array[string] | 否 | - | 参数顺序，确保参数按正确顺序绑定 |
| `out` | string | 否 | - | 输出变量名，查询结果保存到此变量 |
| `timeout` | duration | 否 | - | 查询超时时间 |
| `tx` | bool | 否 | false | 是否在事务中执行 |

### sql.exec - SQL执行
```yaml
- sql.exec:
    ds: main
    sql: "UPDATE users SET name=? WHERE id=?"
    params: { name: "{{ body.name }}", id: "{{ toint(path.id) }}" }
    order: [name, id]
```

#### sql.exec 字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 数据源名称 |
| `sql` | string | 是 | - | SQL执行语句，支持INSERT/UPDATE/DELETE |
| `params` | object | 否 | {} | 参数对象 |
| `order` | array[string] | 否 | - | 参数顺序 |
| `out` | string | 否 | - | 输出变量名，保存执行结果（影响行数） |
| `timeout` | duration | 否 | - | 执行超时时间 |
| `tx` | bool | 否 | false | 是否在事务中执行 |

**示例**：
```yaml
# 查询示例
- sql.query:
    ds: main
    sql: "SELECT id, name, email FROM users WHERE status = ? AND created_at > ?"
    params:
      status: "active"
      created_at: "{{ sub_days(now, 30) }}"
    order: [status, created_at]
    out: active_users
    timeout: 10s

# 执行示例  
- sql.exec:
    ds: main
    sql: "INSERT INTO users (name, email, status) VALUES (?, ?, ?)"
    params:
      name: "{{ body.name }}"
      email: "{{ body.email }}"
      status: "active"
    order: [name, email, status]
    out: insert_result

# 获取最后插入ID
- sql.query:
    ds: main
    sql: "SELECT LAST_INSERT_ID() as id"
    out: last_id
```

## KV 与 Cache

KV（键值存储）和Cache（缓存）操作支持Redis等数据源，用于高性能数据存取。

### kv.set - KV存储设置
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 数据源名称，在project.yaml中配置 |
| `key` | string | 是 | - | 键名，支持模板函数 |
| `value` | any | 是 | - | 值，支持模板函数 |
| `ttl` | duration | 否 | - | 过期时间，如 `30s`, `5m`, `1h` |
| `out` | string | 否 | - | 输出变量名，保存操作结果 |

### kv.get - KV存储获取
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 数据源名称，在project.yaml中配置 |
| `key` | string | 是 | - | 键名，支持模板函数 |
| `out` | string | 是 | - | 输出变量名，保存获取的值 |

### cache.set - 缓存设置
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `key` | string | 是 | - | 缓存键名，支持模板函数 |
| `value` | any | 是 | - | 缓存值，支持模板函数 |
| `ttl` | duration | 否 | - | 过期时间，如 `30s`, `5m`, `1h` |
| `out` | string | 否 | - | 输出变量名，保存操作结果 |

### cache.get - 缓存获取
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `key` | string | 是 | - | 缓存键名，支持模板函数 |
| `out` | string | 是 | - | 输出变量名，保存获取的缓存值 |

**示例**：
```yaml
# KV存储操作示例
- kv.set:
    ds: redis_cache
    key: "user_profile:{{ user.id }}"
    value: "{{ user }}"
    ttl: 1h
    out: set_result

- kv.get:
    ds: redis_cache  
    key: "user_profile:{{ path.user_id }}"
    out: cached_user

# 缓存操作示例
- cache.set:
    key: "feature_flag:new_ui"
    value: "enabled"
    ttl: 24h
    out: cache_set_result

- cache.get:
    key: "feature_flag:new_ui"
    out: ui_flag

# 条件缓存获取，如果缓存不存在则查询数据库
- branch:
    if: "{{ not cache.get({key: 'user:' + path.id, out: 'tmp_user'}) }}"
    then:
      - sql.query:
          ds: main
          sql: "SELECT * FROM users WHERE id = ?"
          args: [ "{{ path.id }}" ]
          out: user
      - cache.set:
          key: "user:{{ path.id }}"
          value: "{{ user }}"
          ttl: 5m
    else:
      - transform:
          mapping: { user: "{{ tmp_user }}" }
```

## 外部 HTTP（重试/断路器）

外部HTTP请求支持调用第三方API服务，包含重试机制和断路器模式。

### http.request 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `method` | string | 是 | - | HTTP方法：`GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD` |
| `url` | string | 是 | - | 请求URL，支持模板函数 |
| `headers` | object | 否 | {} | 请求头，支持模板函数 |
| `params` | object | 否 | {} | URL查询参数，支持模板函数 |
| `body` | any | 否 | - | 请求体，支持模板函数 |
| `retry` | int | 否 | 0 | 重试次数 |
| `backoff` | int | 否 | 100 | 重试间隔毫秒数 |
| `timeoutMs` | int | 否 | 5000 | 超时时间（毫秒） |
| `circuit` | object | 否 | - | 断路器配置 |
| `circuit.key` | string | 否 | - | 断路器唯一标识 |
| `circuit.threshold` | int | 否 | 3 | 失败阈值，达到后触发断路器 |
| `circuit.openMs` | int | 否 | 30000 | 断路器开启时间（毫秒） |
| `circuit.fallback` | any | 否 | - | 断路器触发时的降级响应 |
| `out` | string | 否 | - | 输出变量名，保存HTTP响应 |

**HTTP响应结构**：
```yaml
# HTTP响应包含以下字段
status: 200          # HTTP状态码
headers: {}         # 响应头
body: any           # 响应体
```

**示例**：
```yaml
# 基本GET请求
- http.request:
    method: GET
    url: "https://api.example.com/users/{{ path.id }}"
    headers:
      Authorization: "Bearer {{ token }}"
      X-Request-ID: "{{ uuid }}"
    out: user_data

# POST请求带JSON body
- http.request:
    method: POST
    url: "https://api.example.com/users"
    headers:
      Content-Type: "application/json"
      Authorization: "Bearer {{ token }}"
    body:
      name: "{{ body.name }}"
      email: "{{ body.email }}"
    out: create_result

# 带重试和断路器的请求
- http.request:
    method: GET
    url: "https://api.example.com/health"
    retry: 3
    backoff: 500
    timeoutMs: 3000
    circuit:
      key: "external_api_health"
      threshold: 5
      openMs: 60000
      fallback:
        status: 200
        body: { status: "degraded", message: "Service temporarily unavailable" }
    out: health_status

# 处理HTTP响应
- branch:
    if: "{{ eq(get(http_response,'status'), 200) }}"
    then:
      - transform:
          mapping:
            user: "{{ get(http_response,'body') }}"
    else:
      - response:
          status: 502
          body:
            error: "upstream_error"
            message: "External service unavailable"
            details: "{{ get(http_response,'body') }}"
```

## 上传保存

文件上传保存操作支持将上传的文件保存到本地或云存储。

### upload.save 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 存储配置名称，在project.yaml中配置 |
| `input` | string | 是 | - | 输入文件字段，如 `file.files`, `file.avatar` |
| `naming` | string | 否 | "original" | 文件名生成策略：`original`, `random`, `sha256`, `md5` |
| `out` | string | 否 | - | 输出变量名，保存上传结果 |
| `path` | string | 否 | - | 存储路径，支持模板函数 |
| `maxSize` | string | 否 | - | 最大文件大小限制，如 `10MB`, `100KB` |
| `allowedTypes` | array[string] | 否 | - | 允许的MIME类型列表 |

**上传结果结构**：
```yaml
# 上传结果包含以下字段
id: string           # 文件唯一标识
name: string         # 保存的文件名
path: string         # 文件存储路径
size: int            # 文件大小（字节）
mime: string         # MIME类型
url: string          # 文件访问URL（如果配置了访问地址）
```

**示例**：
```yaml
# 基本文件上传
- upload.save:
    ds: local_storage
    input: "{{ file.avatar }}"
    naming: sha256
    out: avatar_result

# 带路径和类型限制的上传
- upload.save:
    ds: cloud_storage
    input: "{{ file.document }}"
    naming: random
    path: "documents/{{ date '2006/01/02' }}"
    maxSize: "10MB"
    allowedTypes: ["application/pdf", "application/msword"]
    out: doc_result

# 多文件上传处理
- upload.save:
    ds: local
    input: "{{ file.files }}"
    naming: sha256
    out: uploaded_files

# 使用上传结果
- transform:
    mapping:
      file_id: "{{ get(uploaded_files,'id') }}"
      file_url: "{{ get(uploaded_files,'url') }}"

- response:
    status: 200
    body:
      success: true
      file_id: "{{ file_id }}"
      download_url: "{{ file_url }}"
```

## 控制流

控制流步骤支持条件分支和循环操作，实现复杂的业务逻辑。

### branch - 条件分支
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `if` | string | 是 | - | 条件表达式，支持模板函数 |
| `then` | array | 是 | - | 条件为真时执行的步骤列表 |
| `else` | array | 否 | - | 条件为假时执行的步骤列表 |
| `out` | string | 否 | - | 输出变量名，保存分支执行结果 |

### loop - 循环
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `items` | any | 是 | - | 要遍历的数组或对象，支持模板函数 |
| `var` | string | 是 | - | 循环变量名，表示当前迭代项 |
| `do` | array | 是 | - | 每次迭代执行的步骤列表 |
| `out` | string | 否 | - | 输出变量名，保存循环执行结果 |
| `indexVar` | string | 否 | - | 索引变量名，表示当前迭代索引 |
| `concurrency` | int | 否 | 1 | 并发执行数，大于1时并行执行 |

**示例**：
```yaml
# 条件分支示例
- branch:
    if: "{{ and(gt(len(users),0), eq(get(users.0,'status'),'active')) }}"
    then:
      - response:
          status: 200
          body: "{{ users }}"
    else:
      - response:
          status: 404
          body:
            error: "not_found"
            message: "No active users found"

# 嵌套条件分支
- branch:
    if: "{{ gt(query.page,0) }}"
    then:
      - sql.query:
          ds: main
          sql: "SELECT * FROM users LIMIT 10 OFFSET ?"
          args: [ "{{ mul(sub(query.page,1),10) }}" ]
          out: paged_users
      - branch:
          if: "{{ gt(len(paged_users),0) }}"
          then:
            - response:
                status: 200
                body: "{{ paged_users }}"
          else:
            - response:
                status: 404
                body: { error: "no_more_data" }
    else:
      - response:
          status: 400
          body: { error: "invalid_page" }

# 循环示例 - 数组遍历
- loop:
    items: "{{ user_ids }}"
    var: user_id
    do:
      - sql.query:
          ds: main
          sql: "SELECT * FROM users WHERE id = ?"
          args: [ "{{ user_id }}" ]
          out: user
      - transform:
          mapping:
            processed_users: "{{ concat(vars.processed_users, user) }}"
    out: all_users

# 循环示例 - 对象遍历
- loop:
    items: "{{ user_roles }}"
    var: role_info
    indexVar: idx
    do:
      - transform:
          mapping:
            role_name: "{{ get(role_info,'key') }}"
            user_count: "{{ get(role_info,'value') }}"
            display_order: "{{ idx }}"
      - sql.exec:
          ds: main
          sql: "INSERT INTO role_stats (role_name, user_count, display_order) VALUES (?, ?, ?)"
          args: [ "{{ role_name }}", "{{ user_count }}", "{{ display_order }}" ]

# 并发循环示例
- loop:
    items: "{{ image_urls }}"
    var: img_url
    concurrency: 5
    do:
      - http.request:
          method: GET
          url: "{{ img_url }}"
          out: image_data
      - transform:
          mapping:
            downloaded_images: "{{ concat(vars.downloaded_images, get(image_data,'body')) }}"
    out: all_images

# 使用循环结果
- response:
    status: 200
    body:
      total: "{{ len(all_images) }}"
      images: "{{ all_images }}"
```

## 插件调用

插件调用步骤支持调用外部插件服务，扩展框架功能。

### plugin.call 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `plugin` | string | 是 | - | 插件名称，在project.yaml中配置 |
| `function` | string | 是 | - | 要调用的插件函数名 |
| `params` | object | 否 | {} | 调用参数，支持模板函数 |
| `retry` | int | 否 | 0 | 重试次数 |
| `backoff` | int | 否 | 100 | 重试间隔毫秒数 |
| `timeoutMs` | int | 否 | 5000 | 超时时间（毫秒） |
| `circuit` | object | 否 | - | 断路器配置 |
| `circuit.key` | string | 否 | - | 断路器唯一标识 |
| `circuit.threshold` | int | 否 | 3 | 失败阈值，达到后触发断路器 |
| `circuit.openMs` | int | 否 | 30000 | 断路器开启时间（毫秒） |
| `circuit.fallback` | any | 否 | - | 断路器触发时的降级响应 |
| `out` | string | 否 | - | 输出变量名，保存插件调用结果 |

**示例**：
```yaml
# 基本插件调用
- plugin.call:
    plugin: image_processor
    function: resize
    params:
      image: "{{ body.image_data }}"
      width: 300
      height: 200
      quality: 85
    out: resized_image

# 带重试和断路器的插件调用
- plugin.call:
    plugin: payment_gateway
    function: create_charge
    params:
      amount: "{{ body.amount }}"
      currency: "{{ body.currency }}"
      customer: "{{ body.customer_id }}"
    retry: 2
    backoff: 500
    timeoutMs: 10000
    circuit:
      key: "payment_service"
      threshold: 5
      openMs: 60000
      fallback:
        status: "failed"
        error: "service_unavailable"
    out: payment_result

# 处理插件调用结果
- branch:
    if: "{{ eq(get(payment_result,'status'),'succeeded') }}"
    then:
      - sql.exec:
          ds: main
          sql: "INSERT INTO payments (id, amount, status) VALUES (?, ?, ?)"
          args: [ "{{ get(payment_result,'id') }}", "{{ get(payment_result,'amount') }}", "completed" ]
      - response:
          status: 200
          body:
            success: true
            payment_id: "{{ get(payment_result,'id') }}"
    else:
      - response:
          status: 400
          body:
            success: false
            error: "{{ get(payment_result,'error') }}"
            message: "Payment processing failed"

# 链式插件调用
- plugin.call:
    plugin: validation_service
    function: validate_email
    params:
      email: "{{ body.email }}"
    out: validation_result

- branch:
    if: "{{ get(validation_result,'is_valid') }}"
    then:
      - plugin.call:
          plugin: notification_service
          function: send_welcome
          params:
            email: "{{ body.email }}"
            name: "{{ body.name }}"
        out: notification_result
    else:
      - response:
          status: 400
          body:
            error: "invalid_email"
            message: "Please provide a valid email address"
```

## 插件管理（运行时）

插件管理步骤支持运行时管理插件状态和配置。

### plugins.status - 获取插件状态
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存插件状态信息 |

### plugins.control - 插件控制
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `action` | string | 是 | - | 控制动作：`enable`, `disable` |
| `names` | array[string] | 是 | - | 插件名称列表 |
| `enabled` | bool | 否 | - | 启用状态（仅对enable/disable动作有效） |
| `out` | string | 否 | - | 输出变量名，保存控制结果 |

### plugins.add - 添加插件
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `name` | string | 是 | - | 插件名称 |
| `executable` | object | 是 | - | 可执行文件路径配置 |
| `executable.windows` | string | 否 | - | Windows平台可执行文件路径 |
| `executable.unix` | string | 否 | - | Unix/Linux平台可执行文件路径 |
| `instances` | int | 否 | 1 | 插件实例数 |
| `timeout` | duration | 否 | "300ms" | 调用超时时间 |
| `queueSize` | int | 否 | 1024 | 请求队列大小 |
| `functions` | array[string] | 否 | - | 支持的函数列表 |
| `config` | object | 否 | {} | 插件配置参数 |
| `env` | object | 否 | {} | 环境变量配置 |
| `out` | string | 否 | - | 输出变量名，保存添加结果 |

### plugins.remove - 移除插件
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `names` | array[string] | 是 | - | 要移除的插件名称列表 |
| `out` | string | 否 | - | 输出变量名，保存移除结果 |

### plugins.restart - 重启插件
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `names` | array[string] | 是 | - | 要重启的插件名称列表 |
| `out` | string | 否 | - | 输出变量名，保存重启结果 |

### plugins.stop - 停止插件
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `names` | array[string] | 是 | - | 要停止的插件名称列表 |
| `out` | string | 否 | - | 输出变量名，保存停止结果 |

### plugins.start - 启动插件
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `names` | array[string] | 是 | - | 要启动的插件名称列表 |
| `out` | string | 否 | - | 输出变量名，保存启动结果 |

**插件状态结构**：
```yaml
# 插件状态包含以下字段
name: string          # 插件名称
status: string        # 状态：running, stopped, error
enabled: bool         # 是否启用
error: string         # 错误信息（如果有）
pid: int              # 进程ID（如果运行中）
start_time: string    # 启动时间
instances: int        # 实例数
queue_size: int       # 队列大小
functions: array      # 支持的函数列表
```

**示例**：
```yaml
# 获取所有插件状态
- plugins.status:
    out: all_plugins

# 启用特定插件
- plugins.control:
    action: enable
    names: ["image_processor", "payment_gateway"]
    enabled: true
    out: enable_result

# 添加新插件
- plugins.add:
    name: "analytics_service"
    executable:
      windows: "plugins/analytics.exe"
      unix: "./plugins/analytics"
    instances: 2
    timeout: "500ms"
    queueSize: 2048
    functions: ["track_event", "get_stats"]
    config:
      api_key: "{{ env.ANALYTICS_API_KEY }}"
      endpoint: "https://analytics.example.com"
    env:
      LOG_LEVEL: "info"
      MAX_WORKERS: "10"
    out: add_result

# 重启插件
- plugins.restart:
    names: ["notification_service"]
    out: restart_result

# 停止插件
- plugins.stop:
    names: ["debug_tool"]
    out: stop_result

# 启动插件
- plugins.start:
    names: ["image_processor"]
    out: start_result

# 移除插件
- plugins.remove:
    names: ["old_service"]
    out: remove_result

# 使用插件状态信息
- branch:
    if: "{{ eq(get(get(all_plugins,'image_processor'),'status'),'running') }}"
    then:
      - response:
          status: 200
          body:
            message: "All required plugins are running"
            plugins: "{{ all_plugins }}"
    else:
      - response:
          status: 503
          body:
            error: "service_unavailable"
            message: "Image processor plugin is not running"
            plugin_status: "{{ get(all_plugins,'image_processor') }}"

# 插件健康检查
- plugins.status:
    out: plugin_statuses
- branch:
    if: "{{ and(eq(get(get(plugin_statuses,'payment_gateway'),'status'),'running'), eq(get(get(plugin_statuses,'notification_service'),'status'),'running')) }}"
    then:
      - response:
          status: 200
          body: { healthy: true }
    else:
      - response:
          status: 503
          body:
            healthy: false
            failed_plugins: "{{ filter(plugin_statuses, func(k,v) { return ne(v.status,'running') }) }}"
```

## 向量库

向量库操作支持向量相似性搜索和向量数据管理。

### vector.search - 向量搜索
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 向量数据源名称，在project.yaml中配置 |
| `collection` | string | 是 | - | 集合名称 |
| `vec` | array[float] | 是 | - | 查询向量，支持模板函数 |
| `topK` | int | 否 | 10 | 返回最相似的前K个结果 |
| `filter` | object | 否 | - | 过滤条件 |
| `out` | string | 是 | - | 输出变量名，保存搜索结果 |

### vector.upsert - 向量插入/更新
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 向量数据源名称，在project.yaml中配置 |
| `collection` | string | 是 | - | 集合名称 |
| `id` | string | 是 | - | 向量ID |
| `vec` | array[float] | 是 | - | 向量数据，支持模板函数 |
| `meta` | object | 否 | - | 元数据 |
| `out` | string | 否 | - | 输出变量名，保存操作结果 |

### vector.delete - 向量删除
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 向量数据源名称，在project.yaml中配置 |
| `collection` | string | 是 | - | 集合名称 |
| `id` | string | 是 | - | 要删除的向量ID |
| `out` | string | 否 | - | 输出变量名，保存删除结果 |

### vector.ensure - 确保集合存在
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 向量数据源名称，在project.yaml中配置 |
| `collection` | string | 是 | - | 集合名称 |
| `size` | int | 是 | - | 向量维度大小 |
| `metric` | string | 否 | "Cosine" | 相似度度量：`Cosine`, `Euclidean`, `Dot` |
| `out` | string | 否 | - | 输出变量名，保存操作结果 |

### vector.upsert_batch - 批量向量插入/更新
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 向量数据源名称，在project.yaml中配置 |
| `collection` | string | 是 | - | 集合名称 |
| `items` | array[object] | 是 | - | 批量数据，每个对象包含id, vec, meta字段 |
| `out` | string | 否 | - | 输出变量名，保存操作结果 |

**搜索结果结构**：
```yaml
# 搜索结果包含以下字段
id: string           # 向量ID
score: float         # 相似度分数
distance: float      # 距离值
meta: object         # 元数据
```

**示例**：
```yaml
# 向量相似性搜索
- vector.search:
    ds: vector_db
    collection: "documents"
    vec: "{{ body.embedding }}"
    topK: 5
    filter:
      category: "technology"
      published: true
    out: search_results

# 插入单个向量
- vector.upsert:
    ds: vector_db
    collection: "documents"
    id: "doc_123"
    vec: "{{ body.embedding }}"
    meta:
      title: "{{ body.title }}"
      category: "{{ body.category }}"
      published: true
    out: upsert_result

# 删除向量
- vector.delete:
    ds: vector_db
    collection: "documents"
    id: "doc_456"
    out: delete_result

# 确保集合存在
- vector.ensure:
    ds: vector_db
    collection: "images"
    size: 512
    metric: "Euclidean"
    out: ensure_result

# 批量插入向量
- vector.upsert_batch:
    ds: vector_db
    collection: "documents"
    items:
      - id: "doc_001"
        vec: [0.1, 0.2, 0.3, 0.4]
        meta: { title: "Document 1", category: "tech" }
      - id: "doc_002"
        vec: [0.5, 0.6, 0.7, 0.8]
        meta: { title: "Document 2", category: "science" }
      - id: "doc_003"
        vec: [0.9, 1.0, 1.1, 1.2]
        meta: { title: "Document 3", category: "tech" }
    out: batch_result

# 处理搜索结果
- branch:
    if: "{{ gt(len(search_results),0) }}"
    then:
      - response:
          status: 200
          body:
            results: "{{ search_results }}"
            count: "{{ len(search_results) }}"
    else:
      - response:
          status: 404
          body:
            error: "no_results"
            message: "No similar documents found"

# 语义搜索管道
- transform:
    mapping:
      query_text: "{{ body.query }}"
- plugin.call:
    plugin: embedding_service
    function: get_embedding
    params:
      text: "{{ query_text }}"
    out: query_embedding
- vector.search:
    ds: vector_db
    collection: "documents"
    vec: "{{ get(query_embedding,'embedding') }}"
    topK: 10
    out: similar_docs
- response:
    status: 200
    body:
      query: "{{ query_text }}"
      results: "{{ similar_docs }}"
```

## 模型应用

模型应用步骤支持应用数据库模型定义到数据库。

### model.apply 步骤字段说明

| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `ds` | string | 是 | - | 数据源名称，在project.yaml中配置 |
| `dir` | string | 是 | - | 模型定义文件目录路径 |
| `out` | string | 否 | - | 输出变量名，保存应用结果 |
| `dryRun` | bool | 否 | false | 是否只预览不实际执行 |
| `force` | bool | 否 | false | 是否强制执行，忽略警告 |

**应用结果结构**：
```yaml
# 应用结果包含以下字段
applied: int           # 成功应用的变更数量
skipped: int          # 跳过的变更数量
errors: array         # 错误信息列表（如果有）
warnings: array       # 警告信息列表
changes: array        # 详细变更信息
```

**示例**：
```yaml
# 应用模型到数据库
- model.apply:
    ds: main
    dir: "configs/models"
    out: apply_result

# 干运行模式，只预览变更
- model.apply:
    ds: main
    dir: "configs/models"
    dryRun: true
    out: preview_result

# 强制应用，忽略警告
- model.apply:
    ds: main
    dir: "configs/models"
    force: true
    out: force_result

# 处理应用结果
- branch:
    if: "{{ eq(get(apply_result,'errors'),0) }}"
    then:
      - response:
          status: 200
          body:
            success: true
            applied: "{{ get(apply_result,'applied') }}"
            message: "Database schema updated successfully"
    else:
      - response:
          status: 500
          body:
            success: false
            errors: "{{ get(apply_result,'errors') }}"
            message: "Failed to apply database schema changes"

# 模型应用管道
- model.apply:
    ds: main
    dir: "configs/models"
    out: schema_result
- branch:
    if: "{{ gt(get(schema_result,'applied'),0) }}"
    then:
      - sql.query:
          ds: main
          sql: "SELECT COUNT(*) as table_count FROM information_schema.tables WHERE table_schema = DATABASE()"
          out: table_count
      - response:
          status: 200
          body:
            applied_changes: "{{ get(schema_result,'applied') }}"
            total_tables: "{{ get(table_count.0,'table_count') }}"
    else:
      - response:
          status: 200
          body:
            message: "No schema changes needed"
            tables: "{{ get(schema_result,'skipped') }}"
```

## 管理与观测

管理与观测步骤提供系统运维、监控和诊断功能，支持零代码方式管理框架运行状态。

### admin.docs - 获取接口文档
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存接口文档信息 |

**文档结果结构**：
```yaml
# 文档结果包含以下字段
endpoints: array[object]  # 接口端点列表
  - path: string          # 接口路径
  - method: string        # HTTP方法
  - module: string        # 模块名称
  - title: string         # 接口标题
  - description: string   # 接口描述
  - auth: string          # 认证方式
  - roles: array[string]  # 访问角色
  - permissions: array[string]  # 所需权限
```

### admin.openapi - 生成OpenAPI规范
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存OpenAPI规范 |

**OpenAPI结果结构**：
```yaml
# 返回完整的OpenAPI 3.0规范文档
openapi: "3.0.3"
info:
  title: "接口文档"
  version: "v1"
paths: object  # 所有接口路径定义
components: object  # 组件定义（schemas、securitySchemes等）
```

### admin.builtin - 获取内置能力列表
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存内置能力信息 |

**内置能力结果结构**：
```yaml
# 内置能力列表
builtins: array[object]  # 内置能力列表
  - name: string         # 能力名称
  - desc: string         # 能力描述
  - category: string     # 分类（响应、校验、数据访问等）
```

### plugins.status - 获取插件状态
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存插件状态信息 |

**插件状态结果结构**：
```yaml
# 插件状态信息
plugins: array[object]  # 插件状态列表
  - name: string        # 插件名称
  - enabled: bool       # 是否启用
  - running: bool       # 是否运行中
  - functions: array[string]  # 支持的函数列表
  - instances: int      # 实例数量
  - queueSize: int      # 队列大小
  - timeout: duration   # 调用超时时间
```

### admin.permissions_scan - 扫描权限声明
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存权限扫描结果 |

**权限扫描结果结构**：
```yaml
# 权限扫描结果
permissions: array[object]  # 权限声明列表
  - name: string           # 权限名称
  - description: string    # 权限描述
  - interfaces: array[string]  # 使用该权限的接口列表
  - modules: array[string] # 使用该权限的模块列表
```

### admin.plugins_usage - 分析插件使用情况
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存插件使用分析结果 |

**插件使用分析结果结构**：
```yaml
# 插件使用分析结果
internal: array[object]    # 内部插件使用情况
  - name: string          # 插件名称
  - endpoints: array[string]  # 使用的端点列表
external: array[object]    # 外部插件使用情况
  - name: string          # 插件名称
  - functions: array[string]  # 使用的函数列表
  - callCount: int        # 调用次数统计
```

### admin.lint - 接口配置检查
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存检查结果 |

**检查结果结构**：
```yaml
# 接口配置检查结果
issues: array[object]  # 问题列表
  - level: string      # 问题级别（error, warning, info）
  - message: string    # 问题描述
  - path: string       # 接口路径
  - module: string     # 模块名称
  - suggestion: string # 修复建议
```

### admin.reload - 热重载配置
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存重载结果 |

**重载结果结构**：
```yaml
# 热重载结果
success: bool          # 是否成功
reloaded: array[string]  # 重新加载的配置列表
  - interfaces: int    # 接口数量
  - plugins: int       # 插件数量
  - modules: int       # 模块数量
error: string          # 错误信息（如果失败）
```

### obs.metrics - 获取运行指标
| 字段 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| `out` | string | 是 | - | 输出变量名，保存指标数据 |

**指标结果结构**：
```yaml
# Prometheus格式的指标数据
metrics: string  # Prometheus文本格式的指标数据，包含：
  - http_requests_total  # HTTP请求总数
  - http_request_duration_seconds  # 请求耗时
  - plugin_calls_total   # 插件调用次数
  - sql_queries_total    # SQL查询次数
  - cache_hits_total     # 缓存命中次数
  - active_connections   # 活跃连接数
```

**示例**：
```yaml
# 获取系统文档
- admin.docs:
    out: api_docs

# 生成OpenAPI规范
- admin.openapi:
    out: openapi_spec

# 获取内置能力列表
- admin.builtin:
    out: builtin_capabilities

# 检查插件状态
- plugins.status:
    out: plugin_status

# 扫描权限声明
- admin.permissions_scan:
    out: permission_scan

# 分析插件使用情况
- admin.plugins_usage:
    out: plugin_usage

# 检查接口配置规范
- admin.lint:
    out: lint_results

# 热重载配置
- admin.reload:
    out: reload_result

# 获取运行指标
- obs.metrics:
    out: system_metrics

# 使用管理信息构建响应
- response:
    status: 200
    body:
      docs: "{{ api_docs }}"
      openapi: "{{ openapi_spec }}"
      builtins: "{{ builtin_capabilities }}"
      plugins: "{{ plugin_status }}"
      permissions: "{{ permission_scan }}"
      metrics: "{{ system_metrics }}"
```

## 模板函数（示例）
- 数值/聚合：`len|count|sum|avg|min|max|round|floor|ceil|abs|pow|add|sub|mul|div|mod`
- 比较/逻辑：`eq|ne|gt|ge|lt|le|and|or|not|if|coalesce`
- 字符串/集合：`upper|lower|join|split|format|concat|unique|slice|range|union|intersect|diff|keys|values|get|indexby|map|filter|sort|sortn|groupby|sumby|avgby|minby|maxby|uniqby|pick|omit|merge|flatten|compact|contains|startswith|endswith|replace`
- 编解码/时间：`json_encode|json_decode|base64_encode|base64_decode|url_encode|url_decode|now|format_time|parse_time|add_duration`
- 散列/安全：`sha256|sha256_concat|uuid|hmac_sha256`