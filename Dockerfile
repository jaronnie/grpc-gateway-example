FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    PATH="/app:${PATH}"

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o app main.go

EXPOSE 9603

EXPOSE 8090 

ENTRYPOINT ["app"]
