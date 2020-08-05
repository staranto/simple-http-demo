FROM golang:1.13.4 AS builder

COPY . /tmp
WORKDIR /tmp

RUN go get -u github.com/gorilla/mux
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux \
    GOARCH=${TARGETARCH} \
    go build -o /tmp/pounder pounder.go
RUN chmod +x /tmp/pounder


FROM alpine

COPY --from=builder /tmp/pounder /tmp/pounder

CMD ["/tmp/pounder"]
