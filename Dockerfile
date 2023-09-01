#Build Stage
FROM golang:1.21.0-alpine3.17 AS builder 
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
#Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
RUN chmod +x start.sh
COPY wait-for.sh .
COPY db/migration ./migration
EXPOSE 8080
RUN chmod +x /app/start.sh
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
