# Package delivery service

## Description

This is an educational http service for package delivering using **Golang**.
At first, you have to register. Then you can authenticate using the TWT token and manage your packages and deliveries

* **Framework**: Fiber
* **ORM**: Gorm
* **Logs**: Zero-log
* **SQL dialect**: pgSql

## Service works with following entities:

* User
* Package
* Delivery

## Layout

```tree
├── cmd
│   └── app
│       └── main.go
├── config
│   └── config.go
├── docs
│   └── swagger
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
├── internal
│   ├── app
│   │   └── app.go
│   └── controller
│       └── http
│           └── v1
│               ├── middleware
│               │   ├── dbtransaction.go
│               │   └── jwttoken.go   
│               ├── delivery.go
│               ├── error.go
│               ├── msg.go
│               ├── package.go
│               ├── response.go
│               ├── router.go
│               ├── user.go
│               └── utils.go
├── entity
│   ├── delivery.go
│   ├── package.go
│   ├── token.go
│   └── user.go
├── tokens
│   └── token.go
├── usecase
│   ├── delivery_repo
│   │   └── delivery_gorm.go
│   ├── package_repo
│   │   └── package_gorm.go
│   ├── token_repo
│   │   └── token_gorm.go
│   ├── user_repo
│   │   └── user_gorm.go
│   ├── delivery.go
│   ├── errors.go
│   ├── interfaces.go
│   ├── package.go
│   └── user.go
├── validations
│   ├── responce.go
│   └── validatons.go
├── pkg
│   ├── database
│   │   └── gormdb.go
│   ├── logger
│   │   └── logger.go
│   └── utils
│       └── password.go
├── Makefile
└── README.md
```

## Run

### Requirements:

* Golang
* Make
* PgSQL

### Config

Config populates from environment variables

| Name        | Description                       | Default value | Expected value         | Requiered |
|:------------|:----------------------------------|:--------------|:-----------------------|:---------:|
| APP_NAME    | Application`s name                | myAppName     | any string             |    ✔️     |
| APP_PORT    | Port which servers`s listening on | :8080         | port starting with `:` |    ✔️     |
| JWT_SECRET  | Secret for JWT token generation   | jwtSecret     | any string             |    ✔️     |
| DB_HOST     | Database sever host               | localhost     | any string             |    ✔️     |
| DB_PORT     | Database sever port               | 5432          | port without `:`       |    ✔️     |
| DB_NAME     | Database name                     | -             | any string             |    ✔️     |
| DB_USER     | Username for database             | -             | any string             |    ✔️     |
| DB_PWD      | Password for database             | -             | any string             |    ✔️     |
| DB_SSL_MODE | Disable or enable SSL mode        | disable       | `enabled`/`disabeld`   |    ✔️     |
| DB_DNS      | Database DNS                      | -             | any string             |     ❌     |


```Shell
make build
./service
```
### API

Swagger is available at `./docs/swagger.yaml`

Or at `localhost:APP_PORT/swagger/index.html` after you run the service

