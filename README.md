## go-boilerplate
Golang service boilerplate using best practices.

Responsibility: Register (CRUD) and Login Users with JWT.

# Dependencies

- Gin-Gonic
- Swaggo
- go-redis
- GORM
- MySQL
- OpenTelemetry
- JWT-Go
- Crypto

## Documentation & Routes

**Swagger**
http://localhost:8000/auth/swagger/index.html

![Print of Swagger](/assets/swagger.png)

**Jaeger UI**
http://0.0.0.0:16686/search

![Print of Jaeger](/assets/trace-redis.png)

_**On Error**_
![Print of Jaeger](/assets/trace-error.png)

_**SQL**_
![Print of Jaeger](/assets/trace-sql.png)

## Composition

Pod = 1 Containers

- Container 1 = API REST

## Execution / Compilation

Set Env. variables of .env-sample:

Compile with:

```
make run
```

Update Swagger Doc:

```
make doc
```

## Tests

```
make test
```

## Path's Organization

- _**cmd**_: binaries entrypoint
- _**internal**_: private packages
    - _**configs**_: env. vars
    - _**controller**_: onde fica as regras/logicas
    - _**delivery**_: delivery layer -> http, grpc, messaging
    - _**domain**_: models / structs
    - _**repository**_: operacoes com storage/interface com banco
    - _**service**_: external services call's
- _**pkg**_: public packages
- _**docs**_: Swagger Documentation
