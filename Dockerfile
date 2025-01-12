FROM golang:1.23.4-alpine3.21 AS builder
LABEL authors="w2k"

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/http/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ARG PORT
ENV PORT=$PORT

EXPOSE $PORT

CMD ["./server"]

