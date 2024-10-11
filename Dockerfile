FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY polymarket_scraper ./polymarket_scraper

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /polymarket_scraper

CMD ["/polymarket_scraper"]

# Stage 2: Final Image
FROM alpine:latest

COPY --from=builder /polymarket_scraper /polymarket_scraper

CMD ["/polymarket_scraper"]