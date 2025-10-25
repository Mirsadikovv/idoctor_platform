FROM golang:1.25.3-alpine3.21 as builder

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
ENV GOPRIVATE=git.sriss.uz

RUN set -ex && \
  apk add --no-cache gcc musl-dev git


ARG CI_JOB_LOGIN
ARG CI_JOB_TOKEN
ARG MODE=dev

WORKDIR /app

COPY go.mod ./

RUN git config --global url."https://git.sriss.uz/".insteadOf "http://git.sriss.uz/"

RUN printf "machine git.sriss.uz\nlogin %s\npassword %s\n" "$CI_JOB_LOGIN"  "$CI_JOB_TOKEN" > ~/.netrc && \
  chmod 600 ~/.netrc

RUN go mod download

COPY . .

RUN if [ "$MODE" == "dev" ]; then \
  find src/module -type f -name "cmd.go" | while read -r file; do \
    dir=$(dirname "$file"); \
    module_name=$(basename "$dir"); \
    echo "Generating Swagger for $module_name..."; \
    (cd "$dir" && go tool swag init -pd --ot json --o "../../docs/$module_name" -g cmd.go > /dev/null 2>&1) || echo "Error in $module_name"; \
    done \
  fi    

ARG HTTP_PORT=80
EXPOSE ${HTTP_PORT}

# Сборка Go-приложения
RUN go build -ldflags "-s -w" -o main .
