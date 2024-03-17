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

FROM scratch AS bin
COPY --from=build /out/example /
