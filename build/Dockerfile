FROM golang:1.12

ADD . /workdir/app
WORKDIR /workdir/

RUN go build -o app .

EXPOSE 8080

ENTRYPOINT ["./app"]