FROM golang:1.25-alpine AS builder
WORKDIR /src
# Copy only go.mod here — a stdlib-only service has no go.sum,
# and naming go.sum as a COPY source would fail the build when it is absent.
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o /out/app ./

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/app /app
RUN mkdir -p /data
EXPOSE 9090
ENTRYPOINT ["/app"]
