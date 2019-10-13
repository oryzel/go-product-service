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
##### Check golang already installed on youd machine
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