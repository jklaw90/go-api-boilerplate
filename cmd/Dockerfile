FROM golang:1.15.2-alpine3.12 AS base-builder
ENV CGO_ENABLED=0
WORKDIR /go/src/hospital
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd cmd/
COPY pkg pkg/
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64


FROM scratch as base-server
COPY --from=base-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base-builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip


FROM base-builder as api-builder
RUN go build -o /go/bin/api cmd/api/main.go

FROM base-server as api
COPY --from=api-builder /go/bin/api /api
ENTRYPOINT ["/api"]