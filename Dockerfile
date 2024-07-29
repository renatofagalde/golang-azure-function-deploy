#build stage
FROM golang:1.22.4-alpine3.20 AS builder
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o . 

## RUN STAGE
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/* .

EXPOSE 8080
CMD ["/app/bootstrap"]

