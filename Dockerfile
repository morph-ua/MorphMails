FROM golang:1.20.5-alpine@sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c AS builder
RUN apk add build-base
WORKDIR /source

ADD go.mod go.sum /source/
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod download -x

ADD internal .
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go build -o /app .

FROM busybox
WORKDIR /
COPY --from=builder /app /usr/local/bin/
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["server"]