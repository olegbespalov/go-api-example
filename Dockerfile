ARG GO_VERSION=1.17

FROM golang:${GO_VERSION} as base

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

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-delve/delve/cmd/dlv

CMD air -c configs/.air.${APP_NAME}.toml