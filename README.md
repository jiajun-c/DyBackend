# DyBackend

## 1. 技术选型

| 框架名            | 用途      | 文档                                                            |
|----------------|---------|---------------------------------------------------------------|
| Kitex          | 微服务框架   | https://www.cloudwego.io/zh/docs/kitex/getting-started/       | 
| hertz          | web框架   | https://github.com/cloudwego/hertz/blob/develop/README_cn.md  |
| gin            | web框架   | https://gin-gonic.com/zh-cn/docs/                             |
| gorm           | ORM框架   | https://gorm.io/zh_CN/docs/index.html                         |
| go-redis       | 中间件     | https://redis.uptrace.dev/                                    |
| Mysql          | 数据库     ||
| Redis          | 缓存      | https://www.redis.com.cn/documentation.html                   |
| MinIO          | 高性能对象存储 | http://docs.minio.org.cn/docs/                                |
| etcd           | 服务发现    | https://doczhcn.gitbook.io/etcd/index                         |
| docker-compose | 部署      | https://yeasy.gitbook.io/docker_practice/compose/compose_file |



## 2. 开发规范

### 2.1 git

使用rebase进行分支合并，[文档](https://git-scm.com/book/zh/v2/Git-%E5%88%86%E6%94%AF-%E5%8F%98%E5%9F%BA)
### 2.2 变量命名规范
- 私有变量和临时变量一律使用小驼峰式命名法，形如"homeController"，公有变量一律使用大驼峰式命名法，形如："UseControllers"
- 常量使用全大写，不同单词之间用_分开(大写帕斯卡式命名法)
- 常见缩写保证大小写一致，比如SQL，URL，别出现Sql，Url这种写法
### 2.3 API规范
[规范接口](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)
### 2.4 文件/包命名规范
- 包名一律使用小写单数形式
- 包名尽可能简短
- go文件使用帕斯卡式命名法

## 3. 项目架构

```text
.
├── cmd // 微服务分层
│   ├── user
│   └── video
├── config  // 配置文件
│   ├── app // 我们每个微服务的配置文件
│   └── redis
├── data    // docker-compose目录映射用的
│   └── etcd  
├── docker-compose.yml
├── go.mod
├── go.sum
├── idl    // idl文件位置
│   └── user.proto
├── internal // 内部包
│   └── config
├── log      // 日志信息
│   └── redis
├── middleware  // 项目中间件
└── README.md
```
### 3.1 微服务内部分层
- dal层：进行数据库的基础操作
- rpc层：进行rpc调用
- service层：业务层，编写业务逻辑

## 4. 运行项目
### 1.环境
暂时只依赖mysql

数据库存放在云服务器上，配置文件中已经写出

### 2. 微服务
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```


