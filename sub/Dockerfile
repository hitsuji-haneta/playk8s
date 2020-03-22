FROM golang:1.13.6-alpine3.10 as builder
COPY . /app
WORKDIR /app
RUN go build -o /main /app/main.go

FROM alpine:3.10
EXPOSE 8080
WORKDIR /app
COPY --from=builder /main .
CMD [ "./main" ]