FROM golang:1.18-alpine

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY . .

RUN go version
RUN ls -al
RUN cat go.mod
RUN go mod tidy
RUN go build -o main ./cmd/api

CMD ["./main"]