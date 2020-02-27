FROM golang:1.13.6-alpine3.10 as builder
COPY . /playk8s
WORKDIR /playk8s
RUN go build -o /go-app /playk8s/main.go

FROM alpine:3.10
EXPOSE 8080
COPY --from=builder /go-app .
COPY --from=builder /playk8s/templates ./templates
CMD [ "./go-app" ]