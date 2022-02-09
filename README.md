## go-boilerplate
Golang API Service boilerplate using best practices of Clean Architecture.

Responsibility: CRUD and Login Users with JWT.

# Dependencies

- Gin-Gonic
- Swaggo
- go-redis
- GORM (MySQL)
- Snowflake
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

## Compose Stack

- API REST
- MySQL
- Redis
- Jaeger

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

- _**assets**_: static files
- _**cmd**_: binaries entrypoint
- _**internal**_: private packages
    - _**configs**_: env. variables
    - _**usecase**_: business logical/rules
    - _**delivery**_: delivery layer -> http, grpc, messaging
    - _**entity**_: entities / schemas
    - _**repository**_: storage operations, database interface
    - _**service**_: external services call's
- _**pkg**_: public packages
- _**docs**_: Swagger Documentation
