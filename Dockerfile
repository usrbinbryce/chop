FROM golang:1.27.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chop_server ./cmd/chop

FROM scratch
COPY --from=builder /app/chop_server /chop_server
EXPOSE 8080
ENTRYPOINT ["/chop_server"]
