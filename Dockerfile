FROM golang:1.17 AS build
WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GoNotify .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/GoNotify .
COPY config/config.yaml .
CMD ["./GoNotify"]
