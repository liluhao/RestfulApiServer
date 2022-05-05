# 基于Go语言实现的RESTful API

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

基于Nginx部署API

