FROM golang:1.17 as builder

WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/
EXPOSE 11434
CMD ["/app/main"]
