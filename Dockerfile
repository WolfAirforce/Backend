FROM golang:1.18-alpine

LABEL org.opencontainers.image.source="https://github.com/wolfairforce/backend"

WORKDIR /app

COPY . .

ENV CFG_FILE_PATH=/tmp/config.json

RUN go mod download
RUN go build -o /app/bin/api ./cmd/api

RUN chmod +x /app/bin/api

EXPOSE 8080

CMD [ "/app/bin/api" ]
