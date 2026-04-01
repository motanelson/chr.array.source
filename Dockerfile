FROM  golang:latest

COPY hello.go /hello.go
RUN chmod +x /hello.go

RUN go build /hello.go

CMD ["/go/hello"]
