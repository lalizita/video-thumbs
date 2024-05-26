FROM golang:1.22 as builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o thumbs main.go

FROM scratch
COPY --from=builder /app/thumbs /thumbs
CMD ["/thumbs"]

EXPOSE 8000