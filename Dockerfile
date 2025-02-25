FROM golang:1.23-alpine AS builder

WORKDIR /go/workspace

COPY main.go main.go

RUN go build main.go

FROM alpine:latest

WORKDIR /workspace

COPY --from=builder /go/workspace/main ./main

CMD ["./main"]
