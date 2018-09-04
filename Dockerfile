FROM golang as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o count .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/count .
EXPOSE 80
ENTRYPOINT ./count
