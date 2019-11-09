FROM gcr.io/cloud-builders/go:alpine as builder
WORKDIR /work
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /work/app /app
CMD ["/app"]