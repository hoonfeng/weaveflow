# 模块: plans

## 接口: 套餐列表
- 方法: GET
- 路径: /api/plans
- 说明: 列出所有可用套餐
- 认证: jwt
- 角色: admin

### 响应
- 状态码: 200
- 返回说明: 返回套餐列表

#### 返回参数
- code (int)
- msg (string)
- data (array)
  - item (object)
    - id (int)
    - name (string)
    - monthly_quota (int) — 月配额
    - price (number) — 单价

#### 示例数据
```json
{
  "code": 123,
  "data": [
    {
      "id": 123,
      "monthly_quota": 123,
      "name": "string",
      "price": 1.23
    }
  ],
  "msg": "string"
}
```

