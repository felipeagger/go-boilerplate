FROM golang:1.16.1-alpine AS builder

# INSTALL PACKAGES
RUN apk update && apk add --no-cache musl-dev gcc build-base libc-dev curl git tzdata ca-certificates &&  \
    update-ca-certificates && echo "America/Sao_Paulo" > /etc/timezone

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src

ADD go.mod go.sum $GOPATH/src/

RUN GOOS=linux go mod download

COPY . $GOPATH/src

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/api cmd/api/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone

# Copy migrations and static executable
WORKDIR /app/internal/repository/migrations
COPY --from=builder /go/src/internal/repository/migrations/ .

WORKDIR /app
COPY --from=builder /go/bin/api .

CMD ["./api"]