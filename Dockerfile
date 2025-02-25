FROM golang:1.22-alpine AS builder

WORKDIR /go/workspace

COPY main.go main.go

RUN go build main.go

FROM alpine:3.17

WORKDIR /workspace

COPY --from=builder /go/workspace/main ./main

CMD ["./main"]
