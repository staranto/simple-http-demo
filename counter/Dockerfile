FROM golang:1.13.4 AS builder

COPY . /tmp
WORKDIR /tmp

RUN go get -u github.com/gorilla/mux
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux \
    GOARCH=${TARGETARCH} \
    go build -o /tmp/counter counter.go
RUN chmod +x /tmp/counter


FROM alpine

COPY --from=builder /tmp/counter /tmp/counter

HEALTHCHECK --interval=3s --timeout=2s --start-period=15s --retries=1 \
    CMD wget --quiet --tries=1 --spider http://localhost:8282/live || exit 1

CMD ["/tmp/counter"]
