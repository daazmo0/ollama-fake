FROM golang:1.17 as builder

WORKDIR /app
COPY . .
RUN go build -o main .

FROM debian:12-slim
WORKDIR /app
COPY --from=builder /app/main /app/
EXPOSE 11434
CMD ["/app/main"]
