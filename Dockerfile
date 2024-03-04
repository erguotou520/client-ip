FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main -ldflags="-s -w"

FROM apine:latest as final

COPY --from=builder /app/main /app/main

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]