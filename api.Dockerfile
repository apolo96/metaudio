# Builder Stage
FROM golang:1.22.2-alpine as builder

WORKDIR /app

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
     go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod \
    go build -o api ./cmd/api

# Runtime Stage
FROM scratch

COPY --from=builder /app/api api

EXPOSE 8000

CMD ["./api"]




