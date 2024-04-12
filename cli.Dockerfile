FROM golang:1.22-alpine AS builder

RUN adduser -D -u 1001 apiuser

WORKDIR /cli

COPY . . 

RUN chown -R apiuser:apiuser /cli
RUN --mount=type=cache,target=/go/pkg/mod \
     go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod \
    go build -tags "pro" -o cli ./cmd/cli 

FROM scratch

COPY --from=builder /cli/cli /cli
COPY --from=builder /etc/passwd /etc/passwd

USER apiuser

ENTRYPOINT ["./cli"]