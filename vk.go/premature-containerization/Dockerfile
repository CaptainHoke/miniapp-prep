FROM golang:alpine AS builder
WORKDIR /src

COPY go.* .
RUN go mod download
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build go build -o /out/example .

FROM scratch AS bin
COPY --from=builder /out/example /
