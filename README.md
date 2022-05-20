# 基于Go语言实现的RESTful API

**注：本项目以上线本人阿里云服务器，具体访问规则在项目功能介绍**

# 项目介绍



通过实现一个账号系统，来演示如何构建一个真实的RESTful API风格服务器，通过项目展示了AP1构建过程中各个流程（准备->设计->开发->测试->部署)的实现方法, 详细为：

* 准备阶段
  * 安装和配置Go开发环境

* 设计阶段
  * API构建技术选型
  * API基本原理
  * API规范设计

* 开发阶段
  * 读取配置文件
  * 管理和记录日志
  * 做数据库的CURD操作
  * 自定义错误Code
  * 读取和返回HTTP请求
  * 进用户行业务逻辑开发
  * 对请求处理逻辑
  * 进行API身份验证
  * 进行HTTPS加密
  * 如何用Makefile管理API源码
  * 给API命令添加版本功能
  * 管理API命令
  * 生成Swagger在线文档
    

* 测试阶段
  * 如何进行单元测试
  * 进行性能测试（函数性能）
  * 做性能分析
  * API性能测试和调优

* 部署阶段
  * 用Nginx部署API服务

 REST风格虽然适用于很多传输协议，但在实际开发中，REST由于天生和HTTP协议相辅相成，因此HTTP协议
已经成了实现RESTful API事实上的标准。在HTTP协议中通过POST、DELETE、PUT、GET方法来对应REST资
源的增、删、改、查操作，具体的对应关系如下：

| HTTP方法 | 行为                     | URI          | 示例说明                |
| -------- | ------------------------ | ------------ | ----------------------- |
| GET      | 获取资源列表             | /users       | 获取用户列表            |
| GET      | 获取一个具体的资源       | /users/admin | 获取admin用户的详细信息 |
| POST     | 创建一个新的资源         | /users       | 创建一个新用户          |
| PUT      | 以整体的方式更新一个资源 | /users/1     | 更新id为1的用户         |
| DELETE   | 获取资源列表             | /users/1     | 删除id为1的用户         |




项目部署环境

* Ubuntu 20.04.4 LTS (GNU/Linux 5.4.0-91-generic x86_64)
* mysql v 5.6
* go sdk v1.18.1

* github.com/dgrijalva/jwt-go v3.2.0+incompatible
* github.com/fsnotify/fsnotify v1.5.1
* github.com/gin-contrib/pprof v1.3.0
* github.com/gin-gonic/gin v1.7.7
* github.com/jinzhu/gorm v1.9.16
* github.com/satori/go.uuid v1.2.0
* github.com/shirou/gopsutil v3.21.11+incompatible
* github.com/spf13/pflag v1.0.5
* github.com/spf13/viper v1.11.0
* github.com/swaggo/gin-swagger v1.4.1
* github.com/swaggo/swag v1.8.1
* github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125
* github.com/willf/pad v0.0.0-20200313202418-172aa767f2a4
* github.com/zxmrlc/log v0.0.0-20200612082315-fe407f734509
* golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
* gopkg.in/go-playground/validator.v9 v9.31.0

# 项目部署指南

* 服务器的MySQL数据库里导入如下db_apiserver数据库

```mysql
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `db_apiserver` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `db_apiserver`;

--
-- Table structure for table `tb_users`
--

DROP TABLE IF EXISTS `tb_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_tb_users_deletedAt` (`deletedAt`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_users`
--

LOCK TABLES `tb_users` WRITE;
/*!40000 ALTER TABLE `tb_users` DISABLE KEYS */;
INSERT INTO `tb_users` VALUES (0,'admin','$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG','2021-05-01 16:25:33','2021-05-30 16:25:33',NULL);
/*!40000 ALTER TABLE `tb_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-05-28  0:25:41

```

* [2.服务器上安装GoSDK v1.18.1详细教程](https://blog.csdn.net/weixin_52690231/article/details/124563906?utm_source=app&app_version=5.3.0&code=app_1562916241&uLinkId=usr1mkqgl919blen)
* 3.新建项目文件夹，并git初始化仓库

* 4.上传代码到服务器，并进行git提交

```go
mkdir restfulapiserver
cd restfulapiserver
git init 
git add .
git commit -m "2022.5.5"
```

* 5.修改conf/config.yaml里addr等信息
* 6.在项目根目录执行如下命令项目即可启动成功

```bash
go run main.go
```

# 项目功能

## 配置文件读取

## 记录和管理API日志

**启动项目后，在项目根目录执行如下命令项目即可查看日志**

```bash
tail -f ./log/apiserver.log
```

## 初始化MySQL建立连接

## API服务器健康状态自检

**检查API Server的服务器硬盘状态：**

[www.foolartist.top:6990/sd/disk](http://www.foolartist.top:6990/sd/disk)

**检查API Server的健康状况状态：**

 [www.foolartist.top:6990/sd/health](http://www.foolartist.top:6990/sd/health)

**检查API Server的CPU状态:**

[www.foolartist.top:6990/sd/cpu](http://www.foolartist.top:6990/sd/cpu)

**检查API Server的内存使用量状态:**

[www.foolartist.top:6990/sd/ram](http://www.foolartist.top:6990/sd/ram)

## Token

**获取Token成功**

```go
curl -XPOST -H "Content-Type: application/json"  -d'{"username":"admin","password":"admin"}' http://www.foolartist.top:6990/login 
```

**无Token禁止请求**

```go
curl -XPOST -H "Content-Type: application/json" -d'{"username":"user1","password":"user1234"}' http://www.foolartist.top:6990/v1/user
```

**有Token成功请求**

```bash
curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user -d'{"username":"user1","password":"user1234"}'
```

## 测试自定义错误信息

**返回errno.ErrBind错误**：

```bash
 curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user
```

**返回errno.InternalServerError错误：**

```bash
 curl -XPOST  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user -d'{"username":"admin"}'
```

**errno.ErrUserNotFound**

```bash
 curl -XPOST  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user  -d '{"password":"admin"}'
```

## 用户业务逻辑

**新增用户**

```bash
curl -XPOST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user -d'{"username":"user1","password":"user1234"}'
```

**查询用户列表**

```bash
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user  -d'{"offset": 0, "limit": 20}'
```

**获取用户详细信息**

```bash
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user/admin
```

**更新用户**

```bash
curl -XPUT -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user/2 -d'{"username":"kong","password":"kongmodify"}'
```

**删除用户**

```bash
curl -XDELETE -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://www.foolartist.top:6990/v1/user/2
```

## HTTPS加密请求

**请求时不携带CA证书和私钥**

```go
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://www.foolartist.top:6991/v1/user/admin
```

**请求时携带CA证书和私钥**

```GO
//需要加-k参数 ,该命令会在根目录
curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://www.foolartist.top:6991/v1/user/admin --cacert conf/server.crt --cert conf/server.crt --key conf/server.key   -k
```

## Makefile管理API项目

```go
//注意，必须得有.git文件存在，否则运行不起来；命令执行后项目目录下会生成apiserver文件
make
```

## API启动脚本

**查看api用法**

```bash
chmod 777 admin.sh
./admin.sh -h
```

**查看api状态**

```bash
 ./admin.sh status
```

**启动api状态**

```bash
//一定要必须得make命令执行完后，以下命令才有效
./admin.sh start
```

**停止api**

```bash
 ./admin.sh stop
```

**重新启动api状态**

```bash
 ./admin.sh restart
```

## 给API增加版本控制功能

```bash
./apiserver -v
```

## 基于Nginx部署API

## go test测试代码

**性能测试**

```go
cd util
go test
go test -v -count 2
```

**执行压力测试**

```
cd util
go test -test.bench=".*"
```

**查看性能并生成函数调用图**

安装graphviz

```bash
apt install graphviz
dot -version
```

在项目根目录下执行以下命令

```bash
//该命令会在根目录下生成cup.profile 和util.tetx.exe文件
go test -bench=".*" -cpuprofile=cpu.profile  ./util

//执行完以下命令进入到交互界面后执行top命令会显示性能
go tool pprof util.test cpu.profile
top

//执行完以上top命令会后继续进入到交互界面，输入svg后会在项目根目录生成svg图，需要以浏览器方式打开该文件进行浏览
svg
```

**测试覆盖率**

```bash
cd  util

//会在util目录下生成utilcover.out文件
go test -coverprofile=cover.out

go tool cover -func=cover.out
```

## API性能分析

**第一种方式获取profile采集信息**

```bash
//执行完以下命令进入到交互界面后执行topN命令会显示性能
go tool pprof http://www.foolartist.top:6990/debug/pprof/profile
```

**第二种方式获取profile采集信息**

[/debug/pprof/ (foolartist.top)](http://www.foolartist.top:6990/debug/pprof/)

## 生成Swagger在线文档

**Swagger在线文档**

[Swagger UI (foolartist.top)](http://www.foolartist.top:6990/swagger/index.html)



# 反馈交流

v :   x15516535379

