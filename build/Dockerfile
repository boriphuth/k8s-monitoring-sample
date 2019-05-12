FROM golang:1.12

ADD . /workdir/
WORKDIR /workdir

RUN go build -o app .

EXPOSE 8080

ENTRYPOINT ["./app"]