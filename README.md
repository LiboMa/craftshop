## This is the demo app of online shopping

#### Reqiurements

env:
```
go v1.12
gin
redis
mysql

```
#### Structure

.
├── README.md
├── assets
│   ├── css
│   ├── images
│   ├── js
│   └── videos
├── common
│   ├── cache.go
│   ├── database.go
│   └── utils.go
├── conf
├── docs
│   ├── addProduct.sh
│   ├── eshop.sql
│   └── seeds.sql
├── main.go
├── orders
├── pkg
├── products
│   ├── doc.go
│   ├── models.go
│   ├── routers.go
│   └── serializer.go
├── test
│   ├── main.go
│   ├── sqlx_main.go
│   ├── testinterface.go
│   └── time.go
├── users
└── vendor

#### Installation

Mysql driver & mapping tools written by GO
go get github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
go get github.com/jmoiron/sqlx
go get -u github.com/go-redis/redis
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jasonlvhit/gocron

#### Build
Linuxbuild="env GOOS=linux GOARCH=amd64 go build"
Macbuild="go build"
#### Examples

go  run main.go
