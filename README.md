# go-product-service

## Description
This is an example of implementation of Simple RESTful using clean architecture (Golang) projects.

### How To Run This Project

>  Make sure you already install mysql in your machine, after mysql installed and running. 
Run the product.sql in your mysql.


> Make sure you already install golang in your machine,
if you don't have golang in your machine you can install through this link https://golang.org/doc/install

##### Check golang already installed on youd machine
```bash
$ go version
go version go1.13 windows/amd64
```
Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

Clone this repository into your GOPATH
##### Check golang already installed on your machine
```bash
$ go version
```
##### clone project on gopath directory
```bash
$ cd $GOPATH/src/github.com/oryzel
$ git clone https://github.com/oryzel/go-product-service.git
```

>you must change config.json for database value appropriate to your database on machine

after that you can run your first golang app using this command
```bash
$  cd $GOPATH/src/github.com/oryzel/go-product-service
$  git run ./main.go
Service run on port :1001

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.11
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1001
```


### How To Deploy this project using docker

>  Make sure you already docker in your machine, please check this link https://docs.docker.com/install/ for the installation.

##### Check docker already installed on your machine
```bash
$ dokcer --version
Docker version 19.03.2, build 6a30dfc
```
You can build your docker image using this command.
> check your config.json and make sure your database host can be access from outside
```bash
$  cd $GOPATH/src/github.com/oryzel
$  docker build -t oryzel-project-service go-product-service
```

after finish build image finish, you can run your docker using this command
```bash
$  cd $GOPATH/src/github.com/oryzel
$  docker run -p 1001:1001 oryzel/go-product-service
```

to make sure your app already running on docker
```bash
$  docker ps
CONTAINER ID        IMAGE                       COMMAND               CREATED             STATUS              PORTS                    NAMES
099c2e249e25        oryzel/go-product-service   "./project-service"   10 minutes ago      Up 10 minutes       0.0.0.0:1001->1001/tcp   nifty_merkle
```