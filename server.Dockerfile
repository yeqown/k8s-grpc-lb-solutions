FROM golang:1.14-alpine3.11 as build

WORKDIR /build
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o app ./examples/grpc-server

# release
FROM alpine as release

WORKDIR /server

COPY --from=build /build/app /server/app