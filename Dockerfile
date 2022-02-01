FROM golang:1.16.1-alpine AS builder

RUN apk update && apk add --no-cache musl-dev gcc build-base libc-dev curl git ca-certificates &&  \
    update-ca-certificates

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src

ADD go.mod go.sum $GOPATH/src/

RUN GOOS=linux go mod download

COPY . $GOPATH/src

RUN GOOS=linux go build -ldflags '-linkmode=external' -o /go/bin/api cmd/api/main.go

FROM alpine:3.14

WORKDIR /app

# UPDATE APK CACHE AND INSTALL PACKAGES | CONFIGURE TIMEZONE | INSTALL AWS DEPS
RUN apk update && apk upgrade && apk add --no-cache \
    tzdata ca-certificates && \
    cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime; echo "America/Sao_Paulo" > /etc/timezone

# Copy our static executable
COPY --from=builder /go/bin/api .