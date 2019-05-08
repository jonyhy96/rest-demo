
FROM golang:1.11-alpine3.8
WORKDIR /go/src/rest-demo
COPY . /go/src/rest-demo
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/rest-demo .
ENTRYPOINT ["/go/bin/rest-demo"]
CMD ["-log_dir","log","-v","5"]