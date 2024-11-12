FROM golang:1.22-alpine as build


WORKDIR /app

COPY . /app

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./api-server cmd/server/main.go




FROM alpine:latest
MAINTAINER charlie "charlie@mt.social"

USER root

COPY --from=build /app/.imhub /app/.imhub
RUN chmod +x /app/hub

EXPOSE 26656 26657 1317 9090

WORKDIR /app

CMD ["imhub"]