
# ----- build -----

FROM golang:alpine as builder

COPY * ./
# RUN pwd # >> /go
RUN go build -ldflags="-s -w" -trimpath .
