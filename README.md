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
1. go get github.com/jmoiron/sqlx

#### Examples

go  run main.go
