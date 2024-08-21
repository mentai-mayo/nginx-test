
# ----- build -----

FROM golang:alpine as builder

COPY * ./
# RUN pwd # >> /go
RUN go build -ldflags="-s -w" -trimpath .

# ----- prod -----

FROM nginx:latest

# copy domain socket server from builder
COPY --from=builder /go/go-domain-sock /home/golang/main
