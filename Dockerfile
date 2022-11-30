FROM golang:1.19.3-alpine AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o portsd ./cmd/rest-memory-service

FROM scratch

COPY --from=builder /build/portsd .
ENTRYPOINT ["./portsd"]
