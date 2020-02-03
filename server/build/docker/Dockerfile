FROM golang:latest as builder

COPY . /backend
WORKDIR /backend

RUN cd cmd/backend && CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o hello

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /backend/cmd/hello .

ENTRYPOINT [ "./backend" ]