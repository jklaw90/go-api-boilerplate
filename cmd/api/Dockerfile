ARG BASE_BUILDER
FROM $BASE_BUILDER as builder
ARG SKAFFOLD_GO_GCFLAGS
RUN echo "Go gcflags: ${SKAFFOLD_GO_GCFLAGS}"
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -mod=readonly -o /go/bin/api cmd/api/main.go

FROM gcr.io/distroless/base
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip
ENV GOTRACEBACK=single
COPY --from=builder /go/bin/api /api
ENTRYPOINT ["/api"]