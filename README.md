# 基于Go语言实现的RESTful API

# 项目部署环境

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

* 1.服务器的MySQL数据库里导入如下db_apiserver数据库

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

* 3.上传仓库代码到服务器addr等信息
* 4.修改conf/config.yaml里
* 5.在项目根目录执行如下命令项目即可启动成功

```go
go run main.go
```

# 项目功能

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

```bash
 curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://127.0.0.1:8081/v1/user/admin
```

```bash
 curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://www.foolartist.top:6991/v1/admin
```

**请求时不携带CA证书和私钥**

```go
 curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://127.0.0.1:8081/v1/user/admin --cacert conf/server.crt --cert conf/server.crt --key conf/server.key
```

```bash
 curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" https://www.foolartist.top:6991/v1/admin --cacert conf/server.crt --cert conf/server.crt --key conf/server.key
 
 
```

## Makefile管理API项目

```go
make
```

## API启动脚本

**查看api用法**

```bash
 ./admin.sh -h
```

**查看api状态**

```bash
 ./admin.sh status
```

**启动api状态**

```bash
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

# 反馈交流

>  v :x15516535379

