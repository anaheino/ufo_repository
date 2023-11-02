FROM golang:1.21.3-alpine AS build
ENV MONGODB_UFO_URI="mongodb://rust-scraper:iwanttobelieve@127.0.0.1:27017/i_want_to_believe?retryWrites=true&w=majority"
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/ufobackend

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/ufobackend .

CMD ["./ufobackend"]