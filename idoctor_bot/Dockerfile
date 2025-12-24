FROM golang:1.25.3-alpine3.21 as builder

# ENV GOPATH /go
# ENV PATH $PATH:$GOPATH/bin

RUN set -ex && \
  apk add --no-cache gcc musl-dev git


ARG CI_JOB_LOGIN
ARG CI_JOB_TOKEN

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Сборка Go-приложения
RUN go build -ldflags "-s -w" -o main .

# Делаем бинарник исполняемым
RUN chmod +x main
