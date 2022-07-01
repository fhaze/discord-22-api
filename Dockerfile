FROM golang:1.18-alpine as builder
WORKDIR /build
COPY . /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/cmd .
EXPOSE 8888
CMD ["/app/cmd"]
