ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk add --no-cache ca-certificates git
COPY . /src
WORKDIR /src

ENV CGO_ENABLED=0
RUN go build -o main /src/main.go

FROM scratch AS final

#RUN apk add --no-cache ca-certificates git && rm -rf /var/cache/apk/*
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/main /main

EXPOSE ${API_PORT}

VOLUME ["/cert-cache"]
ENTRYPOINT [ "/main" ]
