# go-server

Go语言实现的简单HTTP服务端

## Environment

* Ubuntu 18.04
* Go 1.12.5
* PostgreSQL 9.5

## Configure

`config.ini`

```
[PostgreSQL]
host=host
user=db_user
dbname=db_name
password=db_password
```

## Run

`go run main.go`

## Dependencies

[gin](https://github.com/gin-gonic/gin)

[gorm](https://github.com/jinzhu/gorm)
