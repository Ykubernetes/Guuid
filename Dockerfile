FROM golang:1.20.11 AS builder


ENV GO111MODULE=on \
    GOPROXY=goproxy.io \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o guuid

FROM scratch

COPY --from=builder /app/guuid /
# USER nonroot:nonroot
EXPOSE 9000
CMD ["/guuid"]

