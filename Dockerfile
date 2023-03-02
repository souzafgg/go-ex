FROM golang:alpine3.17 AS builder

WORKDIR /app
COPY main.go .
RUN go build -o app

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/app .
CMD [ "./app" ]