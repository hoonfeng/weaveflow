# 模块: subscriptions

## 接口: 创建订阅
- 方法: POST
- 路径: /api/subscriptions
- 说明: 为租户创建或切换订阅套餐
- 认证: jwt
- 角色: admin

### 请求参数
- [body] tenant_id (int) 必填:true 说明:租户ID
- [body] plan_id (int) 必填:true 说明:套餐ID

### 响应
- 状态码: 201

#### 示例数据
```
map[plan_id:{{ plan_id }} tenant_id:{{ tenant_id }}]
```

## 接口: 当前订阅
- 方法: GET
- 路径: /api/subscriptions/current
- 说明: 查询租户当前有效订阅与配额
- 认证: jwt
- 角色: admin

### 请求参数
- [query] tenant_id (int) 必填:true 说明:租户ID

### 响应
- 状态码: 200

#### 示例数据
```
{{ rows }}
```

