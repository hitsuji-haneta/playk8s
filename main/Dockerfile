FROM golang:1.13.6-alpine3.10 as builder
COPY . /app
WORKDIR /app
RUN go build -o /playk8s /app/main.go

FROM alpine:3.10
EXPOSE 8080
WORKDIR /app
RUN echo fuga
RUN mkdir templates
COPY --from=builder /playk8s .
CMD [ "./playk8s" ]