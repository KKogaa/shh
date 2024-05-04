FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8080 

CMD ["/app/main"]

