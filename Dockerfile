FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o app

FROM debian:bookworm-slim
WORKDIR /root/

COPY --from=builder /app/app .

ENV POSTGRES_DSN="postgres://postgres:postgres@db:5432/articles?sslmode=disable"

EXPOSE 8080
CMD ["./app"]
