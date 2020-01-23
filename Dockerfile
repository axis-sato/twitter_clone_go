FROM golang:1.13.6-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache alpine-sdk \
    && go build -o dist/app \
    && go get -u github.com/go-delve/delve/cmd/dlv \
    && go get gopkg.in/urfave/cli.v2@master \
    && go get github.com/oxequa/realize \
    && go get github.com/volatiletech/sqlboiler \
    && go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql \
    && go get bitbucket.org/liamstask/goose/cmd/goose


FROM alpine

WORKDIR /app

COPY --from=build /go/app/dist/app .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]