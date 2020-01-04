FROM golang:1.13.5 as builder
WORKDIR /go/src/cli
COPY ./src/ .
RUN GOOS=darwin GOARCH=amd64 go install ./...
# RUN GOOS=linux GOARCH=amd64 go install ./...

From alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/ /app
