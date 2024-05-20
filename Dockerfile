FROM golang:1.22-alpine AS builder
ENV CGO_ENABLED=1
RUN apk add --update \
  alpine-sdk \
  sqlite-dev
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine
RUN apk add --update \
  alpine-sdk \
  sqlite-dev
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]

