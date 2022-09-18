# Task
面向医院的任务管理服务

## 缘由

## 技术选型
### 前提
1. 全新的团队，没有技术债务
2. 唯一后端开发（golang技术栈）
3. 功能明确，交付时间紧

### 基于前提
1. 选择golang
2. 尽量选择开箱即用，社区活跃的开源产品

### 对比
1. gin+自建脚手架 vs go-micro vs kratos vs go-zero

2. sql vs nosql
  结构化清晰，选用sql，几乎mysql就是唯一选择，就不详细对比了

3. 考虑缓存
  数据量不大，不考虑缓存

4. 考虑高可用
  仅供快速部署演示使用，仅考虑单节点部署

### 结论  

服务框架：kratos

数据库：mysql

部署：docker-compose

## quick start
```bash
docker-compose up -d
```

## 接口文档
[apifox](https://www.apifox.cn/apidoc/shared-b25823ed-651f-48bf-8edb-e23095887149)
或
[swagger](http://127.0.0.1:8000/q/swagger-ui)


## 项目结构
```bash
├── api                                   # 接口定义          
│   └── task                        # 任务接口
│       └── v1
├── bin                                   # 编译后的二进制文件
├── cmd                                   # 服务启动入口
│   └── task_manager                # 任务管理服务
├── configs                               # 配置文件
├── ent                                   # 数据库实体
│   ├── employee
│   ├── enttest
│   ├── hook
│   ├── hospital
│   ├── migrate
│   ├── predicate
│   ├── runtime
│   ├── schema
│   └── task
├── internal                              # 业务逻辑代码
│   ├── biz                         # 业务逻辑
│   ├── conf                        # 配置结构
│   ├── data                        # 数据访问层
│   ├── server                      # 服务启动
│   └── service                     # 服务实现
├── task                                  # swagger文档
│   └── v1
└── third_party                           # 第三方库
    ├── errors
    ├── google
    │   ├── api
    │   └── protobuf
    │       └── compiler
    ├── openapi
    │   └── v3
    └── validate
```

## 数据表设计


## 疑问
- Return a list of tasks created under a particular company.
    - 这里的company指的是医院吗？全文没有其他地方提到company了，最贴切的解释就是医院。

