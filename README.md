# Task
面向医院的任务管理服务

## 技术选型
服务框架：kratos
数据库：mysql

## quick start
```bash
docker-compose up -d
```

## 接口文档
[apifox](https://www.apifox.cn/apidoc/shared-b25823ed-651f-48bf-8edb-e23095887149)

## 项目结构
```bash
├── api                      # 接口定义          
│   └── task           # 任务接口
│       └── v1
├── bin                      # 编译后的二进制文件
├── cmd                      # 服务启动入口
│   └── task_manager   # 任务管理服务
├── configs                  # 配置文件
├── ent                      # 数据库实体
│   ├── employee
│   ├── enttest
│   ├── hook
│   ├── hospital
│   ├── migrate
│   ├── predicate
│   ├── runtime
│   ├── schema
│   └── task
├── internal                 # 业务逻辑代码
│   ├── biz            # 业务逻辑
│   ├── conf           # 配置结构
│   ├── data           # 数据访问层
│   ├── server         # 服务启动
│   └── service        # 服务实现
├── task                     # openapi文档
│   └── v1
└── third_party              # 第三方库
    ├── errors
    ├── google
    │   ├── api
    │   └── protobuf
    │       └── compiler
    ├── openapi
    │   └── v3
    └── validate
```

## 疑问
- Return a list of tasks created under a particular company.
    - 这里的company指的是医院吗？全文没有其他地方提到company了，最贴切的解释就是医院。

