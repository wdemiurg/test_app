FROM golang:1.16-alpine

WORKDIR /app


COPY *.go ./

RUN go mod init main.go
RUN go mod tidy
RUN go build -o /test_app

EXPOSE 8080

CMD [ "/test_app" ]
