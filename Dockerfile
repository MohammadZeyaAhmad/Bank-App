#Build Stage
FROM golang:1.21.0-alpine3.17 AS builder 
WORKDIR /app
COPY . .
RUN go build -o main main.go
#Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
RUN chmod +x start.sh
COPY wait-for.sh .
EXPOSE 8080
RUN chmod +x /app/start.sh
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
