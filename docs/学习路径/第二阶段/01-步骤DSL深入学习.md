# 第二阶段：核心功能掌握（3-5天）

## 2.1 步骤DSL深入学习

### 目标
掌握核心步骤DSL的使用，包括响应构建、参数校验和数据转换。

### 2.1.1 步骤DSL概述

步骤DSL是框架的核心，用于定义接口的处理逻辑。每个步骤按顺序执行，前一步骤的输出可以作为后一步骤的输入。

**基本语法：**
```yaml
steps:
  - step_type:       # 步骤类型
      param1: value1 # 步骤参数
      param2: value2
  - step_type:
      param3: value3
```

### 2.1.2 response - 响应构建

**构建HTTP响应：**
```yaml
steps:
  - response:
      status: 200                 # HTTP状态码
      headers:                   # 响应头
        Content-Type: application/json
        X-Custom-Header: value
      body:                      # 响应体（对象或字符串）
        code: 0
        message: "success"
        data: "{{ previous_step_output }}"
      wrap:                      # 包装响应（推荐使用）
        code: 0
        msg: "success"
        data: "{{ previous_step_output }}"
```

**响应示例：**
```yaml
# 返回简单文本
steps:
  - response:
      status: 200
      body: "Hello, World!"

# 返回JSON对象
steps:
  - response:
      status: 200
      body:
        message: "操作成功"
        timestamp: "{{ now }}"

# 使用wrap包装响应
steps:
  - response:
      status: 200
      wrap:
        code: 0
        msg: "success"
        data:
          user:
            name: "John"
            email: "john@example.com"
```

### 2.1.3 validate - 参数校验

**单目标校验（与引擎实现一致）：**
```yaml
steps:
  - validate:
      target: query.page
      type: int
      min: 1
  - validate:
      target: path.id
      type: int
      required: true
      min: 1
  - validate:
      target: body.email
      type: string
      regex: "^\\S+@\\S+\\.\\S+$"
  - validate:
      target: file.files
      required: true
      maxSize: 10MB
      types: ["image/png",".jpg"]
```

**校验规则要点：**
- `target` 取值路径（query/path/body/ctx.vars）；其余约束作用于该值
- 文件校验支持 `maxSize` 与 `types`（MIME 与扩展名均可），实现参考 internal/core/engine.go:74–91

### 2.1.4 transform - 数据转换

**数据映射和转换：**
```yaml
steps:
  - transform:
      mapping:                   # 字段映射
        fullName: "{{ body.firstName }} {{ body.lastName }}"
        userAge: "{{ toint(body.age) }}"
        isAdult: "{{ toint(body.age) >= 18 }}"
        formattedTime: "{{ format_time(now, '2006-01-02') }}"
      out: transformed_data     # 输出变量名
```

**转换/模板函数（节选，与引擎一致）：**
- `toint/tofloat/tostring/tobool`
- `upper/lower/json_encode/sha256/sha256_concat/uuid`
- 数学：`add/sub/mul/div/mod/sum/avg/min/max`
- 字符串：`join/split/format/concat`
- 时间：`now`（`unix/ms/RFC3339`），`format_time`，`parse_time`，`add_duration`

**转换示例：**
```yaml
# 用户信息转换
steps:
  - transform:
      mapping:
        userId: "{{ toint(path.id) }}"
        userInfo:
          name: "{{ upper(body.name) }}"
          email: "{{ lower(body.email) }}"
          age: "{{ toint(body.age) }}"
          isAdult: "{{ toint(body.age) >= 18 }}"
          registrationDate: "{{ format_time(now, '2006-01-02 15:04:05') }}"
```

### 2.1.5 分支与循环

**条件分支：**
```yaml
steps:
  - branch:
      if: "{{ gt(toint(query.page), 1) }}"
      then:
        - transform: { mapping: { note: "page>1" } }
      else:
        - transform: { mapping: { note: "page<=1" } }
```

**循环处理：**
```yaml
steps:
  - loop:
      items: "{{ body.items }}"
      var: item
      do:
        - transform:
            mapping:
              item_id: "{{ toint(item.id) }}"
      out: user_data
```

### 2.1.5 步骤组合示例

**完整的用户注册接口：**
```yaml
module: auth
endpoint: register
method: POST
path: /api/auth/register

steps:
  # 1. 参数校验
  - validate:
      body:
        username: { type: string, required: true, min: 3, max: 20 }
        password: { type: string, required: true, min: 6, max: 50 }
        email: { type: string, required: true, pattern: "^\\S+@\\S+\\.\\S+$" }
        age: { type: int, required: false, min: 0, max: 150 }

  # 2. 数据转换
  - transform:
      mapping:
        hashedPassword: "{{ sha256(body.password) }}"
        normalizedEmail: "{{ lower(body.email) }}"
        userAge: "{{ toint(default(body.age, 0)) }}"
        isAdult: "{{ toint(default(body.age, 0)) >= 18 }}"
      out: processed_data

  # 3. 构建响应
  - response:
      status: 201
      wrap:
        code: 0
        msg: "用户注册成功"
        data:
          username: "{{ body.username }}"
          email: "{{ processed_data.normalizedEmail }}"
          age: "{{ processed_data.userAge }}"
          status: "active"
          registeredAt: "{{ now }}"
```

### 2.1.6 实践任务

**完成以下实践任务：**

1. **创建参数校验接口**
   - 创建一个POST接口，接收用户信息
   - 校验用户名、密码、邮箱格式
   - 校验年龄范围（0-150）

2. **创建数据转换接口**
   - 创建一个接口，接收原始数据
   - 使用transform步骤进行数据转换和格式化
   - 返回处理后的数据

3. **组合使用多个步骤**
   - 创建完整的处理流程：校验 → 转换 → 响应
   - 使用模板函数进行数据处理

### 2.1.7 常见问题

**Q: 校验失败时会发生什么？**
A: 校验失败会返回400错误，包含详细的错误信息。

**Q: 如何访问前一步骤的输出？**
A: 使用 `{{ previous_step_output }}` 或步骤的 `out` 参数指定的变量名。

**Q: transform步骤的out参数是必须的吗？**
A: 是的，需要指定输出变量名以便后续步骤使用。

### 2.1.8 下一步学习

掌握步骤DSL基础后，继续学习：[数据操作能力](./02-数据操作能力.md) 掌握SQL操作和数据库交互。