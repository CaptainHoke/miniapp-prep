FROM golang:1.22.1-alpine AS base
WORKDIR /src

COPY go.* .
RUN go mod download
COPY . .

FROM base AS build
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /out/example .

FROM base AS test
RUN --mount=type=cache,target=/root/.cache/go-build \
    go test -v .

FROM golangci/golangci-lint:v1.56.2-alpine AS lint-base

FROM base AS lint
COPY --from=lint-base /usr/bin/golangci-lint /usr/bin/golangci-lint
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 1m0s ./...

FROM scratch AS bin
COPY --from=build /out/example /
