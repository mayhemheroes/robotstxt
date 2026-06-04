FROM golang:1.21 AS builder

ENV GOTOOLCHAIN=auto

WORKDIR /src
COPY . .

RUN go test -c -o /parsebytes-fuzz .

FROM ubuntu:22.04
COPY --from=builder /parsebytes-fuzz /
