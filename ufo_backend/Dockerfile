FROM golang:1.21.4-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/ufobackend
FROM alpine:latest
EXPOSE 8080
ENV localhost "0.0.0.0"
WORKDIR /app
COPY --from=build /app/ufobackend .
ENV MONGODB_UFO_URI mongodb://rust-scraper:iwanttobelieve@172.20.0.5:27017/i_want_to_believe?retryWrites=true&w=majority&authSource=admin
CMD ["./ufobackend"]