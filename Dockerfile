FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/bin/api ./cmd/api

EXPOSE 8080

RUN chmod +x /app/bin/api

CMD [ "/app/bin/api" ]
