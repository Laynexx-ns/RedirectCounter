FROM golang1.24:alpine3 as bulder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder ./service .

EXPOSE 8080

CMD ["./server"]