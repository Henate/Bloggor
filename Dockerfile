FROM golang:latest

WORKDIR $GOPATH/src/Bloggor
COPY . $GOPATH/src/Bloggor
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]