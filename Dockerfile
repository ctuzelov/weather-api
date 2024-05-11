FROM golang:alpine AS builder
LABEL project-name="weather"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o weather cmd/weather/main.go

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .
# COPY wait-for.sh .
RUN chmod +x weather

CMD ["./weather"]

EXPOSE 9003 8003
