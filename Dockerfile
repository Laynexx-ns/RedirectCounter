FROM golang:1.24-alpine3.20 as builder


WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /server .

EXPOSE 8080

CMD ["./server"]