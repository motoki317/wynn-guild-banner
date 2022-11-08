FROM golang:alpine AS build

ENV GOCACHE=/tmp/go/cache
WORKDIR /go/src/github.com/motoki317/wynn-guild-banner

COPY ./go.* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/tmp/go/cache CGO_ENABLED=0 go build -o /wynn-guild-banner -ldflags "-s -w"

FROM alpine:latest

WORKDIR /work

COPY --from=build /wynn-guild-banner ./

ENTRYPOINT ["./wynn-guild-banner"]
