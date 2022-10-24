# purpose: to test for local deployment
FROM golang:1.16-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o server main.go

#RUN STAGE
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/server .
COPY app.env .

EXPOSE 8080

CMD [ "/app/server" ]