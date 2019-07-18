FROM golang:1.12.4 as build

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache
ADD pkg/sim-common/ /pkg/sim-common/
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o gin_scaffold main.go

FROM alpine:3.9 as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/gin_scaffold /
COPY --from=build /go/release/conf/app-prd.ini ./conf/app.ini
CMD ["/gin_scaffold"]