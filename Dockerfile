ARG GO_VERSION=1.17

FROM golang:${GO_VERSION} as builder

ARG APP_NAME
ARG APP_PORT

ENV APP_PORT=${APP_PORT} \
    APP_NAME=${APP_NAME} \
    APP_PATH="/var/app"

COPY . ${APP_PATH}

WORKDIR ${APP_PATH}

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"

RUN go build -a -o /main cmd/api/*.go

FROM builder as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go get -u github.com/go-delve/delve/cmd/dlv

CMD $(go env GOPATH)/bin/air -c configs/.air.${APP_NAME}.toml

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /main ./

RUN chmod +x ./main

ENTRYPOINT ["./main"]