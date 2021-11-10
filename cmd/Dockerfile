FROM golang:1.16
# skaffold 
WORKDIR /go/src/boilerplate
COPY go.mod go.sum ./
RUN go mod download
COPY cmd cmd/