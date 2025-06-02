FROM golang:1.24-alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o expense-tracker .

RUN chmod +x expense-tracker


FROM alpine:latest
RUN mkdir /app

COPY --from=builder app/expense-tracker /app/expense-tracker

ENTRYPOINT ["/app/expense-tracker"]