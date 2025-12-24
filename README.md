# API

✔ 不写逻辑
✔ 不访问数据库

# Controller

✔ 只做：

- 参数校验
- 调用 service
- 返回统一响应

# Service

✔ 可以写事务
✔ 可以组合多个 DAO
✔ 不关心 HTTP

# Dao

✔ 不写业务判断
✔ 不处理参数合法性

# 数据库管理

```bash
# 生成框架
migrate create -ext sql -dir migrations -seq init

# 升级
migrate \
  -path migrations \
  -database "mysql://root:ossdbg1@tcp(127.0.0.1:3306)/yunyi?multiStatements=true" \
  up

# 回滚一版
migrate -path migrations -database "mysql://root:ossdbg1@tcp(127.0.0.1:3306)/yunyi?multiStatements=true" down 1
```
