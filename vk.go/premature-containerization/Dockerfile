FROM golang:alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /out/example .

FROM scratch AS bin
COPY --from=builder /out/example /
