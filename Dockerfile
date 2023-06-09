FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/acs-dl/identity-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/identity-svc /go/src/github.com/acs-dl/identity-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/identity-svc /usr/local/bin/identity-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["identity-svc"]
