FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o dns_ddoser ./cmd/dns_ddoser

FROM scratch
COPY --from=builder /app/dns_ddoser /dns_ddoser
ENTRYPOINT [ "/dns_ddoser" ]
