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
