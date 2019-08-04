FROM golang:1.12-alpine

WORKDIR /go/src/github.com/abyssparanoia/rapid-go-woker/
COPY . .

ENV GO111MODULE=on

ENV PORT 8080
EXPOSE 8080

RUN apk --no-cache --update upgrade \
    && apk add --no-cache git alpine-sdk \
    && go get -u github.com/codegangsta/gin

CMD gin --port 8081 --appPort 8080 --immediate run main.go routing.go dependency.go environment.go
