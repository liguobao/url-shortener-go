FROM golang:1.14

WORKDIR $GOPATH/src/url-shortener-go
COPY . $GOPATH/src/url-shortener-go
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./url-shortener-go"]