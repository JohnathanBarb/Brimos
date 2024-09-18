FROM golang:1.23 AS go-lang

WORKDIR /build

COPY . ./

RUN go mod tidy

RUN go build -o /server

EXPOSE 9000
ENTRYPOINT ["/server"]
